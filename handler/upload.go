package handler

import (
	"cloud_storage/file"
	"cloud_storage/file/util"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type UploadHandler struct {
}

func (uh *UploadHandler) Handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Print("UploadHandler")
	//fmt.Fprintf(w, "UploadHandler") //这个写入到w的是输出到客户端的

	if r.Method == "GET" {
		uh.showUploadPage(w)
	} else if r.Method == "POST" {
		uh.receiveFile(w, r)
	}
}

// showUploadPage
func (uh *UploadHandler) showUploadPage(w http.ResponseWriter) {
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
func (uh *UploadHandler) receiveFile(w http.ResponseWriter, r *http.Request) {
	// 接收文件
	f, head, err := r.FormFile("file")
	if err != nil {
		log.Println("r.FormFile(\"file\") ", err.Error())
		return
	}
	defer f.Close()

	// 文件 meta
	newfileMeta := file.NewFileMeta()
	newfileMeta.SetUploadAt(time.Now().Format("2006-01-02 15:04:05"))
	newfileMeta.SetName(head.Filename)
	newfileMeta.SetLocation("tmp/" + head.Filename)

	// 本地创建文件
	newFile, err := os.Create(newfileMeta.GetLocation())
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
	newfileMeta.SetSize(size)
	newFile.Seek(0, 0)
	newfileMeta.SetSha1(util.FileSha1(newFile))

	// 更新 fileMetaDict
	file.UpdateFileMetaDict(newfileMeta)

	// 上传成功 页面跳转
	// 根据当前路由 重定向
	currentRoute := r.URL.Path
	http.Redirect(w, r, currentRoute+"/success", http.StatusFound)
}
