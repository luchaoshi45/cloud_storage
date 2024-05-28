package main

import (
	"bufio"
	"cloud_storage/configurator"
	"cloud_storage/db/mysql"
	"cloud_storage/db/oss"
	"cloud_storage/rabbitmq"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func ProcessTransfer(msg []byte) bool {
	log.Println(string(msg))

	pubData := rabbitmq.TransferData{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	fin, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	err = oss.Bucket().PutObject(
		pubData.DestLocation,
		bufio.NewReader(fin))
	if err != nil {
		log.Println(err.Error())
		return false
	}

	tFile := mysql.NewFile()
	dir := strings.Split(pubData.DestLocation, "/")[0]
	_, err = tFile.UpdateDir(pubData.FileHash, dir+"/")
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func main() {
	log.Println("文件转移服务启动，开始监听转移任务队列...")
	cfg := configurator.GetRabbitMQConfig()
	rabbitmq.StartConsume(
		cfg.GetAttr("OssQueue"),
		"transfer_oss",
		ProcessTransfer)
}
