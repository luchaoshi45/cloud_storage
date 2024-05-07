package file

import (
	"cloud_storage/db/mysql"
	"encoding/json"
	"net/http"
)

func (f *File) GetFileMeta(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sha1 := r.Form["sha1"][0]
	userFile, err := mysql.NewUserFile().Query(sha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 转化 fileMeta 为 json
	data, err := json.Marshal(userFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 返回
	w.Write(data)
}
