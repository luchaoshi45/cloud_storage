package file

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (f *File) Delete(c *gin.Context) {
	// 获取 SHA1 值
	sha1 := c.Request.Form.Get("sha1")

	// 得到 userFile
	userFile := mysql.NewFile()
	_, err := userFile.Query(sha1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "File not found",
		})
		return
	}

	err = file.SafeRemove(userFile.GetLocation())
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -2,
			"message": "SafeRemove err",
		})
		return
	}
	_, err = userFile.Delete(sha1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -3,
			"message": "userFile Delete err",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}
