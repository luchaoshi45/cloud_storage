[参考](https://blog.csdn.net/big_white_py/article/details/103940139)
### 1 拉取MySQL5.7镜像到本地
```
docker search redis
docker pull redis # 拉取 redis 镜像
docker images # 查看 redis 镜像

mkdir ~/tools/redis;cd ~/tools/redis # 创建文件夹
wget http://download.redis.io/redis-stable/redis.conf # 下载配置文件

# 启动容器
docker run \
--name redis \
-d -p 6379:6379 \
-v ~/tools/redis/redis.conf:/etc/redis.conf \
redis redis-server /etc/redis.conf

# redis log 报错 使用下面的命令
docker run -d --name redis \
-p 6379:6379 \
redis



docker ps 查看容器运行情况
docker exec -it redis /bin/bash # 进入redis
redis-cli # 测试连接
redis-cli -h 127.0.0.1 -p 6379 -a testupload

docker restart redis
docker stop redis
docker rm redis

```