package file

import (
	"io"
	"log"
	"net/http"
)

func (f *File) UploadSuccess(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "uploadSuccess")
	if err != nil {
		log.Println("io.WriteString(w, \"uploadSuccess\") ", err.Error())
	}
}
