package file

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
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
	size, err := io.Copy(newFile, f)
	if err != nil {
		log.Println("io.Copy(newFile, file) ", err.Error())
		return
	}
	newFile.Seek(0, 0)

	// 更新数据库
	userFile := mysql.NewFile()
	userFile.SetAttrs(map[string]interface{}{
		"UploadAt": time.Now().Format("2006-01-02 15:04:05"),
		"Name":     head.Filename,
		"Dir":      "tmp/",
		"Size":     size,
		"Sha1":     file.FileSha1(newFile),
		"UserId":   0,
	})
	success := userFile.Insert()
	if success == false {
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
