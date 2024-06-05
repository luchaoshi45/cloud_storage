package user_file

import (
	"cloud_storage/db/mysql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FileQuery : 查询批量的文件元信息
func (uf *UserFile) FileQuery(c *gin.Context) {
	limitCnt, _ := strconv.Atoi(c.Request.FormValue("limit"))
	username := c.Request.FormValue("username")
	//fileMetas, _ := meta.GetLastFileMetasDB(limitCnt)
	userFile := mysql.NewUserFile()
	userFiles, err := userFile.QueryUserFileMetas(username, limitCnt)
	//fmt.Println(userFiles)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "QueryUserFileMetas failed",
			"code": -1,
		})
		return
	}
	juserFiles, _ := json.Marshal(userFiles)
	c.Data(http.StatusOK, "application/json", juserFiles)
}
