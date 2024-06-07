package handler

import (
	uploadProto "cloud_storage/service/upload/proto"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
	"net/http"
)

var (
	uploadCil uploadProto.UploadService
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

	uploadCil = uploadProto.NewUploadService("go.micro.service.upload", service.Client())
}

// UploadEntry : 查询批量的文件元信息
func UploadEntry(c *gin.Context) {
	rpcResp, err := uploadCil.UploadEntry(context.TODO(), &uploadProto.ReqEntry{})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  rpcResp.Code,
		"entry": rpcResp.Entry,
		"msg":   "ok",
	})
}
