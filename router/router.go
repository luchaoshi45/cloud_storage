package router

import (
	"cloud_storage/handler"
	"net/http"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

// RouterDict 路由字典
var RouterDict = make(map[string]handlerFunc)

func Router() {
	addEntry("/file/upload", handler.NewUploadHandler().Handler)
	addEntry("/file/upload/success", handler.NewUploadSuccessHandler().Handler)
	config()
}

// 向全局字典中添加新的条目
func addEntry(key string, value handlerFunc) {
	RouterDict[key] = value
}

func config() {
	for k, v := range RouterDict {
		http.HandleFunc(k, v)
	}
}
