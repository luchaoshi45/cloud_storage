package user

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"net/http"
)

func (su *User) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//data, err := os.ReadFile("./static/view/signup.html")
		//if err != nil {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//w.Write(data)
		http.Redirect(w, r, "/static/view/signup.html", http.StatusFound)
		return
	}
	r.ParseForm()

	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < minNameLen || len(passwd) < minPwdLen {
		w.Write([]byte("Invalid parameter"))
		return
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := file.Sha1([]byte(passwd + pwdSalt))
	// 将用户信息注册到用户表中
	suc := mysql.NewUser(username, encPasswd).SignUp()
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}
