# 微服务
## 解释
client 端处理 HTTP 请求， 转化为 RPC 请求 <br>
Consul 动态选择 server 端

## 启动
```shell
# account rpc server
go run service/account/main.go

# account client(40001 端口)
go run service/apigw/main.go


# upload client(40002 端口)
# uploadEntry rpc server
go run service/upload/main.go

# download client(40003 端口)
# downloadEntry rpc server
go run service/download/main.go

# transfer client
go run service/transfer/main.go
```

## 测试网址
http://10.181.105.230:40001/user/signup <br/>
http://10.181.105.230:40001/user/signin <br/>


