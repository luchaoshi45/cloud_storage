package file

import (
	"io"
	"log"
	"net/http"
)

func (f *File) UploadDuplicate(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "uploadDuplicate")
	if err != nil {
		log.Println("io.WriteString(w, \"uploadDuplicate\") ", err.Error())
	}
}
