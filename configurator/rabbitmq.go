package configurator

import (
	"encoding/json"
	"fmt"
	"os"
)

type RabbitMQConfig struct {
	config      string
	Url         string `json:"url"`
	Exchange    string `json:"exchange"`
	OssQueue    string `json:"oss_queue"`
	OssErrQueue string `json:"oss_err_queue"`
	RoutingKey  string `json:"routing_key"`
}

var rabbitMQConfig *RabbitMQConfig

func init() {
	rabbitMQConfig = &RabbitMQConfig{config: "config/rabbitmq/conn.json"}
	err := rabbitMQConfig.parse()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func GetRabbitMQConfig() Config {
	return rabbitMQConfig
}

func (rmqc *RabbitMQConfig) parse() error {
	file, err := os.ReadFile(rmqc.config)
	if err != nil {
		return err
	}
	// 解析 JSON 数据
	err = json.Unmarshal(file, rmqc)
	if err != nil {
		return err
	}
	return nil
}

func (rmqc *RabbitMQConfig) GetAttr(attr string) string {
	switch attr {
	case "Url":
		return rmqc.Url
	case "Exchange":
		return rmqc.Exchange
	case "OssQueue":
		return rmqc.OssQueue
	case "OssErrQueue":
		return rmqc.OssErrQueue
	case "RoutingKey":
		return rmqc.RoutingKey
	default:
		return "" // 或者返回错误信息
	}
}
