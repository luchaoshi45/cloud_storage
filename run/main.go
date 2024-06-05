package main

import (
	"cloud_storage/router"
)

func main() {
	r := router.Router()
	_ = r.Run(":42200")
}
