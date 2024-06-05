package user

import (
	"cloud_storage/db/mysql"
	"cloud_storage/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Info ： 查询用户信息
func (su *User) Info(c *gin.Context) {
	// 1. 解析请求参数
	username := c.Request.FormValue("username")

	// 3. 查询用户信息
	user := mysql.User{}
	user, err := user.GetUserInfo(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "GetUserInfo failed",
			"code": -1,
		})
		return
	}

	// 4. 组装并且响应用户数据
	resp := handler.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
}
