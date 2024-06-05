package handler

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"cloud_storage/service/account/proto"
	"context"
)

type User struct {
}

const (
	// 用于加密的盐值(自定义)
	pwdSalt    = "*#890"
	minPwdLen  = 1
	minNameLen = 1
)

func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, resp *proto.RespSignup) error {
	username := req.Username
	passwd := req.Password

	if len(username) < minNameLen || len(passwd) < minPwdLen {
		resp.Code = -1
		resp.Mesaage = "Invalid parameter"
		return nil
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := file.Sha1([]byte(passwd + pwdSalt))

	// 将用户信息注册到用户表中
	suc := mysql.NewUser(username, encPasswd).SignUp()
	if !suc {
		resp.Code = -2
		resp.Mesaage = "SignUp failed"
		return nil
	}

	resp.Code = 0
	resp.Mesaage = "SignUp succeeded"
	return nil
}
