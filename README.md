# gf2-demo

`gf2-demo` 是一个基于 [GoFrameV2](https://github.com/gogf/gf) 用来快速开发后端服务的脚手架, 目标使开发者只需关注业务逻辑的编写, 快速且规范地交付项目.

## 💌 Features

- 优化工程目录结构, 使支持多个可执行命令
- 规范业务错误码, 中间件统一拦截响应, 规范响应格式
- 完善 HTTP 服务访问日志、HTTP 服务错误日志、开发者打印的日志、其他可执行命令的日志配置
- 多环境管理: 开发环境、测试环境、生产环境
- 项目的二进制文件可打印当前版本信息
- 链路跟踪中间件, 默认使用客户端按规范传递的`X-Request-Id`
- 通过 Makefile 管理项目: `make run`, `make run.cli`, `make build`, `make build.cli` 等

## 🚀 Quick Start

### 安装

请提前安装 Go 环境, 要求 Go 版本: `1.15+`

```sh
git clone --depth 1 git@github.com:windvalley/gf2-demo.git

cd gf2-demo

# 安装gf
make cli
```

### 热更新(Live reload)

```sh
cd gf2-demo

# 运行 gf2-demo-api
make run

# 运行 gf2-demo-cli
make run.cli
```

> 默认加载配置文件: `manifest/config/config.yaml`

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

- [工程目录](#工程目录)
- [环境管理](#环境管理)
  - [开发环境](#开发环境)
  - [测试环境](#测试环境)
  - [生产环境](#生产环境)
- [多命令管理](#多命令管理)
  - [目录设计](#目录设计)
  - [配置文件](#配置文件)
- [错误码管理](#错误码管理)
  - [规范制定](#规范制定)
  - [维护业务错误码](#维护业务错误码)
  - [测试效果](#测试效果)
- [日志管理](#日志管理)
  - [HTTP 服务日志](#HTTP-服务日志)
    - [1. HTTP 服务日志配置](#1-HTTP-服务日志配置)
    - [2. 生成的日志示例](#2-生成的日志示例)
  - [开发者打印的通用日志](#开发者打印的通用日志)
    - [1. 通用日志配置](#1-通用日志配置)
    - [2. 如何打印日志](#2-如何打印日志)
    - [3. 生成的日志示例](#3-生成的日志示例)
- [链路跟踪](#链路跟踪)
- [版本管理](#版本管理)
- [开发流程](#开发流程)
  - [1. 设计表结构, 创建物理表](#1-设计表结构-创建物理表)
  - [2. 自动生成数据层相关代码](#2-自动生成数据层相关代码)
  - [3. 编写 api 层代码](#3-编写-api-层代码)
  - [4. 编写 model 层代码](#4-编写-model-层代码)
  - [5. 编写 service 层代码](#5-编写-service-层代码)
  - [6. 编写 controller 层代码](#6-编写-controller-层代码)
  - [7. 路由注册](#7-路由注册)
- [项目部署](#项目部署)
- [使用 Makefile 管理项目](#使用-Makefile-管理项目)

### 工程目录

```sh
├── api  # 对外接口定义: 对外提供服务的输入/输出数据结构定义, 路由path定义, 数据校验等
│   └── v1
│       └── hello.go
├── bin  # make build 和 make build.cli 生成的二进制可执行文件所在目录
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
│   ├── controller  # 对外接口实现: 接收/解析用户输入参数的入口/接口层
│   │   └── hello.go
│   ├── dao  # 数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 CURD 方法. dao层通过框架的ORM抽象层组件与底层真实的数据库交互
│   ├── logic  # 业务封装: 业务逻辑封装管理, 特定的业务逻辑实现和封装. 往往是项目中最复杂的部分. logic层的业务逻辑需要通过调用dao来实现数据的操作, 调用dao时需要传递do数据结构对象, 用于传递查询条件、输入数据. dao执行完毕后通过Entity数据模型将数据结果返回给service(logic)层
│   │   ├── logic.go
│   │   └── middleware  # 中间件
│   │       ├── accessuser.go
│   │       ├── middleware.go
│   │       ├── response.go
│   │       └── traceid.go
│   ├── model  # 数据结构管理模块, 管理数据实体对象, 以及输入与输出数据结构定义. 这里的model不仅负责维护数据实体对象(entity)结构定义, 也包括所有的输入/输出数据结构定义, 被api/dao/service共同引用
│   │   ├── do  # 领域对象: 用于dao数据操作中业务模型与实例模型转换, 由工具维护, 用户不能修改
│   │   └── entity  # 数据模型: 是模型与数据集合的一对一关系, 通常和数据表一一对应, 由工具维护, 用户不能修改
│   ├── packed
│   │   └── packed.go
│   └── service  # 业务接口: 用于业务模块解耦的接口定义层. 具体的接口实现在logic中进行注入
│       └── middleware.go
├── manifest  # 交付清单: 包含程序编译、部署、运行、配置的文件
│   ├── config  # 配置文件存放目录
│   │   ├── config.prod.yaml  # 生产环境
│   │   ├── config.test.yaml  # 测试环境
│   │   └── config.yaml  # 开发环境
│   ├── deploy  # 和部署相关的文件
│   │   └── kustomize
│   │       ├── base
│   │       │   ├── deployment.yaml
│   │       │   ├── kustomization.yaml
│   │       │   └── service.yaml
│   │       └── overlays
│   │           └── develop
│   │               ├── configmap.yaml
│   │               ├── deployment.yaml
│   │               └── kustomization.yaml
│   └── docker  # Docker镜像相关依赖文件, 脚本文件等等
│       ├── Dockerfile
│       └── docker.sh
├── resource  # 静态资源文件: 这些文件往往可以通过资源打包/镜像编译的形式注入到发布文件中
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

### 环境管理

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
> - `GF_GERROR_BRIEF=true` 表示 HTTP 服务日志错误堆栈中不包含 gf 框架堆栈.
> - 配置文件在通过 `make build` 或 `make build.cli` 编译时已经打包到二进制文件中, 所以在部署时只需部署二进制文件即可.

#### 生产环境

配置文件: `manifest/config/config.prod.yaml`

运行:

同测试环境, 只不过指定的配置文件不同, 略.

### 多命令管理

#### 目录设计

举例:

命令 1: `cmd/gf2-demo-api/gf2-demo-api.go` -> `internal/cmd/apiserver/apiserver.go`

命令 2: `cmd/gf2-demo-cli/gf2-demo-cli.go` -> `internal/cmd/cli/cli.go`

#### 配置文件

默认不同命令在相同环境下使用同一个配置文件, 比如 `gf2-demo-api` 和 `gf2-demo-cli` 在开发环境下都使用 `manifest/config/config.yaml` 作为配置文件.

不过也可以使用各自独立的配置文件, 只需要在运行时通过环境变量或命令行参数指定需要使用的配置文件即可, 比如:

`./gf2-demo-cli -c cli_config.yaml` 或
`GF_GCFG_FILE=cli_config.yaml ./gf2-demo-cli`

### 错误码管理

#### 规范制定

- 统一响应格式

  不论是正确还是错误响应, 响应体都统一使用如下格式:

```json
{
  "code": "string",
  "msg": "string",
  "data": null
}
```

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

#### 维护业务错误码

请在 `internal/codes/biz_codes.go` 文件中维护业务错误码.

```go
package codes

// http status, bisiness code, message
var (
	CodeOK          = New(200, "OK", "")
	CodePartSuccess = New(202, "PartSuccess", "part success")

	CodeNotAuthorized    = New(401, "NotAuthorized", "resource is not authorized")
	CodePermissionDenied = New(403, "PermissionDenied", "permission denied")
	CodeNotFound         = New(404, "NotFound", "resource does not exist")
	CodeValidationFailed = New(400, "ValidationFailed", "validation failed")

	CodeInternal = New(500, "InternalError", "an error occurred internally")
	CodeUnknown  = New(500, "UnknownError", "unknown error")
)
```

#### 测试效果

```sh
curl --location --request GET 'localhost:9000/v1/hello' -i
```

- 正确响应

```text
HTTP/1.1 200 OK
Content-Type: application/json
Server: GoFrame HTTP Server
Trace-Id: 10c9769ce5cf4117c19a595c2d781e94
Date: Wed, 08 Feb 2023 09:38:41 GMT
Content-Length: 34

{"code":"OK","data":null,"msg":""}
```

- 401 错误

```text
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Server: GoFrame HTTP Server
Trace-Id: a89b7652b1cf41170d6e5233fbb76a21
Date: Wed, 08 Feb 2023 09:34:56 GMT
Content-Length: 83

{"code":"NotAuthorized","data":null,"msg":"resource is not authorized: some error"}
```

- 500 错误

```text
HTTP/1.1 500 Internal Server Error
Content-Type: application/json
Server: GoFrame HTTP Server
Trace-Id: 70cd58a9d8cf4117376a265eb84137e5
Date: Wed, 08 Feb 2023 09:37:45 GMT
Content-Length: 73

{"code":"InternalError","data":null,"msg":"an error occurred internally"}
```

### 日志管理

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
    ctxKeys: ["user", "mail"] # 自动打印Context的指定变量到日志中. 默认为空
    rotateExpire: "1d"
    rotateBackupExpire: "30d"
    rotateBackupLimit: 30
    rotateCheckInterval: "1h"
```

##### 2. 生成的日志示例

```sh
curl --location --request GET 'localhost:9000/v1/hello' \
    --header 'X-Consumer-Custom-ID: windvalley' \
    --header 'X-Consumer-Username: windvalley@sre.im'
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

#### 开发者打印的通用日志

##### 1. 通用日志配置

```yaml
# manifest/config/config.yaml
logger:
  path: "logs/" # 日志文件目录, 如果为空, 表示不记录到文件
  file: "{Y-m-d}.log" # 日志文件格式. 默认为"{Y-m-d}.log"
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
    file: "cli_{Y-m-d}.log"
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
g.Log().Errorf(ctx, "xxx failed")
// gf2-demo-cli的日志
g.Log("cli").Debug(ctx, "hello world")
g.Log("cli").Warningf(ctx, "warning message")
```

##### 3. 生成的日志示例

```sh
# 普通格式
2023-02-08 17:02:31.906 [INFO] {389b4e7aeccd41175dd0bc18211c2519} {windvalley, windvalley@sre.im} /Users/xg/github/gf2-demo/internal/controller/hello.go:33: hello world
2023-02-08 17:02:31.906 [ERRO] {389b4e7aeccd41175dd0bc18211c2519} {windvalley, windvalley@sre.im} /Users/xg/github/gf2-demo/internal/controller/hello.go:34: xxx failed

# json格式
{"Time":"2023-02-08 17:04:08.957","TraceId":"d0e7f61203ce41171374033689322f91","CtxStr":"windvalley, windvalley@sre.im","Level":"INFO","CallerPath":"/Users/xg/github/gf2-demo/internal/controller/hello.go:33:","Content":"hello world"}
{"Time":"2023-02-08 17:04:08.957","TraceId":"d0e7f61203ce41171374033689322f91","CtxStr":"windvalley, windvalley@sre.im","Level":"ERRO","CallerPath":"/Users/xg/github/gf2-demo/internal/controller/hello.go:34:","Content":"xxx failed"}
```

### 链路跟踪

- 用于链路跟踪的响应 Header 为: `Trace-Id`, 会优先使用客户端传递的请求 Header `X-Request-Id` 的值作为 `Trace-Id` 的值, 如果不存在会自动生成.

- 服务内部如果需要调用其他服务的接口, 请使用 `g.Client()`, 因为他会给请求自动注入`Trace-Id`, 这样不同 API 服务之间的日志就可以通过 `Trace-Id` 串起来了.

> NOTE:
>
> - 浏览器请求时会自动携带 Header: `X-Request-Id`
> - 请参考文档: https://goframe.org/pages/viewpage.action?pageId=49745257

### 版本管理

1. 给项目仓库打 tag

```sh
git tag v0.3.0
git push --tags
```

2. 使用 Makefile 编译

```sh
# For gf2-demo-api
make build

# For gf2-demo-cli
make build.cli
```

3. 查看二进制文件版本信息

```sh
$ bin/darwin_amd64/gf2-demo-api -v

输出:

App Version: v0.3.0-7-g1898e82
Git Commit:  2023-02-08 14:55:39 1898e82dbcb4c2e8a091eb12fc96ead2f04f5993
Build Time:  2023-02-08 15:31:20
Go Version:  go1.17.6
GF Version:  v2.3.1
```

### 开发流程

#### 1. 设计表结构, 创建物理表

#### 2. 自动生成数据层相关代码

1. gf 工具配置

```yaml
# hack/config.yaml
gfcli:
  gen:
    dao:
      - link: "mysql:root:123456@tcp(127.0.0.1:3306)/gf2-demo"
```

2. 自动生成 `internal/dao`, `internal/model/do`, `internal/model/entity`

```sh
make dao
```

#### 3. 编写 api 层代码

位置: `api/v1/`.

定义业务侧数据结构, 提供对外接口的输入/输出数据结构, 定义访问路由 path, 请求数据校验, api 文档等.

#### 4. 编写 model 层代码

位置: `internal/model/`

定义数据侧数据结构，提供对内的数据处理的输入/输出数据结构.
在 GoFrame 框架规范中, 这部分输入输出模型名称以 `XxxInput` 和 `XxxOutput` 格式命名, 需要在 `internal/model` 目录下创建文件.

> 参考: https://goframe.org/pages/viewpage.action?pageId=7295964

#### 5. 编写 service 层

1. 编写具体的业务实现(`internal/logic/`)

调用数据访问层(`internal/dao/`), 编写具体的业务逻辑.

这里是业务逻辑的重心, 绝大部分的业务逻辑都应该在这里编写.

2. 自动生成 service 接口(`internal/service/`)

```sh
make service
```

3. 将业务实现注入到服务接口(依赖注入)

拿中间件举例:

`internal/logic/middleware/middleware.go`

```go
package middleware

import "gf2-demo/internal/service"

type (
	sMiddleware struct{}
)

// 服务注册/依赖注入
func init() {
	service.RegisterMiddleware(new())
}

func new() *sMiddleware {
	return &sMiddleware{}
}
```

4. 程序启动后自动注册服务

在程序入口文件中 `cmd/gf2-demo-api/gf2-demo-api.go` 倒入 logic 包.

```go
package main

import _ "gf2-demo/internal/logic"
```

> 参考: https://goframe.org/pages/viewpage.action?pageId=49770772

#### 6. 编写 controller 层

位置: `internal/controller/`

解析 api 层(`api/v1/`)定义的业务侧用户输入数据结构, 组装为 model 层(`internal/model/`)定义的数据侧输入数据结构实例, 调用 `internal/service/` 层的服务, 最后直接将结果或错误 return 即可(响应中间件会统一拦截处理, 按规范响应用户).

#### 7. 路由注册

位置: `internal/cmd/apiserver/`

路由分组注册, 调用 controller 层(`internal/controller/`), 对外暴露接口.

### 项目部署

> 参考: https://goframe.org/pages/viewpage.action?pageId=1114403

### 使用 Makefile 管理项目

> NOTE:
>
> MacOS 环境下, 需要安装 gsed 命令.

```sh
# 安装最新版gf
make cli

# 物理表有增加或表结构有更新时, 自动生成或更新数据层相关代码
make dao

# internal/logic/ 有代码变动后, 使用此命令自动生成 internal/service/ 接口代码
make service

# 开发环境热启动 gf2-demo-api
make run

# 开发环境热启动 gf2-demo-cli
make run.cli

# 编译 gf2-demo-api
make build

# 编译 gf2-demo-cli
make build.cli
```

## 📜 References

- https://goframe.org/display/gf
- https://pkg.go.dev/github.com/gogf/gf/v2
