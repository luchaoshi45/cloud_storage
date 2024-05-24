package configurator

import (
	"encoding/json"
	"os"
)

type CephConfig struct {
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

func (cc *CephConfig) Parse(location string) (Config, error) {
	file, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}
	// 解析 JSON 数据
	err = json.Unmarshal(file, cc)
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (cc *CephConfig) GetAttr(attr string) string {
	switch attr {
	case "Endpoint":
		return cc.Endpoint
	case "AccessKey":
		return cc.AccessKey
	case "SecretKey":
		return cc.SecretKey
	default:
		return "" // 或者返回错误信息
	}
}
