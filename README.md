## goctl和protoc相关安装

``` 
`goctl安装`
go install github.com/zeromicro/go-zero/tools/goctl@latest

`protoc相关安装`
https://github.com/protocolbuffers/protobuf/releases 下载对应版本
解压后，将bin目录下配置到PATH环境变量中

`protoc-gen-go安装`
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

`protoc-gen-go-grpc安装`
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
``` 

## goctl生成代码

``` 
`生成api相关文件`
goctl api new hello01

进入hello01目录，运行以下命令，下载依赖
go mod tidy
```

## 编写逻辑和修改api

在logic目录下文件写具体逻辑

根目录下有一个文件 hello01.api 这是自动生成代码的api接口描述文件，通过修改这里的api接口说明，便可以重新生成相关文件

```
`生成api相关文件`
goctl api go -api hello01.api -dir .
```

## 在部署时，配置文件并不会打包到二进制包中

我们需要指定配置文件，通过命令行参数的形式指定
```
go run hello01.go -f ./hello01-api.toml
```

## user-api下 goctl 生成 model mysql 
```
go get github.com/go-sql-driver/mysql
goctl model mysql ddl --src ./internal/model/user.sql --dir ./internal/model
```


## grpc
```
生成proto文件
goctl rpc -o greet.proto

用proto文件生成rpc相关文件
goctl rpc protoc greet.proto --go_out=./grpc-server --go-grpc_out=./grpc-server --zrpc_out=./grpc-server
```

grpcurl工具下载
```
https://github.com/fullstorydev/grpcurl/releases 

获取grpc服务列表 --plaintext选项代表适用http链接
grpcurl -plaintext localhost:8080 list

获取服务细节 或者 方法细节
grpcurl -plaintext localhost:8080 describe

调用方法
grpcurl -d {\"ping\":\"ping1\"} -plaintext localhost:8080 greet.Greet/Ping

也可以先启动命令，然后再输入输出，参数输入完后Windows下Ctrl+z(其他系统一般ctrl+d)结束输入，返回结果
grpcurl -d @ -plaintext localhost:8080 greet.Greet/Ping
{"ping":"myping"}
```