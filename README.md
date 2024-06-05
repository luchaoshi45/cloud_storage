# Cloud Storage 基于微服务架构的分布式云存储系统

## 用户服务集群
```markdown
[README.md](service/README.md)
```

### 目录
#### [config](../config)    配置文件
#### [configurator](../configurator)    配置文件解析器
#### [db](../db)    数据库
- [ceph](../db/ceph)
- [mysql](../db/mysql)
- [oss](../db/oss)
- [redis](../db/redis)
#### [file](../file)    用户文件
#### [handler](../handler)    路由服务函数
#### [rabbitmq](../rabbitmq)    消息队列 用于上传文件的异步处理
#### [router](../router)    路由
#### [run](run)    普通方式运行
#### [service](service)    微服务方式运行
#### [service](../service)    两个 main 函数
#### [static](../static)    静态资源
#### [test](../test)    测试文件

