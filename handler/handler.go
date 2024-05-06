package handler

import "net/http"

type Handler interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

func NewUploadHandler() Handler {
	return &UploadHandler{}
}

func NewUploadSuccessHandler() Handler {
	return &UploadSuccess{}
}

func NewUploadDuplicateHandler() Handler {
	return &UploadDuplicate{}
}

func NewGetFileMetaHandler() Handler {
	return &GetFileMetaHandler{}
}

func NewDownloadHandler() Handler {
	return &DownloadHandler{}
}

func NewUpdateFileMetaHandler() Handler {
	return &UpdateFileMetaHandler{}
}

func NewDeleteHandler() Handler {
	return &DeleteHandler{}
}

func NewFileNotFound() Handler {
	return &FileNotFound{}
}
