package router

import (
	"cloud_storage/handler/user"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	u := user.NewUser()
	//f := file.NewFile()
	//uf := user_file.NewUserFile()

	// gin framework 包括Logger, Recovery
	router := gin.Default()

	// 1 静态文件
	router.Static("/static/", "./static")
	// 2 不需要拦截器验证就能访问的接口
	router.GET("/user/signup", u.SignUpGet)
	router.POST("/user/signup", u.SignUpPost)
	router.GET("/user/signin", u.SignIndGet)
	router.POST("/user/signin", u.SignInPost)
	// 拦截器 下面的都需要 token 校验
	router.Use(HttpInterceptor())
	return router
}

//package router
//
//import (
//	"cloud_storage/handler/file"
//	"cloud_storage/handler/user"
//	"cloud_storage/handler/user_file"
//	"net/http"
//	"regexp"
//	"sync"
//)
//
//// handlerFunc
//type handlerFunc func(http.ResponseWriter, *http.Request)
//
//// routerDict 路由字典
//var routerDict map[string]handlerFunc
//var once sync.Once
//
//// GetRouterDict 获得 RouterDict 路由字典
//func GetRouterDict() map[string]handlerFunc {
//	once.Do(func() {
//		routerDict = make(map[string]handlerFunc)
//	})
//	return routerDict
//}
//
//// Router 初始化
//func Router() {
//	GetRouterDict()
//	// 设置静态文件路径
//	fs := http.FileServer(http.Dir("./static"))
//	http.Handle("/static/", http.StripPrefix("/static/", fs))
//
//	u := user.NewUser()
//	f := file.NewFile()
//	uf := user_file.NewUserFile()
//
//	addEntry("/file/upload", f.Upload)
//	//addEntry("/static/view/", f.Upload)
//	addEntry("/file/upload/success", f.UploadSuccess)
//	addEntry("/file/upload/duplicate", f.UploadDuplicate)
//	addEntry("/file/scan", f.GetFileMeta)
//	addEntry("/file/download", f.Download)
//	addEntry("/file/update/name", f.UpdateFileMeta)
//	addEntry("/file/delete", f.Delete)
//	addEntry("/file/404", f.FileNotFound)
//
//	// 分块上传接口
//	addEntry("/file/mpupload/init", file.InitMultipartUpload)
//	addEntry("/file/mpupload/uppart", file.UploadPart)
//	addEntry("/file/mpupload/complete", file.CompleteUpload)
//
//	addEntry("/user/signup", u.SignUp)
//	addEntry("/user/signin", u.SignIn)
//	addEntry("/user/info", u.Info)
//
//	addEntry("/user_file/query", uf.FileQuery)
//	config()
//}
//
//// 向 routerDict 中添加新的条目
//func addEntry(key string, value handlerFunc) {
//	routerDict[key] = value
//}
//
//// 配置路由
//func config() {
//	for k, v := range routerDict {
//		// 编译正则表达式
//		re := regexp.MustCompile(`^/user/sign`)
//
//		// 判断字符串是否匹配正则表达式
//		if re.MatchString(k) {
//			http.HandleFunc(k, v)
//		} else {
//			re := regexp.MustCompile(`^/file/`)
//			if re.MatchString(k) {
//				http.HandleFunc(k, v)
//			} else {
//				http.HandleFunc(k, HttpInterceptor(http.HandlerFunc(v)))
//			}
//		}
//
//	}
//}
