package handler

import (
	"cloud_storage/file"
	"io"
	"net/http"
	"os"
)

type DownloadHandler struct {
}

func (dh *DownloadHandler) Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sha1 := r.Form["sha1"][0]
	fileMeta := file.GetFileMeta(sha1)
	f, err := os.Open(fileMeta.GetLocation())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileMeta.GetName()+"\"")
	w.Write(data)
}
