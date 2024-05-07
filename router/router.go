package router

import (
	"cloud_storage/handler/file"
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
	addEntry("/file/upload", (&file.Upload{}).Handler)
	addEntry("/file/upload/success", (&file.UploadSuccess{}).Handler)
	addEntry("/file/upload/duplicate", (&file.UploadDuplicate{}).Handler)
	addEntry("/file/scan", (&file.GetFileMeta{}).Handler)
	addEntry("/file/download", (&file.Download{}).Handler)
	addEntry("/file/update/name", (&file.UpdateFileMeta{}).Handler)
	addEntry("/file/delete", (&file.Delete{}).Handler)
	addEntry("/file/404", (&file.FileNotFound{}).Handler)

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
