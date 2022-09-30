SET CGO_ENABLE=0
SET GOOS=linux
SET GOARCH=amd64

pathprefix=/pathto/go-spring-demo

# 生成swagger
chmod +x ${pathprefix}/gen/v1/swagger.sh
${pathprefix}/gen/v1/swagger.sh ${pathprefix}

# 生成pb
chmod +x ${pathprefix}/gen/protobuf/generator.sh
${pathprefix}/gen/protobuf/generator.sh ${pathprefix}

go build ${pathprefix}/main.go
