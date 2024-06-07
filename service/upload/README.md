# account 微服务

## 安装插件
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

export PATH=$PATH:$HOME/go/bin
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc

go get "go-micro.dev/v4"
go get "github.com/go-micro/plugins/v4/registry/consul"
```

## 运行 protoc 命令
```shell
protoc \
--proto_path=service/upload/proto \
--go_out=service/upload/proto \
--micro_out=service/upload/proto \
service/upload/proto/upload.proto
```

