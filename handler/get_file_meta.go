package handler

import (
	"cloud_storage/file"
	"encoding/json"
	"net/http"
)

type GetFileMetaHandler struct {
}

func (gfmh *GetFileMetaHandler) Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	hash := r.Form["sha1"][0]
	fileMeta := file.GetFileMeta(hash)
	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
