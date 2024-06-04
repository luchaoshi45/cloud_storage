### consul
https://blog.csdn.net/zhoupenghui168/article/details/134611262 <br>

#### 拉取镜像
```shell
docker pull consul:1.15.1
```

#### 创建数据卷
```shell
# 在运行之前，我们最好是先创建一个数据卷用于持久化Consul的数据
docker volume create consul-data
```

#### 单机部署
```shell
docker run -id --name=consul \
-p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8500:8500 -p 8600:8600 \
-v consul-data:/consul/data consul:1.15.1 agent -server \
-ui -node=n1 -bootstrap-expect=1 -client=0.0.0.0


可见上述除了挂载数据卷之外，我们还暴露了几个端口，它们的作用如下：
8300 TCP协议，用于Consul集群中各个节点相互连结通信的端口
8301 TCP或者UDP协议，用于Consul节点之间相互使用Gossip协议健康检查等交互
8302 TCP或者UDP协议，用于单个或多个数据中心之间的服务器节点的信息同步
8500 HTTP协议，用于API接口或者我们上述的网页管理界面访问
8600 TCP或者UDP协议，作为DNS服务器，用于通过节点名查询节点信息

所以如果是在服务器上面部署，记得配置好防火墙放行上述端口。在Spring Cloud模块集成Consul服务发现时，需要配置8500端口。
除此之外，我们来看一下命令最后的几个参数：

agent 表示启动一个Agent进程
-server 表示该节点类型为Server节点（下面会讲解集群中的节点类型）
-ui 开启网页可视化管理界面
-node 指定该节点名称，注意每个节点的名称必须唯一不能重复！上面指定了第一台服务器节点的名称为n1，那么别的节点就得用其它名称
-bootstrap-expect 最少集群的Server节点数量，少于这个值则集群失效，这个选项必须指定，由于这里是单机部署，因此设定为1即可
-advertise 这里要指定本节点外网地址，用于在集群时告诉其它节点自己的地址，如果是在自己电脑上或者是内网搭建单节点/集群则不需要带上这个参数
-client 指定可以外部连接的地址，0.0.0.0表示外网全部可以连接

除此之外，还可以加上-datacenter参数自定义一个数据中心名，同一个数据中心的节点数据中心名应当指定为一样！
```

#### 集群部署
```shell
Server节点：这是Consul集群的核心组成部分，用于维护集群的状态、处理查询请求、执行一致性协议以及提供服务发现和健康检查等功能
Client节点：用于向集群提交查询请求，并将请求转发给Server节点处理，作为服务发现和健康检查的代理，这类节点有着负载均衡、健康检查和故障转移等作用，降低Server节点的压力，搭建集群时，Client节点不是必须的
可加上 -advertise=123.20.15.63 指定外网访问地址

# 服务器1
docker run -id --name=consul1 \
-p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8500:8500 -p 8600:8600 \
-v consul-data:/consul/data consul:1.15.1 agent -server \
-ui -bootstrap-expect=3 -client=0.0.0.0 -bind=0.0.0.0

# 查看服务器1 ip -join 参数
docker inspect -f '{{.NetworkSettings.IPAddress}}' consul1

# 服务器2
docker run -id --name=consul2 \
-p 8501:8500 \
consul:1.15.1 agent -server \
-ui -bootstrap-expect=3 -client=0.0.0.0 -bind=0.0.0.0 -join 172.17.0.8

# 服务器3
docker run -id --name=consul3 \
-p 8502:8500 \
consul:1.15.1 agent -server \
-ui -bootstrap-expect=3 -client=0.0.0.0 -bind=0.0.0.0 -join 172.17.0.8

docker run --name consulClient1 -d -p 8503:8500 \
consul:1.15.1 agent -ui -bind=0.0.0.0 -client=0.0.0.0 -join 172.17.0.8

# 验证
http://10.181.105.230:8500/ui


docker exec -it consul1 sh
docker exec -it consul1 consul members

```

#### 清除数据
```shell
如果Consul运行时出现任何问题或者无法启动，可以先停止容器，然后清除所有的数据，
上述我们已经把容器中数据目录映射到数据卷了，我们删除数据卷中所有内容即可

docker stop consul1 consul2 consul3
docker rm consul1 consul2 consul3

rm -rf /var/lib/docker/volumes/consul-data/_data/*

```





