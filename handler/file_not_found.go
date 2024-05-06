package handler

import (
	"io"
	"log"
	"net/http"
)

// FileNotFound 文件找不到
type FileNotFound struct {
}

func (us *FileNotFound) Handler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "fileNotFound")
	if err != nil {
		log.Println("io.WriteString(w, \"fileNotFound\") ", err.Error())
	}
}
