package main

import (
	"cloud_storage/db/mysql"
	"cloud_storage/router"
	"log"
	"net/http"
)

func sysInit() {
	// mysql 数据库
	mysql.MySqlConn("config/mysql/master_conn.json")

	// 配置路由
	router.Router()
	// 设置监听的端口
	err := http.ListenAndServe(":42200", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	sysInit()
}
