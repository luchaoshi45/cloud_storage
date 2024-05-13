package user

import (
	"cloud_storage/db/mysql"
	"cloud_storage/handler"
	"net/http"
)

// Info ： 查询用户信息
func (su *User) Info(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	token := r.Form.Get("token")

	// 2. 验证token是否有效
	isValidToken := IsTokenValid(token)
	if !isValidToken {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// 3. 查询用户信息
	user := mysql.User{}
	user, err := user.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// 4. 组装并且响应用户数据
	resp := handler.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

// IsTokenValid : token是否有效
func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	// TODO: 判断token的时效性，是否过期
	// TODO: 从数据库表tbl_user_token查询username对应的token信息
	// TODO: 对比两个token是否一致
	return true
}
