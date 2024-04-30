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
	sha1 := r.Form["sha1"][0]

	// 得到 fileMeta
	fileMeta, err := file.GetFileMeta(sha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 转化 fileMeta 为 json
	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 返回
	w.Write(data)
}
