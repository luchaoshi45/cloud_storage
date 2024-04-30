### 1 拉取MySQL5.7镜像到本地
```
sudo docker pull mysql:5.7
docker ps -a
docker images

docker exec -it mysql-master mysql -uroot -p123456
docker exec -it mysql-slave mysql -uroot -p123456

docker exec -it mysql-master bash
docker exec -it mysql-slave bash
sudo docker stop mysql-master
sudo docker rm mysql-master
# 如果你只需要跑一个mysql实例，不做主从，那么执行以下命令即可
sudo docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456
mysql:5.7
```


### 2 准备MySQL配置文件
[config/mysql/master.conf](master.conf)
[config/mysql/slave.conf](slave.conf)


### 3 Docker分别运行 MySQL 主/从两个容器
#### mysql主节点
```
mkdir -p ../../database/mysql-master
# 挂载需要绝对路径
sudo docker run -d --name mysql-master -p 3306:3306 \
-v /data/xyun/code/cloud_storage/config/mysql/master.conf:/etc/mysql/mysql.conf.d/mysqld.cnf \
-v /data/xyun/database/mysql-master:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
# 查看 3306 端口
netstat -tuln | grep 3306
```
#### mysql从节点
```
mkdir -p ../../database/mysql-slave
# 挂载需要绝对路径
sudo docker run -d --name mysql-slave -p 3307:3306 \
-v /data/xyun/code/cloud_storage/config/mysql/slave.conf:/etc/mysql/mysql.conf.d/mysqld.cnf \
-v /data/xyun/database/mysql-slave:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
# 查看 3307 端口
netstat -tuln | grep 3307
```


### 4 MySQL 主从节点配置同步信息
####  MySQL 主节配置
```
# 登录主要节点 xx.xx.xx.xx 是你本机的内网ip,可以通过 ifconfig 查看
mysql -u root -h 10.181.105.230 -P3306 -p123456
# 在 mysql client中执行
create user slave identified by 'slave';
GRANT REPLICATION SLAVE ON *.* TO 'slave'@'%' IDENTIFIED BY 'slave';
flush privileges;
create database fileserver default character set utf8mb4;

```
#### MySQL 从节配置
```
# 主节对话框 查看状态
SHOW MASTER STATUS;
# 查看mysql master的容器独立ip地址
docker inspect --format='{{.NetworkSettings.IPAddress}}' mysql-master
# 登录从节点
mysql -u root -h 10.181.105.230 -P3307 -p123456
# 在 mysql client 中执行 注意其中的日志文件和数值要和上面show master status的值对应
stop slave;
create database fileserver default character set utf8mb4;
CHANGE MASTER TO MASTER_HOST='172.17.0.2', \
MASTER_PORT=3306, \
MASTER_USER='slave', \
MASTER_PASSWORD='slave', \
MASTER_LOG_FILE='log.000004', \
MASTER_LOG_POS=1035;
start slave;
# 检查
show slave status \G;
/*
...
Slave_IO_Running: Yes 
Slave_SQL_Running: Yes 
...
*/

### 测试
docker exec -it mysql-master mysql -uroot -p123456

use fileserver;
CREATE TABLE example_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

docker exec -it mysql-slave mysql -uroot -p123456
use fileserver;
SHOW TABLES;

```