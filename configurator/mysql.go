package configurator

import (
	"encoding/json"
	"os"
)

type MysqlConfig struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	IP           string `json:"ip"`
	Port         string `json:"port"`
	DatabaseName string `json:"databaseName"`
}

func (mc *MysqlConfig) Parse(location string) (Config, error) {
	file, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}
	// 解析 JSON 数据
	err = json.Unmarshal(file, mc)
	if err != nil {
		return nil, err
	}
	return mc, nil
}

func (mc *MysqlConfig) GetAttr(attr string) string {
	switch attr {
	case "User":
		return mc.User
	case "Password":
		return mc.Password
	case "IP":
		return mc.IP
	case "Port":
		return mc.Port
	case "DatabaseName":
		return mc.DatabaseName
	default:
		return "" // 或者返回错误信息
	}
}
