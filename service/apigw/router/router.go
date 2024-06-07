package router

import (
	"cloud_storage/service/apigw/handler"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Static("/static/", "./static")
	// 注册
	router.GET("/user/signup", handler.SignUpGet)
	router.POST("/user/signup", handler.SignUpPost)
	// 登录
	router.GET("/user/signin", handler.SignIndGet)
	router.POST("/user/signin", handler.SignInPost)

	// 拦截器
	router.Use(handler.Authorize())

	// 用户查询
	router.POST("/user/info", handler.UserInfo)
	// 用户文件查询
	router.POST("/user_file/query", handler.UserFiles)

	// 返回上传接口地址
	router.GET("/get/upload/entry", handler.UploadEntry)
	router.GET("/get/download/entry", handler.DownloadEntry)

	return router
}
