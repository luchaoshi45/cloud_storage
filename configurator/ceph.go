package configurator

import (
	"encoding/json"
	"fmt"
	"os"
)

type CephConfig struct {
	config    string
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

var cephConfig *CephConfig

func init() {
	cephConfig = &CephConfig{config: "config/ceph/conn.json"}
	err := cephConfig.parse()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func GetCephConfig() Config {
	return cephConfig
}

func (cc *CephConfig) parse() error {
	file, err := os.ReadFile(cc.config)
	if err != nil {
		return err
	}
	// 解析 JSON 数据
	err = json.Unmarshal(file, cc)
	if err != nil {
		return err
	}
	return nil
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
