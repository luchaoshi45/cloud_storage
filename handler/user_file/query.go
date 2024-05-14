package user_file

import (
	"cloud_storage/db/mysql"
	"encoding/json"
	"net/http"
	"strconv"
)

// FileQuery : 查询批量的文件元信息
func (uf *UserFile) FileQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
	username := r.Form.Get("username")
	//fileMetas, _ := meta.GetLastFileMetasDB(limitCnt)
	userFile := mysql.NewUserFile()
	userFiles, err := userFile.QueryUserFileMetas(username, limitCnt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(userFiles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
