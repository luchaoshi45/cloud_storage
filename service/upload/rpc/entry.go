package rpc

import (
	upProto "cloud_storage/service/upload/proto"
	"context"
	"fmt"
	"net"
)

// UploadEntry : 配置上传入口地址
var UploadEntry = ":40002"

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// 检查地址类型并跳过回环地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no suitable IP address found")
}

// UploadServiceHost : 上传服务监听的地址
var UploadServiceHost = "0.0.0.0:40002"

// Upload : upload结构体
type Upload struct{}

// UploadEntry : 获取上传入口
func (u *Upload) UploadEntry(
	ctx context.Context,
	req *upProto.ReqEntry,
	res *upProto.RespEntry) error {
	ip, _ := getLocalIP()
	res.Entry = ip + UploadEntry
	return nil
}
