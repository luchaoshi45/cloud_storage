package main

import (
	upProto "cloud_storage/service/upload/proto"
	"cloud_storage/service/upload/router"
	upRpc "cloud_storage/service/upload/rpc"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"time"
)

func startRPCService() {
	// 创建一个新的 Consul 注册中心
	consulRegistry := consul.NewRegistry(
		registry.Addrs("localhost:8500"), // 指定 Consul 地址和端口
	)
	service := micro.NewService(
		micro.Name("go.micro.service.upload"), // 服务名称
		micro.RegisterTTL(time.Second*10),     // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5), // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
		micro.Registry(consulRegistry),
	)
	service.Init()

	err := upProto.RegisterUploadServiceHandler(service.Server(), new(upRpc.Upload))
	if err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

func startAPIService() {
	router := router.Router()
	err := router.Run(upRpc.UploadServiceHost)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// rpc 服务
	go startRPCService()

	// api 服务
	startAPIService()
}
