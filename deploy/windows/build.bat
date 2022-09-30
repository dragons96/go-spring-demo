set CGO_ENABLE=0
set GOOS=windows
set GOARCH=amd64

set pathprefix=D:/go_product/bigdata/go-spring-demo

:: 生成swagger
call %pathprefix%/gen/v1/swagger.bat %pathprefix%

:: 生成pb
call %pathprefix%/gen/protobuf/generator.bat %pathprefix%

:: 打包
go build %pathprefix%/main.go

