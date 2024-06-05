package file

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (f *File) UpdateFileMeta(c *gin.Context) {
	// ParseForm
	opType := c.Request.Form.Get("op")
	if opType != "0" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "Forbidden opType",
		})
		return
	}
	sha1 := c.Request.Form.Get("sha1")
	newFileName := c.Request.Form.Get("name")

	// 得到 userFile
	userFile, err := mysql.NewFile().Query(sha1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": "mysql.NewFile().Query(sha1)",
		})
		return
	}

	// 记住 Location
	oldLocation := userFile.GetLocation()
	// 更新 userFile 和数据库中的 Name
	userFile, err = userFile.Update(sha1, newFileName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -3,
			"message": "userFile.Update(sha1, newFileName)",
		})
		return
	}
	// Rename File
	err = file.SafeRename(oldLocation, userFile.GetLocation())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -4,
			"message": "file.SafeRename(oldLocation, userFile.GetLocation())",
		})
		return
	}

	// Write
	data, err := json.Marshal(userFile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -5,
			"message": "json.Marshal(userFile)",
		})
		return
	}
	c.Data(http.StatusOK, "application/json", data)
}
