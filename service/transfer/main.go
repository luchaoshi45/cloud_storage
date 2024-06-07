package main

//import (
//	"cloud_storage/configurator"
//	"cloud_storage/rabbitmq"
//	"fmt"
//	"go-micro.dev/v4"
//	"time"
//)
//
//func main() {
//	log.Println("文件转移服务启动，开始监听转移任务队列...")
//	cfg := configurator.GetRabbitMQConfig()
//	rabbitmq.StartConsume(
//		cfg.GetAttr("OssQueue"),
//		"transfer_oss",
//		ProcessTransfer)
//}

import (
	"cloud_storage/configurator"
	"cloud_storage/rabbitmq"
	"cloud_storage/service/transfer/handler"
	"fmt"
	"go-micro.dev/v4"
	"log"
	"time"
)

func startRPCService() {
	service := micro.NewService(
		micro.Name("go.micro.service.transfer"), // 服务名称
		micro.RegisterTTL(time.Second*10),       // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5),   // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
	)
	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startTranserService() {
	log.Println("文件转移服务启动，开始监听转移任务队列...")
	cfg := configurator.GetRabbitMQConfig()
	rabbitmq.StartConsume(
		cfg.GetAttr("OssQueue"),
		"transfer_oss",
		handler.ProcessTransfer)
}

func main() {
	// 文件转移服务
	go startTranserService()

	// rpc 服务
	startRPCService()
}
