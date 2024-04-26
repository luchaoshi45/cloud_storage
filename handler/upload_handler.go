package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type UploadHandler struct {
}

func NewUploadHandler() Handler {
	return &UploadHandler{}
}

func (uh *UploadHandler) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("UploadHandler Handler")
	//fmt.Fprintf(w, "UploadHandler Handler") //这个写入到w的是输出到客户端的

	if r.Method == "GET" {
		data, err := os.ReadFile("static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "internel server error")
			log.Println("os.ReadFile(\"static/view/index.html\") err")
		}
		_, err = io.WriteString(w, string(data))

		if err != nil {
			log.Println("io.WriteString(w, string(data)) err")
		}

	} else if r.Method == "POST" {

	}

}
