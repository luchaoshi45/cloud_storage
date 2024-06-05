# 微服务
## 解释
client 端处理 HTTP 请求， 转化为 RPC 请求 <br>
Consul 动态选择 server 端

## 启动
```shell
# server 端
go run service/account/main.go

# client 端 40001 端口
go run service/apigw/main.go

```

## 测试网址
http://10.181.105.230:40001/user/signup <br/>


