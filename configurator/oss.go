package configurator

import (
	"encoding/json"
	"fmt"
	"os"
)

type OssConfig struct {
	config          string
	Bucket          string `json:"bucket"`
	Endpoint        string `json:"endpoint"`
	AccessKey       string `json:"access_key"`
	AccessKeySecret string `json:"access_key_secret"`
}

var ossConfig *OssConfig

func init() {
	ossConfig = &OssConfig{config: "config/oss/conn.json"}
	err := ossConfig.parse()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func GetOssConfig() Config {
	return ossConfig
}

func (oc *OssConfig) parse() error {
	file, err := os.ReadFile(oc.config)
	if err != nil {
		return err
	}
	// 解析 JSON 数据
	err = json.Unmarshal(file, oc)
	if err != nil {
		return err
	}
	return nil
}

func (oc *OssConfig) GetAttr(attr string) string {
	switch attr {
	case "Bucket":
		return oc.Bucket
	case "Endpoint":
		return oc.Endpoint
	case "AccessKey":
		return oc.AccessKey
	case "AccessKeySecret":
		return oc.AccessKeySecret
	default:
		return "" // 或者返回错误信息
	}
}
