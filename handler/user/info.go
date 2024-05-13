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
