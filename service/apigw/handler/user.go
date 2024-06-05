package handler

import (
	"cloud_storage/service/account/proto"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
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
		"msg":  resp.Mesaage,
	})

}
