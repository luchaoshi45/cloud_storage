package file

import (
	"cloud_storage/db/mysql"
	"cloud_storage/db/oss"
	"io"
	"net/http"
	"os"
	"strings"
)

func (f *File) Download(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sha1 := r.Form["sha1"][0]

	// 得到 userFile
	userFile, err := mysql.NewFile().Query(sha1)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 得到 文件位置
	location := userFile.GetLocation()
	dir := strings.Split(location, "/")[0]

	if dir == "oss" {
		signedURL := oss.DownloadURL("oss/" + sha1)
		w.Header().Set("Content-Type", "application/text/plain")
		w.Write([]byte(signedURL))
	} else if dir == "tmp" {
		file, err := os.Open(location)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// 读取文件到内存
		data, err := io.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// 返回
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=\""+userFile.Name+"\"")
		w.Write(data)
	}
}
