package handler

import (
	"io"
	"log"
	"net/http"
)

// uploadSuccess 上传文件成功
type uploadSuccess struct {
}

func NewUploadSuccessHandler() Handler {
	return &uploadSuccess{}
}

func (us *uploadSuccess) Handler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "uploadSuccess")
	if err != nil {
		log.Println("io.WriteString(w, \"uploadSuccess\") ", err.Error())
	}
}
