package file

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func (f *File) Upload(w http.ResponseWriter, r *http.Request) {
	//fmt.Print("Upload")
	//fmt.Fprintf(w, "Upload") //这个写入到w的是输出到客户端的

	if r.Method == "GET" {
		f.showUploadPage(w)
	} else if r.Method == "POST" {
		f.receiveFile(w, r)
	}
}

// showUploadPage
func (f *File) showUploadPage(w http.ResponseWriter) {
	// 读取 static/view/index.html 文件
	data, err := os.ReadFile("static/view/index.html")
	if err != nil {
		_, _ = io.WriteString(w, "internel server error")
		log.Println("os.ReadFile(\"static/view/index.html\") ", err.Error())
	}

	// 将读到的 html 文件以 string 的形式返回客户端
	_, err = io.WriteString(w, string(data))
	if err != nil {
		log.Println("io.WriteString(w, string(data)) ", err.Error())
	}
}

// receiveFile
func (this *File) receiveFile(w http.ResponseWriter, r *http.Request) {

	// 秒传
	if this.FastUpload(w, r) {
		return
	}

	// 接收文件
	f, head, err := r.FormFile("file")
	if err != nil {
		log.Println("r.FormFile(\"file\") ", err.Error())
		return
	}
	defer f.Close()

	// 本地创建文件
	newFile, err := os.Create("tmp/" + head.Filename)
	if err != nil {
		log.Println("os.Create(\"/tmp/\" + head.Filename) ", err.Error())
		return
	}

	// copy 文件
	fileSize, err := io.Copy(newFile, f)
	if err != nil {
		log.Println("io.Copy(newFile, file) ", err.Error())
		return
	}
	newFile.Seek(0, 0)

	// 更新数据库
	tFile := mysql.NewFile()
	fileSha1 := file.FileSha1(newFile)
	tFile.SetAttrs(map[string]interface{}{
		"UploadAt": time.Now().Format("2006-01-02 15:04:05"),
		"Name":     head.Filename,
		"Dir":      "tmp/",
		"Size":     fileSize,
		"Sha1":     fileSha1,
	})
	success := tFile.Insert()
	if success == false {
		// 上传失败 页面跳转
		// 根据当前路由 重定向
		currentRoute := r.URL.Path
		http.Redirect(w, r, currentRoute+"/duplicate", http.StatusFound)
	} else {
		// 上传UserFile
		userFile := mysql.NewUserFile()
		username := r.Form.Get("username")
		success = userFile.Insert(username, fileSha1, head.Filename, fileSize)
		if !success {
			// 上传失败 页面跳转
			// 根据当前路由 重定向
			currentRoute := r.URL.Path
			http.Redirect(w, r, currentRoute+"/duplicate", http.StatusFound)
		} else {
			// 上传成功 页面跳转
			// 根据当前路由 重定向
			currentRoute := r.URL.Path
			http.Redirect(w, r, currentRoute+"/success", http.StatusFound)
		}
	}
}

// FastUpload 秒传接口
func (f *File) FastUpload(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	// 1.1 解析请求参数
	username := r.Form.Get("username")

	// 1.2 filehash
	tf, head, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file from request", http.StatusBadRequest)
		return false
	}
	defer tf.Close()

	// 创建SHA-1哈希对象
	h := sha1.New()
	// 将文件内容传入哈希对象
	filesize, err := io.Copy(h, tf)
	if err != nil {
		http.Error(w, "Failed to read file content", http.StatusInternalServerError)
		return false
	}
	// 计算哈希值
	hashInBytes := h.Sum(nil)
	// 将哈希值转换成16进制字符串
	filehash := hex.EncodeToString(hashInBytes)

	// 1.3 filename
	filename := head.Filename

	// 2. 从文件表中查询相同hash的文件记录
	fileMeta, err := mysql.NewFile().Query(filehash)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	// 3. 查不到记录则返回秒传失败
	if fileMeta == nil {
		//resp := handler.RespMsg{
		//	Code: -1,
		//	Msg:  "秒传失败，请访问普通上传接口",
		//}
		//w.Write(resp.JSONBytes())
		return false
	}

	// 4. 上传过则将文件信息写入用户文件表， 返回成功
	userFile := mysql.NewUserFile()
	suc := userFile.Insert(username, filehash, filename, int64(filesize))
	if suc {
		//resp := handler.RespMsg{
		//	Code: 0,
		//	Msg:  "秒传成功",
		//}
		//w.Write(resp.JSONBytes())
		return true
	}
	//resp := handler.RespMsg{
	//	Code: -2,
	//	Msg:  "秒传失败，请稍后重试",
	//}
	//w.Write(resp.JSONBytes())
	return false
}
