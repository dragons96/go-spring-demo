## Grpc 集成
### 1. 安装依赖
```shell
# grpc cli工具
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# grpc 依赖
go get -u github.com/go-spring/starter-grpc@v1.1.1
go get -u google.golang.org/grpc@v1.49.0
go get -u google.golang.org/protobuf@v1.28.1
```