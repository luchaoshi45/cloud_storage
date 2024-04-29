package handler

import (
	"cloud_storage/file"
	"encoding/json"
	"net/http"
)

type UpdateFileMetaHandler struct {
}

func (ufmh *UpdateFileMetaHandler) Handler(w http.ResponseWriter, r *http.Request) {
	// POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// ParseForm
	r.ParseForm()
	opType := r.Form.Get("op")
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	sha1 := r.Form.Get("sha1")
	newFileName := r.Form.Get("name")

	// 更新
	fileMeta := file.GetFileMeta(sha1)
	// 记住 Location
	oldLocation := fileMeta.GetLocation()
	// 更新 Location
	fileMeta.SetName(newFileName)
	// 更新 FileMetaDict
	file.UpdateFileMetaDict(fileMeta)
	// Rename File
	err := file.SafeRename(oldLocation, fileMeta.GetLocation())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write
	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
