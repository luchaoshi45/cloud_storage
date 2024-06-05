package router

import (
	"cloud_storage/service/apigw/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/static/", "./static")

	router.GET("/user/signup", handler.SignUpGet)
	router.POST("/user/signup", handler.SignUpPost)
	return router
}
