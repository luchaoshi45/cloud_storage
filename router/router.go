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
	// 设置静态文件路径
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	u := user.NewUser()
	f := file.NewFile()

	addEntry("/file/upload", f.Upload)
	addEntry("/file/upload/success", f.UploadSuccess)
	addEntry("/file/upload/duplicate", f.UploadDuplicate)
	addEntry("/file/scan", f.GetFileMeta)
	addEntry("/file/download", f.Download)
	addEntry("/file/update/name", f.UpdateFileMeta)
	addEntry("/file/delete", f.Delete)
	addEntry("/file/404", f.FileNotFound)

	addEntry("/user/signup", u.SignUp)
	addEntry("/user/signin", u.SignIn)
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
