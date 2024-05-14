package file

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"net/http"
)

func (f *File) Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sha1 := r.Form.Get("sha1")

	// 得到 userFile
	userFile := mysql.NewFile()
	_, err := userFile.Query(sha1)
	if err != nil {
		http.Redirect(w, r, "/file/404", http.StatusFound)
		return
	}

	err = file.SafeRemove(userFile.GetLocation())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = userFile.Delete(sha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
