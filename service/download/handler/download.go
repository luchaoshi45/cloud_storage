package handler

import (
	"cloud_storage/db/mysql"
	"cloud_storage/db/oss"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(c *gin.Context) {
	// 获取 SHA1 值
	sha1 := c.Query("sha1")
	// 得到 userFile
	userFile, err := mysql.NewFile().Query(sha1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "mysql.NewFile().Query(sha1)",
		})
		return
	}

	// 得到 文件位置
	location := userFile.GetLocation()
	dir := strings.Split(location, "/")[0]

	if dir == "oss" {
		signedURL := oss.DownloadURL("oss/" + sha1)
		// 如果文件存储在 OSS 中，生成带签名的 URL，并返回给客户端
		c.String(http.StatusOK, signedURL)
		//w.Header().Set("Content-Type", "application/text/plain")
		//w.Write([]byte(signedURL))
	} else if dir == "tmp" {
		file, err := os.Open(location)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -2,
				"message": "os.Open(location)",
			})
			return
		}
		defer file.Close()

		// 读取文件到内存
		data, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    -2,
				"message": "io.ReadAll(file)",
			})
			return
		}

		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename=\""+userFile.Name+"\"")
		c.Data(http.StatusOK, "application/octet-stream", data)
	}
}
