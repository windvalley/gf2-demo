# gf2-demo

[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://go.dev)
[![Version](https://img.shields.io/github/v/release/windvalley/gf2-demo?include_prereleases)](https://github.com/windvalley/gf2-demo/releases)
![Supports](https://img.shields.io/badge/Supports-macOS,%20Linux-orange)
[![LICENSE](https://img.shields.io/github/license/windvalley/gf2-demo)](LICENSE)

`gf2-demo` 是一个基于 [GoFrameV2](https://github.com/gogf/gf) 用来快速开发后端服务的脚手架, 目标使开发者只需关注业务逻辑的编写, 快速且规范地交付项目.

## 💌 Features

- 优化工程目录结构, 使支持多个可执行命令
- 多环境管理: 开发环境、测试环境、生产环境
- 编译的二进制文件可打印当前应用的版本信息
- 中间件统一拦截响应, 规范响应格式, 规范业务错误码
- 完善 HTTP 服务访问日志、HTTP 服务错误日志、SQL 日志、开发者打印的日志、其他可执行命令的日志配置
- 封装 `Redis` 常用工具库: `rediscache`, `redislock`, `redismq`, `redisdelaymq`, `redispubsub`
- 通过工具自动生成数据库层、服务接口层、控制器层代码
- 完整的增删改查接口示例和完善的开发流程文档, 帮助开发者快速上手
- 项目部署遵循不可变基础设施原则, 不论是传统单体部署还是容器云部署方式
- 通过 `Makefile` 管理项目: `make run`, `make build`, `make dao`, `make service` 等
- 增加 `golangci-lint` 配置文件 `.golangci.yml`, 统一团队代码风格, 保障团队代码质量
- 适合个人开发者高质量完成项目, 也适合团队统一后端技术框架, 规范高效管理

## 🚀 Quick Start

### 安装

```sh
git clone --depth 1 git@github.com:windvalley/gf2-demo.git

cd gf2-demo

# 安装gf
make cli

# 导入测试用mysql库表
mysql -uroot -p'123456' < manifest/sql/gf2_demo.sql
```

> 请提前安装 Go 环境, 要求 Go 版本: `1.15+`

### 热更新(Live reload)

开发环境下使用.

```sh
cd gf2-demo

# 运行 gf2-demo-api
make run

# 运行 gf2-demo-cli
make run.cli
```

> 默认加载配置文件: `manifest/config/config.yaml`

### 访问测试

```sh
$ curl -X GET -i 'localhost:9000/v1/demo/windvalley'

HTTP/1.1 200 OK
Content-Type: application/json
Server: GoFrame HTTP Server
Trace-Id: 506dccff4a08431731f5d0259180c3b8
Date: Sun, 12 Feb 2023 09:03:24 GMT
Content-Length: 130

{"code":"OK","message":"","traceid": "506dccff4a08431731f5d0259180c3b8","data":{"id":1,"fielda":"windvalley","created_at":"2008-08-08 08:08:08","updated_at":"2008-08-08 08:08:08"}}
```

### 编译二进制文件

```sh
cd gf2-demo

# 编译 gf2-demo-api
make build

# 编译 gf2-demo-cli
make build.cli
```

会生成如下二进制文件:

```text
bin
├── darwin_amd64
│   └── gf2-demo-api
│   └── gf2-demo-cli
└── linux_amd64
    └── gf2-demo-api
    └── gf2-demo-cli
```

### 打印帮助信息

```sh
$ bin/darwin_amd64/gf2-demo-api -h

USAGE
    gf2-demo-api [OPTION]

OPTION
    -v, --version   print version info
    -c, --config    config file (default config.yaml)
    -h, --help      more information about this command

EXAMPLE
    Dev:
    ./gf2-demo-api
    Test:
    ./gf2-demo-api -c config.test.yaml
    or
    GF_GCFG_FILE=config.test.yaml GF_GERROR_BRIEF=true ./gf2-demo-api
    Prod:
    ./gf2-demo-api -c config.prod.yaml
    or
    GF_GCFG_FILE=config.prod.yaml GF_GERROR_BRIEF=true ./gf2-demo-api

DESCRIPTION
    An API server demo using GoFrame V2

Find more information at: https://github.com/windvalley/gf2-demo
```

## 📄 Documentation

- [工程目录](#工程目录-)
- [环境管理](#环境管理-)
  - [开发环境](#开发环境)
  - [测试环境](#测试环境)
  - [生产环境](#生产环境)
- [多命令管理](#多命令管理-)
  - [目录设计](#目录设计)
  - [配置文件](#配置文件)
- [错误码管理](#错误码管理-)
  - [规范制定](#规范制定)
  - [业务错误码](#业务错误码)
  - [响应示例](#响应示例)
- [日志管理](#日志管理-)
  - [HTTP 服务日志](#http-服务日志)
    - [1. HTTP 服务日志配置](#1-http-服务日志配置)
    - [2. 生成的日志示例](#2-生成的日志示例)
  - [SQL 日志](#sql-日志)
    - [1. SQL 日志配置](#1-sql-日志配置)
    - [2. 生成的日志示例](#2-生成的日志示例)
  - [开发者打印的通用日志](#开发者打印的通用日志)
    - [1. 通用日志配置](#1-通用日志配置)
    - [2. 如何打日志](#2-如何打日志)
    - [3. 生成的日志示例](#3-生成的日志示例)
- [链路跟踪](#链路跟踪-)
- [版本管理](#版本管理-)
  - [1. 写版本变更文档](#1-写版本变更文档)
  - [2. 给项目仓库打 tag](#2-给项目仓库打-tag)
  - [3. 使用 Makefile 编译](#3-使用-makefile-编译)
  - [4. 查看二进制文件版本信息](#4-查看二进制文件版本信息)
- [Redis](#redis-)
  - [Redis 配置](#redis-配置)
  - [Redis 工具库](#redis-工具库)
- [开发流程](#开发流程-)
  - [1. 设计表结构, 创建物理表](#1-设计表结构-创建物理表)
  - [2. 自动生成数据层相关代码](#2-自动生成数据层相关代码)
  - [3. 编写 api 层代码](#3-编写-api-层代码)
  - [4. 自动生成 controller 层框架代码](#4-自动生成-controller-层框架代码)
  - [5. 编写 model 层代码](#5-编写-model-层代码)
  - [6. 编写 service 层代码](#6-编写-service-层代码)
  - [7. 编写 controller 层代码](#7-编写-controller-层代码)
  - [8. 路由注册](#8-路由注册)
  - [9. 接口访问测试](#9-接口访问测试)
- [代码质量](#代码质量-)
- [项目部署](#项目部署-)
  - [Systemctl](#systemctl)
  - [Supervisor](#supervisor)
  - [Docker](#docker)
  - [优雅关闭测试](#优雅关闭测试)
- [使用 Makefile 管理项目](#使用-makefile-管理项目-)
- [变更项目名称](#变更项目名称-)

### 工程目录 [⌅](#-documentation)

```sh
├── CHANGELOG.md  # 版本变更管理
├── Dockerfile  # 用于制作容器镜像
├── Makefile  # 用于项目管理
├── README.md  # 项目文档
├── api  # 对外接口定义: 对外提供服务的输入/输出数据结构定义, 路由path定义, 数据校验等
│   ├── api.go  # 接口模块通用结构
│   └── demo  # demo模块
│       ├── demo.go  # demo模块的api interface, 由make ctrl自动生成
│       └── v1  # 版本控制
│           └── demo.go  # 开发者按照规范编写的接口文件, make ctrl会根据本文件自动生成controller代码
├── bin  # make build 和 make build.cli 生成的二进制可执行文件所在目录, 不要提交到仓库
│   ├── darwin_amd64
│   │   ├── gf2-demo-api
│   │   └── gf2-demo-cli
│   └── linux_amd64
│       ├── gf2-demo-api
│       └── gf2-demo-cli
├── cmd  # 项目的可执行文件入口
│   ├── gf2-demo-api  # API服务
│   │   └── gf2-demo-api.go  # 注意: 编译时会使用入口文件的名字作为二进制文件名称
│   └── gf2-demo-cli  # 项目的其他可执行服务, 比如可以是: 命令行工具或Daemon后台程序等和项目相关的辅助应用
│       └── gf2-demo-cli.go
├── hack  # 存放项目开发工具、脚本等内容. 例如: gf工具的配置, 各种shell/bat脚本等文件
│   └── config.yaml  # gf 工具的配置文件, 比如 gf gen/gf build 等会使用这里的配置内容
│   └── change_project_name.sh  # 将示例项目名称改成你自己的项目名称
├── internal
│   ├── cmd  # 对应外层 cmd 目录
│   │   ├── apiserver  # 对应 gf2-demo-api, 命令配置, 路由注册等
│   │   │   └── apiserver.go
│   │   └── cli  # 对应 gf2-demo-cli, 命令配置等
│   │       └── cli.go
│   ├── codes  # 业务错误码定义维护
│   │   ├── biz_codes.go
│   │   └── codes.go
│   ├── consts  # 项目所有通用常量定义
│   │   └── consts.go
│   ├── controller  # 控制器层: 接收/解析用户输入参数的入口
│   │   └── demo  # 本目录和目录下文件由make ctrl自动生成, 除了demo_new.go不能修改, 其他文件均需要添加具体的控制器实现(比如和service联动)
│   │       ├── demo.go
│   │       ├── demo_new.go
│   │       ├── demo_v1_create.go
│   │       ├── demo_v1_delete.go
│   │       ├── demo_v1_get_list.go
│   │       ├── demo_v1_get_one.go
│   │       └── demo_v1_update.go
│   ├── dao  # 数据访问对象, 由make dao自动生成. 这是一层抽象对象, 用于和底层数据库交互, 仅包含最基础的 CURD 方法. dao层通过框架的ORM抽象层组件与底层真实的数据库交互
│   │   ├── demo.go
│   │   └── internal
│   │       └── demo.go
│   ├── logic  # 业务封装: 业务逻辑封装管理, 特定的业务逻辑实现和封装. 往往是项目中最复杂的部分. logic层的业务逻辑需要通过调用dao来实现数据的操作, 调用dao时需要传递do数据结构对象, 用于传递查询条件、输入数据. dao执行完毕后通过Entity数据模型将数据结果返回给service(logic)层
│   │   ├── logic.go
│   │   ├── demo  # demo服务的具体实现
│   │   │   └── demo.go
│   │   └── middleware  # 中间件
│   │       ├── middleware.go
│   │       ├── accessuser.go
│   │       ├── response.go  # 统一拦截规范响应
│   │       └── traceid.go
│   ├── model  # 数据结构管理模块, 管理数据实体对象, 以及输入与输出数据结构定义. 这里的model不仅负责维护数据实体对象(entity)结构定义, 也包括所有的输入/输出数据结构定义, 被api/dao/service共同引用
│   │   ├── demo.go  # 输入/输出数据结构定义
│   │   ├── do  # 领域对象: 用于dao数据操作中业务模型与实例模型转换. NOTE: 由工具维护(make dao), 不要手动修改
│   │   │   └── demo.go
│   │   └── entity  # 数据模型: 是模型与数据集合的一对一关系, 通常和数据表一一对应. NOTE: 由工具维护(make dao), 不要手动修改
│   │   │   └── demo.go
│   ├── packed
│   │   └── packed.go
│   └── service  # 业务接口层: 用于业务模块解耦的接口定义层. 具体的接口实现在logic中进行注入. NOTE: 由工具维护(make service), 不要手动修改
│       ├── demo.go
│       └── middleware.go
├── manifest  # 交付清单: 包含应用配置文件, 部署文件等
│   ├── config  # 配置文件存放目录, 可通过gf build/make build打包到二进制文件中
│   │   ├── config.prod.yaml  # 生产环境
│   │   ├── config.test.yaml  # 测试环境
│   │   └── config.yaml  # 开发环境
│   ├── deploy  # 和部署相关的文件
│   │   ├── kustomize  # Kubernetes集群化部署的Yaml模板, 通过kustomize管理
│   │   │   ├── base
│   │   │   │   ├── deployment.yaml
│   │   │   │   ├── kustomization.yaml
│   │   │   │   └── service.yaml
│   │   │   └── overlays
│   │   │       └── develop
│   │   │           ├── configmap.yaml
│   │   │           ├── deployment.yaml
│   │   │           └── kustomization.yaml
│   │   ├── supervisor  # 通过 supervisor 管理服务
│   │   │   ├── deploy.sh  # 一键部署脚本
│   │   │   ├── gf2-demo-api.ini  # 本项目生产环境supervisor配置文件
│   │   │   └── gf2-demo-api_test.ini  # 本项目测试环境supervisor配置文件
│   │   └── systemctl  # 通过systemctl管理服务
│   │       ├── deploy.sh  # 一键部署脚本
│   │       ├── gf2-demo-api.service  # 生产环境服务文件
│   │       └── gf2-demo-api_test.service  # 测试环境服务文件
│   └── sql
│       └── gf2_demo.sql  # 用于创建示例表的sql文件
├── resource  # 静态资源文件: 这些文件往往可以通过资源打包/镜像编译的形式注入到发布文件中, 纯后端api服务一般用不到此目录
│   ├── i18n
│   ├── public
│   │   ├── html
│   │   ├── plugin
│   │   └── resource
│   │       ├── css
│   │       ├── image
│   │       └── js
│   └── template
└── utility  # 通用工具类
    ├── accessuser.go
    └── version.go
```

### 环境管理 [⌅](#-documentation)

#### 开发环境

配置文件: `manifest/config/config.yaml`

运行:

`make run` 或 `./gf2-demo-api`

> 会默认加载配置文件 config.yaml

#### 测试环境

配置文件: `manifest/config/config.test.yaml`

运行:

- 通过环境变量指定配置文件: `GF_GCFG_FILE=config.test.yaml GF_GERROR_BRIEF=true ./gf2-demo-api`
- 通过命令行参数指定配置文件: `./gf2-demo-api -c config.test.yaml`

> NOTE:
>
> - 通过命令行参数指定配置文件优先于环境变量.
> - -c 参数指定的配置文件可以使用绝对路径, 如果不包含路径, 默认依次在如下路径搜索配置文件: `./`(二进制所在的当前目录) > `./config/` > `./manifest/config/`.
> - `GF_GERROR_BRIEF=true` 表示 HTTP 服务日志错误堆栈中不包含 gf 框架堆栈.
> - 配置文件在通过 `make build` 或 `make build.cli` 编译时已经打包到二进制文件中, 所以在部署时只需部署二进制文件即可.

#### 生产环境

配置文件: `manifest/config/config.prod.yaml`

运行:

同测试环境, 只不过指定的配置文件不同, 略.

### 多命令管理 [⌅](#-documentation)

#### 目录设计

举例:

- 命令 1: `cmd/gf2-demo-api/gf2-demo-api.go` -> `internal/cmd/apiserver/apiserver.go`
- 命令 2: `cmd/gf2-demo-cli/gf2-demo-cli.go` -> `internal/cmd/cli/cli.go`

#### 配置文件

默认不同命令在相同环境下使用同一个配置文件, 比如 `gf2-demo-api` 和 `gf2-demo-cli` 在开发环境下都使用 `manifest/config/config.yaml` 作为配置文件.

不过也可以使用各自独立的配置文件, 只需要在运行时通过环境变量或命令行参数指定需要使用的配置文件即可, 比如:

`./gf2-demo-cli -c cli_config.yaml` 或
`GF_GCFG_FILE=cli_config.yaml ./gf2-demo-cli`

### 错误码管理 [⌅](#-documentation)

#### 规范制定

- 统一响应格式

  不论是正确还是错误响应, 响应体都统一使用如下格式:

  ```json
  {
    "code": "string",
    "message": "string",
    "traceid": "string",
    "data": null
  }
  ```

  > 💡 **响应 header 中已经有了`Trace-Id`了, 为什么响应 json 中还要加一个`traceid` ?**
  >
  > 目的是在遇到错误问题进行排查时, 减少不必要的沟通成本, 毕竟很多用户容易忽略响应 header,
  > 在响应体中直接体现 `traceid` 更直接. 这样在快速拿到用户反馈的 `traceid` 后,
  > 我们就可以很快找到对应的日志从而高效解决问题了.

- 业务码  
  统一使用字符串表示, 如: `"code": "ValidationFailed"`

- HTTP 状态码
  - 正确响应
    - `200`: 成功的响应
    - `202`: 部分成功的响应
  - 客户端错误
    - `401`: 未通过访问认证
    - `403`: 请求的资源未获得授权
    - `404`: 请求的资源不存在
    - `400`: 其他所有客户端错误, 比如请求参数验证失败等
  - 服务端错误
    - `500`: 所有服务端错误

#### 业务错误码

请在 `internal/codes/biz_codes.go` 文件中维护业务错误码.

```go
package codes

//  http status, bisiness code, message
var (
	CodeOK          = New(200, "OK", "")
	CodePartSuccess = New(202, "PartSuccess", "part success")

	CodePermissionDenied = New(401, "AuthFailed", "authentication failed")
	CodeNotAuthorized    = New(403, "NotAuthorized", "resource is not authorized")
	CodeNotFound         = New(404, "NotFound", "resource does not exist")
	CodeValidationFailed = New(400, "ValidationFailed", "validation failed")
	CodeNotAvailable     = New(400, "NotAvailable", "not available")

	CodeInternal = New(500, "InternalError", "an error occurred internally")
)
```

#### 响应示例

- 正确响应

  ```text
  HTTP/1.1 200 OK
  Content-Type: application/json
  Server: GoFrame HTTP Server
  Trace-Id: 10c9769ce5cf4117c19a595c2d781e94
  Date: Wed, 08 Feb 2023 09:38:41 GMT
  Content-Length: 34

  {
      "code": "OK",
      "message": "",
      "traceid": "10c9769ce5cf4117c19a595c2d781e94",
      "data": null
  }
  ```

- 401 错误

  ```text
  HTTP/1.1 401 Unauthorized
  Content-Type: application/json
  Server: GoFrame HTTP Server
  Trace-Id: a89b7652b1cf41170d6e5233fbb76a21
  Date: Wed, 08 Feb 2023 09:34:56 GMT
  Content-Length: 83

  {
      "code": "AuthFailed",
      "message": "authentication failed",
      "traceid": "a89b7652b1cf41170d6e5233fbb76a21",
      "data": null
  }
  ```

- 500 错误

  ```text
  HTTP/1.1 500 Internal Server Error
  Content-Type: application/json
  Server: GoFrame HTTP Server
  Trace-Id: 70cd58a9d8cf4117376a265eb84137e5
  Date: Wed, 08 Feb 2023 09:37:45 GMT
  Content-Length: 73

  {
      "code": "InternalError",
      "message": "an error occurred internally",
      "traceid": "70cd58a9d8cf4117376a265eb84137e5",
      "data": null
  }
  ```

### 日志管理 [⌅](#-documentation)

#### HTTP 服务日志

##### 1. HTTP 服务日志配置

```yaml
# manifest/config/config.yaml
server:
  # 服务日志(包括访问日志和server错误日志)
  logPath: "logs/" # 日志文件存储目录路径, 建议使用绝对路径. 默认为空, 表示关闭
  logStdout: true # 日志是否输出到终端. 默认为true
  errorStack: true # 当Server捕获到异常时是否记录堆栈信息到日志中. 默认为true
  errorLogEnabled: true # 是否记录异常日志信息到日志中. 默认为true
  errorLogPattern: "error-{Ymd}.log" # 异常错误日志文件格式. 默认为"error-{Ymd}.log"
  accessLogEnabled: true # 是否记录访问日志(包含异常的访问日志). 默认为false
  accessLogPattern: "access-{Ymd}.log" # 访问日志文件格式. 默认为"access-{Ymd}.log"

  # 针对服务日志的扩展配置
  logger:
    file: "{Ymd}.log" # 这里只记录server启动过程中gf默认打印的日志, 默认 {Y-m-d}.log; 日志所在路径为server.logPath指定的目录
    ctxKeys: ["user", "mail"] # 自动打印Context的指定变量到日志中. 默认为空
    rotateExpire: "1d"
    rotateBackupExpire: "30d"
    rotateBackupLimit: 30
    rotateCheckInterval: "1h"
```

##### 2. 生成的日志示例

```sh
$ curl -X GET 'localhost:9000/v1/demo' \
    -H 'X-Consumer-Custom-ID: windvalley' \
    -H 'X-Consumer-Username: windvalley@sre.im'
```

- 服务访问日志示例

  ```sh
  # 普通格式
  2023-02-08 16:50:51.992 {10fde08349cd4117115968787a401378} {windvalley, windvalley@sre.im} 401 "GET http localhost:9000 /v1/hello HTTP/1.1" 0.004, ::1, "", "PostmanRuntime/7.28.0"

  # json格式
  {"Time":"2023-02-08 16:53:13.118","TraceId":"a8b1bf5f6acd41177931ba72f7411788","CtxStr":"windvalley, windvalley@sre.im","Level":"","Content":"401 \"GET http localhost:9000 /v1/hello HTTP/1.1\" 0.002, ::1, \"\", \"PostmanRuntime/7.28.0\""}
  ```

- 服务错误日志示例

  ```sh
  # 普通格式
  2023-02-08 16:55:25.984 {2068374f89cd41170d329c50fe5a5fc8} {windvalley, windvalley@sre.im} 401 "GET http localhost:9000 /v1/hello HTTP/1.1" 0.003, ::1, "", "PostmanRuntime/7.28.0", 0, "resource is not authorized", "{Code:NotAuthorized HttpCode:401}"
  Stack:
  1. resource is not authorized: some error
  1).  gf2-demo/internal/controller.(*cHello).Hello
          /Users/xg/github/gf2-demo/internal/controller/hello.go:25
  2).  gf2-demo/internal/logic/middleware.(*sMiddleware).ResponseHandler
          /Users/xg/github/gf2-demo/internal/logic/middleware/response.go:16
  3).  gf2-demo/internal/logic/middleware.(*sMiddleware).AccessUser
          /Users/xg/github/gf2-demo/internal/logic/middleware/accessuser.go:25
  4).  gf2-demo/internal/logic/middleware.(*sMiddleware).TraceID
          /Users/xg/github/gf2-demo/internal/logic/middleware/traceid.go:27
  2. some error

  # json格式
  {"Time":"2023-02-08 16:54:28.757","TraceId":"18323afc7bcd411710d9f134cc2ec9d5","CtxStr":"windvalley, windvalley@sre.im","Level":"ERRO","Content":"401 \"GET http localhost:9000 /v1/hello HTTP/1.1\" 0.003, ::1, \"\", \"PostmanRuntime/7.28.0\", 0, \"resource is not authorized\", \"{Code:NotAuthorized HttpCode:401}\"\nStack:\n1. resource is not authorized: some error\n   1).  gf2-demo/internal/controller.(*cHello).Hello\n        /Users/xg/github/gf2-demo/internal/controller/hello.go:25\n   2).  gf2-demo/internal/logic/middleware.(*sMiddleware).ResponseHandler\n        /Users/xg/github/gf2-demo/internal/logic/middleware/response.go:16\n   3).  gf2-demo/internal/logic/middleware.(*sMiddleware).AccessUser\n        /Users/xg/github/gf2-demo/internal/logic/middleware/accessuser.go:25\n   4).  gf2-demo/internal/logic/middleware.(*sMiddleware).TraceID\n        /Users/xg/github/gf2-demo/internal/logic/middleware/traceid.go:27\n2. some error\n"}
  ```

#### SQL 日志

##### 1. SQL 日志配置

```yaml
# manifest/config/config.yaml

# doc: https://goframe.org/pages/viewpage.action?pageId=1114245
database:
  # sql日志
  logger:
    path: "logs/"
    file: "sql-{Ymd}.log"
    level: "all"
    stdout: true
    ctxKeys: ["user", "mail"]
    rotateExpire: "1d"
    rotateBackupExpire: "30d"
    rotateBackupLimit: 30
    rotateCheckInterval: "1h"
```

##### 2. 生成的日志示例

```sh
# 普通格式
2023-02-12 16:52:32.330 [DEBU] {508ad625b3074317ec81cd791f1a5993} {windvalley, windvalley@sre.im} [  2 ms] [default] [gf2_demo] [rows:5  ] SHOW FULL COLUMNS FROM `demo`
2023-02-12 16:52:32.331 [DEBU] {508ad625b3074317ec81cd791f1a5993} {windvalley, windvalley@sre.im} [  1 ms] [default] [gf2_demo] [rows:1  ] SELECT * FROM `demo` WHERE `fielda`='windvalley' LIMIT 1

# json格式
{"Time":"2023-02-12 16:55:02.420","TraceId":"28a0d817d6074317a9e647156d712d81","CtxStr":"windvalley, windvalley@sre.im","Level":"DEBU","Content":"[  3 ms] [default] [gf2_demo] [rows:5  ] SHOW FULL COLUMNS FROM `demo`"}
{"Time":"2023-02-12 16:55:02.421","TraceId":"28a0d817d6074317a9e647156d712d81","CtxStr":"windvalley, windvalley@sre.im","Level":"DEBU","Content":"[  1 ms] [default] [gf2_demo] [rows:1  ] SELECT * FROM `demo` WHERE `fielda`='windvalley' LIMIT 1"}
```

#### 开发者打印的通用日志

##### 1. 通用日志配置

```yaml
# manifest/config/config.yaml
logger:
  path: "logs/" # 日志文件目录, 如果为空, 表示不记录到文件; 建议目录和server.logPath保持一致
  file: "{Ymd}.log" # 日志文件格式. 默认为 {Y-m-d}.log; 建议和server.logger.file保持一致
  level: "all" # DEBU < INFO < NOTI < WARN < ERRO < CRIT, 也支持ALL, DEV, PROD常见部署模式配置名称. level配置项字符串不区分大小写
  stStatus: 0 # 是否打印错误堆栈(1: enabled - default; 0: disabled). 如果开启, 使用g.Log().Error 将会打印错误堆栈
  ctxKeys: ["user", "mail"] # 自动打印Context的变量到日志中. 默认为空
  stdout: true # 日志是否同时输出到终端. 默认true
  stdoutColorDisabled: false # 关闭终端的颜色打印. 默认false
  writerColorEnable: false # 日志文件是否带上颜色. 默认false, 表示不带颜色
  rotateExpire: "1d" # 多长时间切分一次日志
  rotateBackupExpire: "30d" # 删除超过多长时间的切分文件, 默认为0, 表示不备份, 切分则删除. 如果启用按时间备份, rotateBackupLimit 必须设置为一个相对较大的数
  rotateBackupLimit: 30 # 最多保留多少个切分文件, 但rotateBackupExpire的配置优先. 默认为0, 表示不备份, 切分则删除. 可以不设置rotateBackupExpire
  rotateCheckInterval: "1h" # 滚动切分的时间检测间隔, 一般不需要设置. 默认为1小时
  format: "" # "json" or other, 也对server服务日志生效

  # 为子项目gf2-demo-cli配置独立的logger
  cli:
    path: "logs/"
    file: "cli-{Ymd}.log"
    level: "all"
    stStatus: 1
    stdout: true
    stdoutColorDisabled: false
    writerColorEnable: false
    rotateExpire: "1d"
    rotateBackupExpire: "30d"
    rotateBackupLimit: 30
    rotateCheckInterval: "1h"
    format: ""
```

##### 2. 如何打日志

```go
// gf2-demo-api的日志
g.Log().Info(ctx, "hello world")
g.Log().Errorf(ctx, "hello %s", "world")

// gf2-demo-cli的日志
g.Log("cli").Debug(ctx, "hello world")
g.Log("cli").Warningf(ctx, "hello %s", "world")
```

##### 3. 生成的日志示例

```sh
# 普通格式
2023-02-08 17:02:31.906 [INFO] {389b4e7aeccd41175dd0bc18211c2519} {windvalley, windvalley@sre.im} /Users/xg/github/gf2-demo/internal/controller/hello.go:33: hello world
2023-02-08 17:02:31.906 [ERRO] {389b4e7aeccd41175dd0bc18211c2519} {windvalley, windvalley@sre.im} /Users/xg/github/gf2-demo/internal/controller/hello.go:34: hello world

# json格式
{"Time":"2023-02-08 17:04:08.957","TraceId":"d0e7f61203ce41171374033689322f91","CtxStr":"windvalley, windvalley@sre.im","Level":"INFO","CallerPath":"/Users/xg/github/gf2-demo/internal/controller/hello.go:33:","Content":"hello world"}
{"Time":"2023-02-08 17:04:08.957","TraceId":"d0e7f61203ce41171374033689322f91","CtxStr":"windvalley, windvalley@sre.im","Level":"ERRO","CallerPath":"/Users/xg/github/gf2-demo/internal/controller/hello.go:34:","Content":"hello world"}
```

### 链路跟踪 [⌅](#-documentation)

- 用于链路跟踪的响应 Header 为: `Trace-Id`, 会优先使用客户端传递的请求 Header `Trace-Id` 的值, 如果不存在会自动生成. 为了便于用户查看`Trace-Id`, 也在响应 json 中加入了 `traceid` 字段.

- 服务内部如果需要调用其他服务的接口, 请使用 `g.Client()`, 因为他会给请求自动注入`Trace-Id`, 这样不同 API 服务之间的日志就可以通过 `Trace-Id` 串起来了.

> 参考: https://goframe.org/pages/viewpage.action?pageId=49745257

### 版本管理 [⌅](#-documentation)

#### 1. 写版本变更文档

`vi CHANGELOG.md`

```text
## v0.3.0

### Added

- xxx
- xxx

### Changed

- xxx
- xxx

### Fixed

- xxx
- xxx
```

#### 2. 给项目仓库打 tag

```sh
git tag v0.3.0
git push --tags
```

#### 3. 使用 Makefile 编译

- gf 工具配置(`hack/config.yaml`)

  ```yaml
  # doc: https://goframe.org/pages/viewpage.action?pageId=3673173
  gfcli:
    # doc: https://goframe.org/pages/viewpage.action?pageId=1115788
    build:
      path: "./bin" # 编译生成的二进制文件的存放目录. 生成的二进制名称默认与程序入口go文件同名
      arch: "amd64"
      system: "linux,darwin"
      # 将项目需要的配置文件打包进二进制, 这样项目部署的时候就可以不用拷贝配置文件了.
      # NOTE: 1) 如果开启了打包功能, 但还是想再使用外部配置文件,
      # 需要通过环境变量GF_GCFG_FILE或-c参数指定配置文件路径.
      # 例如：
      # GF_GCFG_FILE=./config.prod.yaml ./gf2-demo-api
      # 或: ./gf2-demo-api -c ./config.prod.yaml
      # 2) 使用外部配置文件的好处: 可以不用重新编译变更配置文件内容, 支持配置文件热更新.
      # 3) 使用外部配置文件的缺点: 配置文件的变化没有经过版本管理, 出现问题不方便回溯历史.
      packSrc: "manifest/config"
      extra: ""
      # 编译时的内置变量可以在运行时通过gbuild包获取, 比如: utility/version.go
      varMap:
        # NOTE: 1) `version` is used by `make build` to generate binary version. Do Not Edit.
        # 2) You can manage the project versions by command `git tag vX.X.X`
        version:
  ```

- 编译

  ```sh
  # For gf2-demo-api
  make build

  # For gf2-demo-cli
  make build.cli
  ```

#### 4. 查看二进制文件版本信息

```sh
$ bin/darwin_amd64/gf2-demo-api -v

输出:

App Version: v0.3.0-7-g1898e82
Git Commit:  2023-02-08 14:55:39 1898e82dbcb4c2e8a091eb12fc96ead2f04f5993
Build Time:  2023-02-08 15:31:20
Go Version:  go1.17.6
GF Version:  v2.3.1
```

### Redis [⌅](#-documentation)

#### Redis 配置

```yaml
# manifest/config/config.yaml

# doc: https://goframe.org/pages/viewpage.action?pageId=1114217
redis:
  # 默认分组, 调用方式: g.Redis()
  default:
    address: 127.0.0.1:6379
    # 数据库索引, 0-15
    db: 0
    # 访问授权密码
    pass:
    # 连接最长存活时间, 默认值: 30s, 建议设置的长一些.
    # 建立连接后, 可用此连接进行多次reids请求,
    # 从建立连接开始计算, 只要超过这个时间就会自动断开连接,
    # 就算此时有redis请求也会先关闭当前连接, 然后重新建一个新的连接,
    # 除非请求是一个阻塞式请求, 比如: BLPop等, 此时连接将始终保持,
    # 直到从队列读取到数据才会断开连接.
    maxConnLifetime: 30m
    # 连接最大空闲时间, 默认值: 10s;
    # 只要连接空闲(没有新请求)超过idleTimeout后, 就会断开连接.
    # 此值应该小于maxConnLifetime, 才有实际意义.
    idleTimeout: 10m
  # 延迟队列分组, 调用方式: g.Redis("delayqueue")
  delayqueue:
    address: 127.0.0.1:6379
    db: 1
    pass:
    maxConnLifetime: 30m
    idleTimeout: 10m
```

#### Redis 工具库

- Redis 缓存: `internal/pkg/rediscache`
- Redis 分布式锁: `internal/pkg/redislock`
- Redis 消息队列: `internal/pkg/redismq`
- Redis 延迟队列: `internal/pkg/redisdelaymq`
- Redis 发布订阅: `internal/pkg/redispubsub`

使用方法可参考代码或每个包下面的 test 文件.

### 开发流程 [⌅](#-documentation)

#### 1. 设计表结构, 创建物理表

- 设计表结构

  ```sql
  -- manifest/sql/gf2_demo.sql
  -- Create demo database
  CREATE DATABASE IF NOT EXISTS `gf2_demo`;

  USE `gf2_demo`;

  -- Create demo table
  DROP TABLE IF EXISTS `demo`;
  CREATE TABLE `demo`
  (
      `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
      `fielda`  varchar(45) NOT NULL COMMENT 'Field demo',
      `fieldb`  varchar(45) NOT NULL COMMENT 'Private field demo',
      `created_at` datetime DEFAULT NULL COMMENT 'Created Time',
      `updated_at` datetime DEFAULT NULL COMMENT 'Updated Time',
      PRIMARY KEY (`id`),
      UNIQUE KEY `idx_fielda` (`fielda`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
  ```

- 创建物理表

  ```sh
  $ mysql -uroot -p'123456' < manifest/sql/demo.sql
  ```

#### 2. 自动生成数据层相关代码

- gf 工具配置

  ```yaml
  # hack/config.yaml
  gfcli:
  gen:
      # doc: https://goframe.org/pages/viewpage.action?pageId=3673173
      dao:
      - link: "mysql:root:123456@tcp(127.0.0.1:3306)/gf2_demo"
          tables: "" # 指定当前数据库中需要执行代码生成的数据表, 多个以逗号分隔. 如果为空, 表示数据库的所有表都会生成. 默认为空
          descriptionTag: true # 用于指定是否为数据模型结构体属性增加desription的标签, 内容为对应的数据表字段注释. 默认 false
          noModelComment: true # 用于指定是否关闭数据模型结构体属性的注释自动生成, 内容为数据表对应字段的注释. 默认 false
          jsonCase: "snake" # 指定model中生成的数据实体对象中json标签名称规则. 默认 CamelLower
          clear: true # 自动删除数据库中不存在对应数据表的本地dao/do/entity代码文件, 默认 false. 线上环境应设置为fasle
  ```

- 自动生成 `internal/dao`, `internal/model/do`, `internal/model/entity`

  ```sh
  $ make dao
  ```

#### 3. 编写 api 层代码

位置: `api/demo/v1/demo.go`

定义业务侧数据结构, 提供对外接口的输入/输出数据结构, 定义访问路由 path, 请求方法, 数据校验, api 文档等.

注意: 目录结构必须遵守这个模式规范 `api/模块名称/版本号/模块名称.go`

示例:

```go
// api/demo/v1/demo.go

type CreateReq struct {
	g.Meta `path:"/demo" method:"post" tags:"DemoService" summary:"Create a demo record"`
	Fielda string `p:"fileda" v:"required|passport|length:4,30"`
	Fieldb string `p:"filedb" v:"required|length:10,30"`
}

type CreateRes struct {
	ID uint `json:"id"`
}
```

> 编写规范请参考文档: https://goframe.org/pages/viewpage.action?pageId=93880327

#### 4. 自动生成 controller 层框架代码

编写完 api 定义代码(`api/demo/v1/demo.go`)后, 在项目根目录执行如下命令行:

```sh
$ make ctrl
```

该命令行会根据开发者编写的 `api/demo/v1/demo.go` api 定义文件自动生成:

1. api interface 文件

   `api/demo/demo.go`

2. controller 层代码

   ```sh
   ├── internal
   │   ├── controller
   │   │   └── demo
   │   │       ├── demo.go
   │   │       ├── demo_new.go  # 不可变更
   │   │       ├── demo_v1_create.go  # 我们只需要在这里填充controller的具体实现
   ```

#### 5. 编写 model 层代码

位置: `internal/model/`

定义数据侧数据结构，提供对内的数据处理的输入/输出数据结构.
在 GoFrame 框架规范中, 这部分输入输出模型名称以 `XxxInput` 和 `XxxOutput` 格式命名, 需要在 `internal/model` 目录下创建文件.

示例:

```go
// internal/model/demo.go

type DemoCreateInput struct {
	Fielda string
	Fieldb string
}

type DemoCreateOutput struct {
	ID uint
}
```

> 参考: https://goframe.org/pages/viewpage.action?pageId=7295964

#### 6. 编写 service 层代码

##### a. 编写具体的业务实现(`internal/logic/`)

调用数据访问层(`internal/dao/`), 编写具体的业务逻辑. 这里是业务逻辑的重心, 绝大部分的业务逻辑都应该在这里编写.

示例:

```go
// internal/logic/demo/demo.go

import (
	"gf2-demo/internal/dao"
	"gf2-demo/internal/model"
	"gf2-demo/internal/model/do"
	"gf2-demo/internal/model/entity"
)

type sDemo struct{}

func New() *sDemo {
	return &sDemo{}
}

func (s *sDemo) Create(ctx context.Context, in model.DemoCreateInput) (*model.DemoCreateOutput, error) {
	notFound, err := s.FieldaNotFound(ctx, in.Fielda)
	if err != nil {
		return nil, err
	}
	if !notFound {
		err1 := gerror.WrapCode(codes.CodeNotAvailable, fmt.Errorf("fielda '%s' already exists", in.Fielda))
		return nil, err1
	}

	id, err := dao.Demo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.DemoCreateOutput{
		ID: uint(id),
	}, nil
}
```

##### b. 自动生成 service 接口代码(`internal/service/`)

```sh
$ make service
```

##### c. 将业务实现注入到服务接口(依赖注入)

示例:

```go
// internal/logic/demo/demo.go

import 	"gf2-demo/internal/service"

type sDemo struct{}

func init() {
	service.RegisterDemo(New())
}
```

##### d. 程序启动时自动注册服务

在程序入口文件 `cmd/gf2-demo-api/gf2-demo-api.go` 中导入 logic 包.

示例:

```go
// cmd/gf2-demo-api/gf2-demo-api.go

package main

import _ "gf2-demo/internal/logic"
```

> 参考: https://goframe.org/pages/viewpage.action?pageId=49770772

#### 7. 编写 controller 层代码

位置: `internal/controller/`

controller 代码文件前面已经通过`make ctrl`自动生成了, 我们只需要在适当的位置填充具体实现即可.

具体实现如何编写:

解析 api 层(`api/demo/v1/demo.go`)定义的业务侧用户输入数据结构, 组装为 model 层(`internal/model/`)定义的数据侧输入数据结构实例, 调用 `internal/service/` 层的服务, 最后直接将结果或错误 return 即可(响应中间件会统一拦截处理, 按规范响应用户).

示例:

```go
// internal/controller/demo/demo_v1_create.go

import (
	"context"

	v1 "gf2-demo/api/demo/v1"
	"gf2-demo/internal/model"
	"gf2-demo/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	data := model.DemoCreateInput{
		Fielda: req.Fielda,
		Fieldb: req.Fieldb,
	}

	result, err := service.Demo().Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{ID: result.ID}, err
}
```

#### 8. 路由注册

位置: `internal/cmd/apiserver/apiserver.go`

路由注册, 调用 controller 层(`internal/controller/`), 对外暴露接口.

示例:

```go
// internal/cmd/apiserver/apiserver.go

import	"gf2-demo/internal/controller/demo"

var (
	Main = gcmd.Command{

		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/v1", func(group *ghttp.RouterGroup) {
				group.Bind(
					// 对外暴露的接口路由集合: 是 "/v1" 和 api/demo/v1/demo.go 文件中的所有XxxReq结构体定义的 path 的组合
					demo.NewV1(),
				)
			})

			s.Run()
			return nil
		},
	}
)
```

#### 9. 接口访问测试

```sh
# 运行
$ make run

# 访问
$ curl -X POST -i 'localhost:9000/v1/demo' -d \
'{
    "fielda": "foobar",
    "fieldb": "LJIYdjsvt83l"
}'

# 输出:

HTTP/1.1 200 OK
Content-Type: application/json
Server: GoFrame HTTP Server
Trace-Id: 0862d01e5aa64317c7fae45b326dabd1
Date: Tue, 14 Feb 2023 09:19:52 GMT
Content-Length: 88

{"code":"OK","message":"","traceid":"0862d01e5aa64317c7fae45b326dabd1","data":{"id":17}}
```

### 代码质量 [⌅](#-documentation)

统一团队代码风格, 保障团队代码质量.

> Github: https://github.com/golangci/golangci-lint <br>
> Documentation: https://golangci-lint.run

#### 安装 golangci-lint

```sh
# 进行代码检查, 如果 golangci-lint 没有安装会自动进行安装
$ make lint
```

或

```sh
$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

#### 使用方法

1. 命令行执行

   ```sh
   # 在仓库根路径执行, 检测仓库内所有Go代码
   $ golangci-lint run
   # 或
   $ make lint

   # 查看所有linters的功能介绍
   $ golangci-lint help linters

   # 查看 .golangci.yml 已启用的 linters
   $ golangci-lint linters

   # 只使用某一个linter来检查代码
   $ golangci-lint run --no-config --disable-all -E errcheck
   ```

2. 集成到编辑器或 IDE

   请参考官方文档: `https://golangci-lint.run/usage/integrations/`

   强烈建议使用此种方式, 可实时提示代码存在的问题, 而不是等到编译的时候才知道哪里出错了,
   不但提高代码质量, 还能提高编码效率.

### 项目部署 [⌅](#-documentation)

#### Systemctl

1. 相关的配置文件及脚本

   - 生产环境 systemctl 服务配置文件: `manifest/deploy/systemctl/gf2-demo-api.service`
   - 测试环境 systemctl 服务配置文件: `manifest/deploy/systemctl/gf2-demo-api_test.service`
   - 部署脚本: `manifest/deploy/systemctl/deploy.sh`

2. 设置目标服务器(修改 `deploy.sh` 脚本)

   ```sh
   # 用于连接目标服务器的用户名.
   # NOTE: 不要使用个人账号部署和运行服务, 建议建立独立的公共账号部署和运行应用.
   deploy_user="vagrant"
   # 目标部署服务器.
   # NOTE:
   # 1. 请提前配置发布机到目标服务器之间的ssh key信任:
   #    1) ssh-keygen -t rsa -b 4096 -C "vagrant@$(hostname)"
   #    2) ssh-copy-id -i ~/.ssh/id_rsa.pub vagrant@$(deploy_server)
   # 2. 目标机器上解决sudo执行命令提示输入密码的问题:
   #    执行visudo, 末尾添加一行: vagrant ALL=(ALL:ALL) NOPASSWD:ALL
   deploy_server="gf2-demo.sre.im"
   # 部署目录的权限应该为$deploy_user: chown -R $deploy_user. $deploy_dir
   deploy_dir=/app/$project_name
   ```

3. 执行部署

   ```sh
   # 部署测试环境
   $ ./manifest/deploy/systemctl/deploy.sh test

   # 部署生产环境
   $ ./manifest/deploy/systemctl/deploy.sh prod
   ```

4. 验证

   首先登录到目标服务器.

   ```sh
   # 默认项目的所有标准输出和标准错误输出都会在此文件中.
   $ tail -f /app/gf2-demo/gf2-demo-api.log

   # 项目常规日志, 包括通过g.Log()打印的日志.
   $ tail -f /app/gf2-demo/logs/2023-02-15.log

   # 项目HTTP服务访问日志
   $ tail -f /app/gf2-demo/logs/access-20230215.log

   # 项目HTTP服务错误日志
   $ tail -f /app/gf2-demo/logs/error-20230215.log

   # sql debug 日志
   $ tail -f /app/gf2-demo/logs/sql-20230215.log
   ```

5. systemctl 常用命令

   ```sh
   # gf2-demo-api.service 配置有变动的时候, 需要重新加载使生效
   $ sudo systemctl daemon-reload

   # 启动
   $ sudo systemctl start gf2-demo-api

   # 关闭: 发送 SIGTERM 信号给主(sh)和子进程(gf2-demo-api),
   # gf2-demo-api程序可通过捕获SIGTERM信号来实现优雅关闭.
   $ sudo systemctl stop gf2-demo-api

   # 重启: 先关闭(SIGTERM), 再启动
   $ sudo systemctl restart gf2-demo-api
   ```

> NOTE:
>
> - 此示例为单台部署, 若部署集群可使用 `gossh`、`ansible` 等工具.
> - 服务器操作系统: `CentOS7.x`, 其他系统类型未验证.

#### Supervisor

1. 相关的配置文件及脚本

   - 生产环境 supervisor 配置文件: `manifest/deploy/supervisor/gf2-demo-api.ini`
   - 测试环境 supervisor 服务配置文件: `manifest/deploy/supervisor/gf2-demo-api_test.ini`
   - 部署脚本: `manifest/deploy/supervisor/deploy.sh`

2. 设置目标服务器(修改 `deploy.sh` 脚本)

   ```sh
   deploy_server="gf2-demo.sre.im"
   deploy_user="vagrant"
   deploy_dir=/app/$project_name
   ```

   相关说明参考 systemctl 部署示例.

3. 在目标服务器上提前安装 supervisor

   基于 CentOS7 系统演示.

   ```sh
   yum update -y
   yum install epel-release -y
   yum install supervisor -y

   systemctl enable supervisord
   systemctl start supervisord
   systemctl status supervisord

   echo_supervisord_conf > /etc/supervisord.conf
   cat >> /etc/supervisord.conf <<EOF
   [include]
   files = supervisord.d/*.ini
   EOF

   mkdir -p /etc/supervisord.d
   ```

4. 执行部署

   ```sh
   # 部署测试环境
   $ ./manifest/deploy/supervisor/deploy.sh test

   # 部署生产环境
   $ ./manifest/deploy/supervisor/deploy.sh prod
   ```

5. supervisorctl 常用命令

   ```sh
   # 启动
   $ sudo supervisorctl start gf2-demo-api

   # 关闭(SIGTERM信号), 可捕获SIGTERM信号, 实现优雅关闭
   $ sudo supervisorctl stop gf2-demo-api

   # 重启: 先关闭(SIGTERM信号), 再启动.
   # NOTE: /etc/supervisord.*相关配置有变动, 重启具体某服务并不会生效
   $ sudo supervisorctl restart gf2-demo-api

   # 重启 supervisor 控制的所有服务.
   # NOTE: 当 /etc/supervisord.*相关配置有变动, 必须执行此命令才能加载生效
   $ sudo supervisorctl reload
   ```

#### Docker

1. Dockerfile

   采用两阶段构建, 镜像体积小; 将依赖库下载剥离出来并且前置, 利用缓存特性提高编译速度.

   ```dockerfile
   # syntax=docker/dockerfile:1

   # Step 1: build binary
   FROM golang:1.17 as builder

   ENV GOPROXY https://goproxy.cn,direct

   WORKDIR /src

   # pre-copy/cache go.mod for pre-downloading dependencies and
   # only redownloading them in subsequent builds if they change
   COPY Makefile ./
   RUN make cli
   COPY go.mod go.sum ./
   RUN go mod download && go mod verify

   COPY . .
   RUN make build OS="linux"


   # Step 2: copy binary from step 1
   FROM loads/alpine:3.8

   ENV GF_GERROR_BRIEF=true

   WORKDIR /app

   COPY --from=builder /src/bin/linux_amd64/gf2-demo-api .

   EXPOSE 9000

   ENTRYPOINT [ "./gf2-demo-api" ]
   ```

2. 制作容器镜像

   ```sh
   $ cd gf2-demo

   $ make image

   $ docker image ls

   REPOSITORY          TAG                            IMAGE ID       CREATED          SIZE
   gf2-demo-api        20230221113306.0d26121.dirty   58e6953c2e1b   15 seconds ago   30.1MB
   ```

3. 运行容器

   ```sh
   # 开发环境
   $ docker run --name gf2-demo -p80:9000 -d gf2-demo-api:20230221113306.0d26121.dirty

   # 测试环境
   $ docker run --name gf2-demo -p80:9000 -e GF_GCFG_FILE=config.test.yaml -d gf2-demo-api:20230221113306.0d26121.dirty

   # 生产环境
   $ docker run --name gf2-demo -p80:9000 -e GF_GCFG_FILE=config.prod.yaml -d gf2-demo-api:20230221113306.0d26121.dirty
   ```

4. 验证

- 查看是否成功运行:  
  浏览器访问 `http://localhost/swagger`, 参看 api 文档是否正常展示.

- 查看二进制应用版本信息

  ```sh
  $ docker exec -it gf2-demo ./gf2-demo-api -v

  # 输出如下:
  App Version: v0.7.0
  Git Commit:  2023-02-17 19:32:05 95390e39485aa29050c2327c263a732267ec3eb3
  Build Time:  2023-02-20 06:18:57
  Go Version:  go1.17.13
  GF Version:  v2.3.2
  ```

- 查看不同环境下, 程序使用的配置文件是否正确

  ```sh
  # 查看容器输出的日志
  $ docker logs gf2-demo

  # 如果配置了日志保存到文件, 也可登录到容器内部进行查看.
  $ docker exec -it gf2-demo sh

  # 输出的部分日志截取:
  2023-02-17 18:52:36.568 [DEBU] {7f0f8d5a279744179740f477f49fbd06} /Users/xg/github/gf2-demo/internal/cmd/apiserver/apiserver.go:79: use config file: &{defaultName:config searchPaths:0xc0000bf6e0 jsonMap:0xc000303720 violenceCheck:false}
  ```

  上面日志中的 `defaultName` 如果为 `config`, 代表开发环境; 为 `config.test.yaml`, 代表测试环境; 为 `config.prod.yaml`, 代表生产环境.

5. 如何优雅关闭

   ```sh
   # 关闭: 会发送SIGTERM信号, gf2-demo捕获该信号经过处理, 可实现优雅关闭
   $ docker stop gf2-demo

   # 重启: 先关闭(SIGTERM信号), 再启动, 可实现优雅关闭
   $ docker restart gf2-demo

   # 强制关闭(SIGKILL信号), gf2-demo无法捕获到SIGKILL信号, 直接退出
   $ docker kill gf2-demo

   # 强制关闭并删除容器(SIGKILL信号)
   $ docker rm -f gf2-demo
   ```

#### 优雅关闭测试

`GoFrame` 从 `V2.4.0` 版本开始已支持捕获`SIGTERM`信号来实现优雅关闭服务. 以上三种部署方式(`Systemctl`/`Supervisor`/`Docker`)在关闭或重启服务的时候均是发送`SIGTERM`信号给服务进程, 所以都能优雅关闭服务.

开始测试, 准备至少 3 个 terminal 窗口.

1. 模拟接口响应延时 8 秒

   测试接口: `GET localhost:9000/v1/demo/:fielda`

   ```go
   // internal/controller/demo.go

   func (c *cDemo) Get(ctx context.Context, req *v1.DemoGetReq) (*v1.DemoGetRes, error) {
   // 加入此行代码, 模拟延迟
   time.Sleep(8 * time.Second)

   demoInfo, err := service.Demo().Get(ctx, req.Fielda)
   if err != nil {
       return nil, err
   }
   ```

2. 配置文件中的`server.gracefulShutdownTimeout`调整为 10 秒

   ```yaml
   server:
   gracefulShutdownTimeout: 10 # 默认 5秒, 建议根据实际业务情况调整
   ```

3. 在窗口 1 启动服务

   ```sh
   $ make run
   ```

   在输出日志中找到服务 PID 为 80273, 在步骤 5 会用到:

   ```text
   2023-04-25 11:31:16.716 [INFO] pid[80273]: http server started listening on [:9000]
   ```

4. 在窗口 2 模拟请求

   ```sh
   $ curl localhost:9000/v1/demo/windvalley
   ```

   请求卡住, 在等待响应中.

5. 在窗口 3 模拟优雅关闭服务

   执行完第 4 步后, 立即执行如下命令:

   ```sh
   $ kill -SYSTERM 80273
   ```

   窗口 1 输出:

   ```text
   2023-04-25 11:41:52.004 80273: server gracefully shutting down by signal: terminated
   ```

   此时, 并没有马上关闭服务, 在等待请求处理完成.
   几秒钟后, 请求处理完成, 服务关闭, 窗口 1 继续输出:

   ```test
   2023-04-25 11:41:57.863 [DEBU] {708a5f7a8710591764de0d572ec0fc19} [  1 ms] [default] [gf2_demo] [rows:1  ] SELECT * FROM `demo` WHERE `fielda`='windvalley' LIMIT 1
   2023-04-25 11:41:57.863 {708a5f7a8710591764de0d572ec0fc19} 200 "GET http localhost:9000 /v1/demo/windvalley HTTP/1.1" 8.005, 127.0.0.1, "", "curl/7.79.1"
   2023-04-25 11:41:58.275 [INFO] pid[80273]: all servers shutdown
   ```

   此时窗口 2 输出:

   ```text
   {"code":"OK","message":"","traceid":"708a5f7a8710591764de0d572ec0fc19","data":{"id":14,"fielda":"windvalley","created_at":"2023-02-14 17:12:59","updated_at":"2023-02-14 17:12:59"}}
   ```

   请求成功.

> NOTE:
>
> 虽然 `GoFrame` 支持优雅重启特性, 但在生产环境下不建议开启:
>
> 1. `Systemctl` 或 `Supervisor` 无法很好的接管和控制父子进程. 优雅关闭特性已足够, 可通过部署负载均衡器(LVS 等)来实现不中断服务.
>
> 2. 对于`Docker`/`K8S`等容器化场景, 最佳实践要求进程和容器本身同生命周期, 更不能开启优雅重启特性. 容器管理平台本身支持容器的平滑启停, 实现不中断服务.

### 使用 Makefile 管理项目 [⌅](#-documentation)

```sh
$ make help

Usage:

    make [TARGETS] [OPTIONS]

Targets:

    cli          Install/Update to the latest Gf Cli tool
    lint         Run golangci-lint
    ctrl         Generate Go files for Controller
    dao          Generate Go files for Dao/Do/Entity
    service      Generate Go files for Service
    run          Run gf2-demo-api for development environment
    run.cli      Run gf2-demo-cli for development environment
    build        Build gf2-demo-api binary
    build.cli    Build gf2-demo-cli binary
    image        Build docker image
    image.push   Build docker image and automatically push to docker repo
    help         Show this help

Options:

    ARCH         The multiple ARCH to build. Default is "amd64";
                 This option is available for: make build*;
                 Example: make build ARCH="amd64,arm64"

    OS           The multiple OS to build. Default is "darwin,linux";
                 This option is available for: make build*;
                 Example: make build OS="linux,darwin,windows"

    TAG          Docker image tag. Default is generated from current git commit;
                 This option is available for: make image*;
                 Example: make image TAG="v0.6.0"
```

使用示例:

```sh
# 安装最新版gf
$ make cli

# 运行 golangci-lint 检查代码
$ make lint

# 物理表有增加或表结构有更新时, 自动生成或更新数据层相关代码
$ make dao

# 自动生成 internal/controller/ 控制器层代码
$ make ctrl

# internal/logic/ 有代码变动后, 使用此命令自动生成 internal/service/ 接口代码
$ make service

# 开发环境热启动 gf2-demo-api
$ make run

# 开发环境热启动 gf2-demo-cli
$ make run.cli

# 编译 gf2-demo-cli
$ make build.cli

# 按照 hack/config.yaml 配置, 编译 gf2-demo-api
$ make build

# 编译 linux_arm64 架构的二进制文件
$ make build OS="linux" ARCH="arm64"

# 制作docker镜像: gf2-demo-api:20230220144044.2d4bb19.dirty
$ make image

# 自定义tag制作docker镜像: gf2-demo-api:v0.3.0
$ make image TAG="v0.3.0"

# 制作镜像并推送到docker仓库
$ make image.push
```

> NOTE:
> 如果是 macOS 系统, 需要提前安装 `gsed` 命令.

### 变更项目名称 [⌅](#-documentation)

请按如下步骤便捷地将本项目名称 `gf2-demo` 改成你自己的项目名称 `new-project`.

1. 变更项目目录名称

   ```sh
   $ mv gf2-demo new-project
   ```

2. 运行变更脚本

   ```sh
   $ cd new-project
   $ hack/change_project_name.sh new-project
   ```

   > NOTE:
   >
   > 1. 如果是 macOS 系统, 需要提前安装 `gsed` 命令. <br>
   > 2. mysql 数据库配置需要根据你的实际情况进行手动修改, 有以下几处：<br>
   >    - `manifest/config/*.yaml` <br>
   >    - `hack/config.yaml` <br>
   >    - `manifest/sql/gf2_demo.sql`

3. 验证

   ```sh
   $ make build

   输出如下:

   bin
   ├── darwin_amd64
   │   └── new-project-api
   └── linux_amd64
       └── new-project-api
   ```

## 📜 References

- https://github.com/gogf/gf
- https://goframe.org/display/gf
- https://pkg.go.dev/github.com/gogf/gf/v2

## ⚖️ License

This project is under the MIT License.
See the [LICENSE](LICENSE) file for the full license text.
