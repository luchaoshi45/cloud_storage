package rpc

import (
	DnProto "cloud_storage/service/download/proto"
	"context"
	"fmt"
	"net"
)

// DownloadEntry : 配置上传入口地址
var DownloadEntry = ":40003"

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

// DownloadServiceHost : 上传服务监听的地址
var DownloadServiceHost = "0.0.0.0:40003"

// Download : upload结构体
type Download struct{}

// DownloadEntry : 获取上传入口
func (d *Download) DownloadEntry(
	ctx context.Context,
	req *DnProto.ReqEntry,
	res *DnProto.RespEntry) error {
	ip, _ := getLocalIP()
	res.Entry = ip + DownloadEntry
	return nil
}
