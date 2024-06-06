package handler

import (
	"cloud_storage/handler"
	"cloud_storage/service/account/proto"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
	"net/http"
)

var (
	userCli proto.UserService
)

func init() {
	// 创建一个新的 Consul 注册中心
	consulRegistry := consul.NewRegistry(
		registry.Addrs("localhost:8500"), // 指定 Consul 地址和端口
	)
	// 创建服务
	service := micro.NewService(
		micro.Registry(consulRegistry),
	)
	service.Init()

	userCli = proto.NewUserService("go.micro.service.user", service.Client())
}

// SignUpGet 响应注册页面
func SignUpGet(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signup.html")
}

// SignUpPost 处理注册请求
func SignUpPost(c *gin.Context) {
	username := c.Request.FormValue("username")
	passwd := c.Request.FormValue("password")

	resp, err := userCli.Signup(context.TODO(), &proto.ReqSignup{
		Username: username,
		Password: passwd,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Message,
	})
}

// SignIndGet 响应登录页面
func SignIndGet(c *gin.Context) {
	c.Redirect(http.StatusFound, "/static/view/signin.html")
}

// SignInPost : 处理登录请求
func SignInPost(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	rpcResp, err := userCli.Signin(context.TODO(), &proto.ReqSignin{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if rpcResp.Code != 0 {
		c.JSON(200, gin.H{
			"msg":  "登录失败",
			"code": rpcResp.Code,
		})
		return
	}

	// 登录成功，返回用户信息
	resp := RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token    string
		}{
			Location: "/static/view/home.html",
			Username: username,
			Token:    rpcResp.Token,
		},
	}

	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
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

// Authorize : http请求拦截器
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		token := c.Request.FormValue("token")

		//验证登录 token 是否有效
		if len(username) < 3 || !IsTokenValid(token) {
			// w.WriteHeader(http.StatusForbidden)
			// token校验失败则跳转到登录页面
			c.Abort()

			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token 无效",
			})
			return
		}
		c.Next()
	}
}

// UserInfo 查询用户信息
func UserInfo(c *gin.Context) {
	// 1. 解析请求参数
	username := c.Request.FormValue("username")

	resp, err := userCli.UserInfo(context.TODO(), &proto.ReqUserInfo{
		Username: username,
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	// 3. 组装并且响应用户数据
	cliResp := handler.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: gin.H{
			"user_name": username,
			"signup_at": resp.SignupAt,
			// TODO: 完善其他字段信息
			"last_active": resp.LastActiveAt,
		},
	}
	c.Data(http.StatusOK, "application/json", cliResp.JSONBytes())
}
