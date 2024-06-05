package main

import (
	"cloud_storage/router"
	"fmt"
)

func sysInit() {
	// 配置路由
	r := router.Router()
	err := r.Run(":42200")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	sysInit()
}
