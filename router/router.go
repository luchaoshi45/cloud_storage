package router

import (
	"cloud_storage/handler"
	"cloud_storage/handler/user"
	"net/http"
	"sync"
)

// handlerFunc
type handlerFunc func(http.ResponseWriter, *http.Request)

// routerDict 路由字典
var routerDict map[string]handlerFunc
var once sync.Once

// GetRouterDict 获得 RouterDict 路由字典
func GetRouterDict() map[string]handlerFunc {
	once.Do(func() {
		routerDict = make(map[string]handlerFunc)
	})
	return routerDict
}

// Router 初始化
func Router() {
	GetRouterDict()
	addEntry("/file/upload", handler.NewUploadHandler().Handler)
	addEntry("/file/upload/success", handler.NewUploadSuccessHandler().Handler)
	addEntry("/file/upload/duplicate", handler.NewUploadDuplicateHandler().Handler)
	addEntry("/file/scan", handler.NewGetFileMetaHandler().Handler)
	addEntry("/file/download", handler.NewDownloadHandler().Handler)
	addEntry("/file/update/name", handler.NewUpdateFileMetaHandler().Handler)
	addEntry("/file/delete", handler.NewDeleteHandler().Handler)
	addEntry("/file/404", handler.NewFileNotFound().Handler)

	addEntry("/user/signup", (&user.Signup{}).Handler)
	config()
}

// 向 routerDict 中添加新的条目
func addEntry(key string, value handlerFunc) {
	routerDict[key] = value
}

// 配置路由
func config() {
	for k, v := range routerDict {
		http.HandleFunc(k, v)
	}
}
