package router

import (
	"cloud_storage/handler"
	"net/http"
)

func Router() {
	uploadHandler := handler.NewUploadHandler()
	http.HandleFunc("/file/upload", uploadHandler.Handler)
}
