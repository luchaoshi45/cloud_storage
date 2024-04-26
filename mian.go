package main

import (
	"cloud_storage/router"
	"log"
	"net/http"
)

func main() {
	// 配置路由
	router.Router()

	// 设置监听的端口
	err := http.ListenAndServe(":42200", nil)
	if err != nil {
		log.Fatal(err)
	}

}
