package file

import (
	"cloud_storage/db/mysql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (f *File) GetFileMeta(c *gin.Context) {
	sha1 := c.Request.Form["sha1"][0]
	userFile, err := mysql.NewFile().Query(sha1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "mysql.NewFile().Query(sha1)",
		})
		return
	}

	// 转化 fileMeta 为 json
	data, err := json.Marshal(userFile)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "json.Marshal(userFile)",
		})
		return
	}
	// 返回
	c.Data(http.StatusOK, "application/json", data)
}
