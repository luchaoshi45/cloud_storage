package handler

import (
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"cloud_storage/service/account/proto"
	"context"
	"fmt"
	"time"
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
		resp.Message = "Invalid parameter"
		return nil
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := file.Sha1([]byte(passwd + pwdSalt))

	// 将用户信息注册到用户表中
	suc := mysql.NewUser(username, encPasswd).SignUp()
	if !suc {
		resp.Code = -2
		resp.Message = "SignUp failed"
		return nil
	}

	resp.Code = 0
	resp.Message = "ok"
	return nil
}

// Signin : 处理登录请求
func (u *User) Signin(ctx context.Context, req *proto.ReqSignin, resp *proto.RespSignin) error {
	username := req.Username
	password := req.Password

	encPasswd := file.Sha1([]byte(password + pwdSalt))

	// 1. 校验用户名及密码
	pwdChecked := mysql.NewUser(username, encPasswd).SignIn()
	if !pwdChecked {
		resp.Code = -1
		resp.Message = "mysql.NewUser(username, encPasswd).SignIn()"
		return nil
	}

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := mysql.NewUserToken(username, token).UpdateToken()
	if !upRes {
		resp.Code = -2
		resp.Message = "mysql.NewUserToken(username, token).UpdateToken()"
		return nil
	}

	// 3. 登录成功, 返回token
	resp.Code = 0
	resp.Message = "ok"
	resp.Token = token
	return nil
}

// GenToken : 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := file.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// UserInfo ： 查询用户信息
func (u *User) UserInfo(ctx context.Context, req *proto.ReqUserInfo, resp *proto.RespUserInfo) error {
	// 1. 解析请求参数
	username := req.Username

	// 3. 查询用户信息
	user := mysql.User{}
	user, err := user.GetUserInfo(username)
	if err != nil {
		resp.Code = -1
		resp.Message = "user.GetUserInfo(username)"
		return nil
	}

	// 3. 组装并且响应用户数据
	resp.Code = 0
	resp.Username = user.UserName
	resp.SignupAt = user.SignupAt
	resp.LastActiveAt = user.LastActive
	resp.Status = int32(user.Status)

	// TODO: 需增加接口支持完善用户信息(email/phone等)
	resp.Email = user.Email
	resp.Phone = user.Phone

	return nil
}
