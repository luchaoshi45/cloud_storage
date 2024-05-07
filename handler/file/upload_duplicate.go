package file

import (
	"io"
	"log"
	"net/http"
)

// UploadDuplicate 上传文件重复
type UploadDuplicate struct {
}

func (us *UploadDuplicate) Handler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "uploadDuplicate")
	if err != nil {
		log.Println("io.WriteString(w, \"uploadDuplicate\") ", err.Error())
	}
}
