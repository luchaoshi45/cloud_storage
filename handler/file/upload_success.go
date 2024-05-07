package file

import (
	"io"
	"log"
	"net/http"
)

// UploadSuccess 上传文件成功
type UploadSuccess struct {
}

func (us *UploadSuccess) Handler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "uploadSuccess")
	if err != nil {
		log.Println("io.WriteString(w, \"uploadSuccess\") ", err.Error())
	}
}
