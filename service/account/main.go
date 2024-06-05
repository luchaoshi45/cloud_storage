package main

import (
	"cloud_storage/service/account/handler"
	"cloud_storage/service/account/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
)

func main() {

	// 创建一个新的 Consul 注册中心
	consulRegistry := consul.NewRegistry(
		registry.Addrs("localhost:8500"), // 指定 Consul 地址和端口
	)

	// 创建服务
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Registry(consulRegistry),
	)

	service.Init()

	// 注册服务处理器
	if err := proto.RegisterUserServiceHandler(service.Server(), new(handler.User)); err != nil {
		log.Fatalf("Failed to register handler: %v", err)
	}

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatalf("Failed to run service: %v", err)
	}
}
