package router

import (
	"cloud_storage/handler"
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
	addEntry("/file/get/meta", handler.NewGetFileMetaHandler().Handler)
	addEntry("/file/download", handler.NewDownloadHandler().Handler)
	addEntry("/file/update/meta", handler.NewUpdateFileMetaHandler().Handler)
	addEntry("/file/delete", handler.NewDeleteHandler().Handler)
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
