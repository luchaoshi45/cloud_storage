package user

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"cloud_storage/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SignIndGet 响应登录页面
func (su *User) SignIndGet(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// SignInPost : 处理登录请求
func (su *User) SignInPost(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	encPasswd := file.Sha1([]byte(password + pwdSalt))

	// 1. 校验用户名及密码
	pwdChecked := mysql.NewUser(username, encPasswd).SignIn()
	if !pwdChecked {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Login failed",
			"code": -1,
		})
		return
	}

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := mysql.NewUserToken(username, token).UpdateToken()
	if !upRes {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "GenToken failed",
			"code": -2,
		})
		return
	}

	// 3. 登录成功后重定向到首页
	resp := handler.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}

	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
}

// GenToken : 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := file.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
