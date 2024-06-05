package user

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignUpGet 响应注册页面
func (su *User) SignUpGet(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// SignUpPost 处理注册请求
func (su *User) SignUpPost(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	if len(username) < minNameLen || len(passwd) < minPwdLen {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Invalid parameter",
			"code": -1,
		})
		return
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := file.Sha1([]byte(passwd + pwdSalt))
	// 将用户信息注册到用户表中
	suc := mysql.NewUser(username, encPasswd).SignUp()
	if suc {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "SignUp succeeded",
			"code": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "SignUp failed",
			"code": -2,
		})
	}
}
