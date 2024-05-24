[参考](https://www.cnblogs.com/niuben/p/17429673.html)
[参考](https://www.cnblogs.com/aganippe/p/16095588.html)
```shell
# 创建Ceph专用网络
sudo docker network create --driver bridge --subnet 172.20.0.0/16 ceph-network
docker network ls
# 拉取docker镜像
sudo docker pull ceph/daemon:latest-luminous
# 搭建mon节点
sudo docker run -d --name ceph-mon \
--network ceph-network --ip 172.20.0.10 \
-e CLUSTER=ceph -e WEIGHT=1.0 -e MON_IP=172.20.0.10 \
-e MON_NAME=ceph-mon -e CEPH_PUBLIC_NETWORK=172.20.0.0/16 \
-v /etc/ceph:/etc/ceph -v /var/lib/ceph/:/var/lib/ceph/ \
-v /var/log/ceph/:/var/log/ceph/ ceph/daemon:latest-luminous mon
```

```shell
# 搭建osd节点
sudo docker exec ceph-mon ceph auth get client.bootstrap-osd -o /var/lib/ceph/bootstrap-osd/ceph.keyring

# 打开配置文件 在追加下面的两行
# osd max object name len = 256
# osd max object namespace len = 64
sudo vi /etc/ceph/ceph.conf

# 分别启动三个容器来模拟集群
sudo docker run -d --privileged=true --name ceph-osd-1 --network ceph-network --ip 172.20.0.11 \
-e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=ceph-mon -e MON_IP=172.20.0.10 -e OSD_TYPE=directory \
-v /etc/ceph:/etc/ceph -v /var/lib/ceph/:/var/lib/ceph/ -v /var/lib/ceph/osd/1:/var/lib/ceph/osd \
-v /etc/localtime:/etc/localtime:ro ceph/daemon:latest-luminous osd
sudo docker run -d --privileged=true --name ceph-osd-2 --network ceph-network --ip 172.20.0.12 \
-e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=ceph-mon -e MON_IP=172.20.0.10 -e OSD_TYPE=directory \
-v /etc/ceph:/etc/ceph -v /var/lib/ceph/:/var/lib/ceph/ -v /var/lib/ceph/osd/2:/var/lib/ceph/osd \
-v /etc/localtime:/etc/localtime:ro ceph/daemon:latest-luminous osd
sudo docker run -d --privileged=true --name ceph-osd-3 --network ceph-network --ip 172.20.0.13 \
-e CLUSTER=ceph -e WEIGHT=1.0 -e MON_NAME=ceph-mon -e MON_IP=172.20.0.10 -e OSD_TYPE=directory \
-v /etc/ceph:/etc/ceph -v /var/lib/ceph/:/var/lib/ceph/ -v /var/lib/ceph/osd/3:/var/lib/ceph/osd \
-v /etc/localtime:/etc/localtime:ro ceph/daemon:latest-luminous osd
```

```shell
# 搭建mgr节点
sudo docker run -d --privileged=true --name ceph-mgr --network ceph-network --ip 172.20.0.14 \
-e CLUSTER=ceph -p 7000:7000 --pid=container:ceph-mon \
-v /etc/ceph:/etc/ceph -v /var/lib/ceph/:/var/lib/ceph/ ceph/daemon:latest-luminous mgr

# 开启管理界面，访问地址 http://127.0.0.1:7000/
sudo docker exec ceph-mgr ceph mgr module enable dashboard
```

```shell
# 搭建rgw节点
sudo docker exec ceph-mon ceph auth get client.bootstrap-rgw -o /var/lib/ceph/bootstrap-rgw/ceph.keyring
sudo docker run -d --privileged=true --name ceph-rgw --network ceph-network --ip 172.20.0.15 \
-e CLUSTER=ceph -e RGW_NAME=ceph-rgw -p 7480:7480 \
-v /var/lib/ceph/:/var/lib/ceph/ -v /etc/ceph:/etc/ceph -v /etc/localtime:/etc/localtime:ro ceph/daemon:latest-luminous rgw
```

```shell
# 检查Ceph状态
sudo docker exec ceph-mon ceph -s
# 测试添加rgw用户
sudo docker exec ceph-rgw radosgw-admin user create --uid="test" --display-name="test user"
```