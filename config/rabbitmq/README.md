[参考](https://blog.csdn.net/big_white_py/article/details/103940139)
### 1 拉取MySQL5.7镜像到本地
```sh
docker pull rabbitmq:3.12-management

mkdir /data/rabbitmq

docker run -d --name rabbitmq \
-p 5672:5672 -p 15672:15672 -p 25672:25672 \
-v /data/rabbitmq:/var/lib/rabbitmq \
rabbitmq:3.12-management
```

http://10.181.105.230:15672

```sh
RabbitMQ默认的登录账号和密码如下：

用户名：guest
密码： guest
```