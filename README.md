# Cloud Storage 基于微服务架构的分布式云存储系统

<img align="right" width="80px" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png">

- 基础文件上传、秒传、分块上传和断点续传
- Ceph 集群和阿里云 OSS
- Rabbitmq 异步处理
- Consul 微服务
- Docker 容器化和 K8s 集群

<img width="1200px" src=service/基于微服务的分布式云存储系统.png>

## [启动微服务](service/README.md)

## 目录
#### [config](../config)    配置文件
#### [configurator](../configurator)    配置文件解析器
#### [db](../db)    数据库
- [ceph](../db/ceph)
- [mysql](../db/mysql)
- [oss](../db/oss)
- [redis](../db/redis)
#### [file](../file)    用户文件
#### [rabbitmq](../rabbitmq)    消息队列 用于上传文件的异步处理
#### [service](service)    微服务
#### [static](../static)    静态资源
#### [test](../test)    测试文件

