package file

import (
	"io"
	"log"
	"net/http"
)

func (f *File) FileNotFound(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "fileNotFound")
	if err != nil {
		log.Println("io.WriteString(w, \"fileNotFound\") ", err.Error())
	}
}
