package configurator

import (
	"encoding/json"
	"fmt"
	"os"
)

type MysqlConfig struct {
	config       string
	User         string `json:"user"`
	Password     string `json:"password"`
	IP           string `json:"ip"`
	Port         string `json:"port"`
	DatabaseName string `json:"databaseName"`
}

var mysqlConfig *MysqlConfig

func init() {
	mysqlConfig = &MysqlConfig{config: "config/mysql/master_conn.json"}
	err := mysqlConfig.parse()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func GetMysqlConfig() Config {
	return mysqlConfig
}

func (mc *MysqlConfig) parse() error {
	file, err := os.ReadFile(mc.config)
	if err != nil {
		return err
	}
	// 解析 JSON 数据
	err = json.Unmarshal(file, mc)
	if err != nil {
		return err
	}
	return nil
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
