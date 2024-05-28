### 命令
```shell
go run service/main.go
go run service/transfer.go
```



### 测试网址
http://10.181.105.230:42200/file/upload
http://10.181.105.230:42200/user/signup
http://10.181.105.230:42200/user/signin

http://10.181.105.230:42200/file/downloadurl?sha1=167cb9619a97d84458f789a2209887e6fd518f9e
http://10.181.105.230:42200/file/get/meta?sha1=22d9ebe1a35871d068c5b83df46f96174d3d86e9
http://10.181.105.230:42200/file/get/meta?sha1=22d9ebe1a35871d068c5b83df46f96174d3d86e9
http://10.181.105.230:42200/file/download?sha1=d267485cad77dbb6d4a434bc682afd5c9acaa876
http://10.181.105.230:42200/file/update/meta?sha1=22d9ebe1a35871d068c5b83df46f96174d3d86e9&name=hh.png
http://10.181.105.230:42200/file/delete?sha1=9943ff60b76eae053e317da40dee9e3105210034
http://10.181.105.230:42200/file/delete?sha1=b090ed884b07d2d98747141aefd25590b8b254f9

### 测试
```
go run test/test.redis.go
go run test/test_mpupload.go
```


### 目录
#### [config](config)    配置文件
#### [configurator](configurator)    配置文件解析器
#### [db](db)    数据库
#### [file](file)    用户文件
#### [handler](handler)    路由服务函数
#### [rabbitmq](rabbitmq)    rabbitmq 消息队列 用于上传文件的异步处理
#### [router](router)    路由
#### [service](service)    两个 main 函数
#### [static](static)    静态资源
#### [test](test)    测试文件
#### [tmp](tmp)    文件临时存储目录
#### [util](file/util)    算法工具