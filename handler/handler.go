package handler

import "net/http"

type Handler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

// NewUploadHandler  New UploadHandler 结构体
func NewUploadHandler() Handler {
	return &UploadHandler{}
}

// NewUploadSuccessHandler  New UploadSuccessHandler 结构体
func NewUploadSuccessHandler() Handler {
	return &uploadSuccess{}
}

// NewGetFileMetaHandler New GetFileMetaHandler 结构体
func NewGetFileMetaHandler() Handler {
	return &GetFileMetaHandler{}
}
