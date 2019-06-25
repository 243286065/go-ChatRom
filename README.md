# go-ChatRom
A web chat program writing by golang

解决依赖的问题：
```
go install github.com/gin-gonic/gin
go install github.com/micro/go-micro

git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
cd $GOPATH/src/
go install google.golang.org/grpc
```

本项目已经使用govendor，将所有的依赖库全部备份。

备份依赖库的命令是:
```
govendor add +external
```


生成go版本的proto：
```
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/protoc-gen-micro
export PATH="$PATH:$GOPATH/bin"

sudo apt  install protobuf-compiler
protoc --proto_path=service/account/proto/ --go_out=service/account/proto --micro_out=service/account/proto service/account/proto/account.proto

```

对于选用consul作为微服务注册中心，启动服务：
```
# 安装consul
sudo apt install consul

# 启动一个consul
consul agent -server -ui -bootstrap-expect=1 -data-dir=/tmp/consul -node=agent-one -advertise=127.0.0.1 -bind=0.0.0.0 -client=0.0.0.0

#-server：服务器模式
#-ui:能webui展示
#-bootstrap-expect：server为1时即选择server集群leader
#-data-dir:consul状态存储文件地址
#-node：指定结点名
#-advertise：本地ip地址
#-client:指定可访问这个服务结点的ip，0.0.0.0表示允许所有client访问

```
然后才可运行我们的微服务：
```
go run service/account/main.go --registry=consul
```
通过`http://127.0.0.1:8500`可以进入web服务查看


待开发：
1. 接入Rabbitmq，实现上下线提醒、文字聊天功能；
2. 接入ceph。实现文件上传、下载功能。