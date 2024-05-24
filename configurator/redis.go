package configurator

import (
	"encoding/json"
	"fmt"
	"os"
)

type RedisConfig struct {
	config   string
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

var redisConfig *RedisConfig

func init() {
	redisConfig = &RedisConfig{config: "config/redis/conn.json"}
	err := redisConfig.parse()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func GetredisConfig() Config {
	return redisConfig
}

func (rc *RedisConfig) parse() error {
	file, err := os.ReadFile(rc.config)
	if err != nil {
		return err
	}
	// 解析 JSON 数据
	err = json.Unmarshal(file, rc)
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisConfig) GetAttr(attr string) string {
	switch attr {
	case "IP":
		return rc.IP
	case "Port":
		return rc.Port
	case "Password":
		return rc.Password
	case "Host":
		return rc.IP + ":" + rc.Port
	default:
		return "" // 或者返回错误信息
	}
}
