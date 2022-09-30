## 创建一个go-spring-demo目录作为项目根目录, 并创建一系列子目录, 结构如下所示：
```text
go-spring-demo/
├- api/                     api swagger目录
│   └- v1/                  服务版本划分
│
├- config/                  配置文件目录
│   ├- application.yml      本地配置文件(文件后缀名可选properties, yaml, yml, toml, tml, 视自己习惯而定)
│   ├- application-dev.yml  开发环境配置(非必须, 视业务场景而定)
│   ├- application-test.yml 测试环境配置(非必须, 视业务场景而定)
│   └- application-prod.yml 生产环境配置
│
├- deploy/                  部署配置
│
├- docs/                    设计和用户文档
│   ├- en/                  英文文档(非必须) 
│   └- zh/                  中文文档(非必须)
│
├- gen/                     自动生成脚本
│   └- v1/                  版本划分
│
├- internal/                业务相关代码
│   ├- cmd/                 命令行工具
│   ├- consts/              常量
│   ├- controller/          控制器
│   ├- dao/                 dao层
│   ├- middleware/          中间件
│   ├- model/               模型
│   ├- pkg/                 其他工具
│   ├- proto/               protobuf相关文件
│   │   └- pb               protoc生成的pb.go文件
│   ├- repository/          数据源
│   └- service/             服务层
│
├- third_party/             外部辅助工具，fork的代码和其他第三方工具
│
├- main.go                  主运行程序
└- README.md                项目概要说明文档
``` 

#### [下一页：项目基础功能](base.md)