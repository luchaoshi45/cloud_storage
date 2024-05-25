package file

import (
	"cloud_storage/db/mysql"
	"cloud_storage/db/oss"
	"io"
	"net/http"
	"os"
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
	file, err := os.Open(userFile.GetLocation())
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
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+userFile.Name+"\"")
	w.Write(data)
}

// DownloadURL : 生成文件的下载地址
func (f *File) DownloadURL(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("sha1")
	// 从文件表查找记录
	//row, _ := mysql.NewFile().Query(filehash)
	//ossPath := row.GetLocation()
	ossPath := "oss/" + filehash
	signedURL := oss.DownloadURL(ossPath)

	w.Write([]byte(signedURL))
}
