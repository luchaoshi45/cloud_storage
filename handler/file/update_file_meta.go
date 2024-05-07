package file

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"encoding/json"
	"net/http"
)

type UpdateFileMeta struct {
}

func (ufmh *UpdateFileMeta) Handler(w http.ResponseWriter, r *http.Request) {
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

	// 得到 userFile
	userFile, err := mysql.NewUserFile().Query(sha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 记住 Location
	oldLocation := userFile.GetLocation()
	// 更新 userFile 和数据库中的 Name
	userFile, err = userFile.Update(sha1, newFileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Rename File
	err = file.SafeRename(oldLocation, userFile.GetLocation())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write
	data, err := json.Marshal(userFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
