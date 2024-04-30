package handler

import (
	"cloud_storage/file"
	"net/http"
)

type DeleteHandler struct {
}

func (dh *DeleteHandler) Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sha1 := r.Form.Get("sha1")

	fileMeta, err := file.GetFileMeta(sha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = file.SafeRemove(fileMeta.GetLocation())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = file.RemoveFileMeta(sha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
