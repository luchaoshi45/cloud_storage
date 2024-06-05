package main

import "cloud_storage/service/apigw/router"

func main() {
	r := router.Router()
	_ = r.Run(":40001")
}
