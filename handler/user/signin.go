package user

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"cloud_storage/handler"
	"fmt"
	"net/http"
	"time"
)

// SignIn : 登录接口
func (su *User) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	encPasswd := file.Sha1([]byte(password + pwdSalt))

	// 1. 校验用户名及密码
	pwdChecked := mysql.NewUser(username, encPasswd).SignIn()
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := mysql.NewUserToken(username, token).UpdateToken()
	if !upRes {
		w.Write([]byte("FAILED"))
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
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: username,
			Token:    token,
		},
	}
	w.Write(resp.JSONBytes())
}

// GenToken : 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := file.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}
