# gf2-demo

`gf2-demo` 是一个基于 [GoFrameV2](https://github.com/gogf/gf) 快速开发后端服务的脚手架, 目标使开发者只需关注业务逻辑的编写, 快速交付项目.

## 💌 Features

- 优化目录结构, 使支持多个可执行命令
- 规范业务错误码, 中间件统一拦截响应, 规范响应格式
- 完善 HTTP 服务访问日志、HTTP 服务错误日志、开发者打印的日志、其他可执行命令的日志配置
- 多环境配置: 开发环境、测试环境、生产环境
- 可输出二进制文件的版本信息
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

### 开发环境

#### 配置文件

`manifest/config/config.yaml`

#### 运行

代码有变动准实时生效.

```sh
# 运行 gf2-demo-api
make run

# 运行 gf2-demo-cli
make run.cli
```

### 测试或生产环境

#### 配置文件

- 测试环境: `manifest/config/config.test.yaml`
- 生产环境: `manifest/config/config.prod.yaml`

#### 编译

```sh
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

#### 运行

运行 `gf2-demo-api` 和 `gf2-demo-cli` 类似, 下面以 `gf2-demo-api` 为例说明.

- 通过命令行参数指定配置文件

```sh
cd bin/darwin_amd64/

# 测试
./gf2-demo-api -c config.test.yaml
# 生产
./gf2-demo-api -c config.prod.yaml
```

- 通过环境变量指定配置文件

```sh
cd bin/darwin_amd64/

# 测试
GF_GCFG_FILE=config.test.yaml GF_GERROR_BRIEF=true ./gf2-demo-api
# 生产
GF_GCFG_FILE=config.prod.yaml GF_GERROR_BRIEF=true ./gf2-demo-api
```

> NOTE:
>
> - 通过命令行参数指定配置文件优先于环境变量.
> - 直接运行 `./gf2-demo-api` 或 `./gf2-demo-cli` 将默认使用 `config.yaml` 配置文件.
> - `GF_GERROR_BRIEF=true` 表示 HTTP 服务日志错误堆栈中不包含 gf 框架堆栈.
> - 配置文件在通过 `make build` 或 `make build.cli` 编译时已经打包到二进制文件中, 所以在部署时只需部署二进制文件即可.

## 📄 Documentation

### 工程目录

```sh
├── api  # 对外接口定义: 对外提供服务的输入/输出数据结构定义, 路由path定义, 数据校验等
│   └── v1
│       └── hello.go
├── bin  # make build 和 make build.cli 生成的二进制可执行文件
│   ├── darwin_amd64
│   │   ├── gf2-demo-api
│   │   └── gf2-demo-cli
│   └── linux_amd64
│       ├── gf2-demo-api
│       └── gf2-demo-cli
├── cmd  # 项目的可执行文件入口
│   ├── gf2-demo-api
│   │   └── gf2-demo-api.go  # 编译时会使用入口文件的名字作为二进制文件名称
│   └── gf2-demo-cli
│       └── gf2-demo-cli.go
├── hack  # 存放项目开发工具、脚本等内容. 例如: gf工具的配置, 各种shell/bat脚本等文件
│   └── config.yaml  # gf 工具的配置文件, gf gen/gf build 等会使用这里的配置内容
├── internal
│   ├── cmd  #  对应外层 cmd 目录
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

### 错误码管理

#### 规则制定

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

// http status, bisiness codes, message
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
  logPath: "logs/" # 日志文件存储目录路径，建议使用绝对路径。默认为空，表示关闭
  logStdout: true # 日志是否输出到终端。默认为true
  errorStack: true # 当Server捕获到异常时是否记录堆栈信息到日志中。默认为true
  errorLogEnabled: true # 是否记录异常日志信息到日志中。默认为true
  errorLogPattern: "error-{Ymd}.log" # 异常错误日志文件格式。默认为"error-{Ymd}.log"
  accessLogEnabled: true # 是否记录访问日志(包含异常的访问日志)。默认为false
  accessLogPattern: "access-{Ymd}.log" # 访问日志文件格式。默认为"access-{Ymd}.log"

  # 针对服务日志的扩展配置
  logger:
    ctxKeys: ["user", "mail"] # 自动打印Context的指定变量到日志中。默认为空
    rotateExpire: "1d"
    rotateBackupExpire: "30d"
    rotateBackupLimit: 30
    rotateCheckInterval: "1h"
```

##### 2. 效果测试

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
# 开发者通过g.Log()打印的通用日志, 对server服务日志无效
logger:
  path: "logs/" # 日志文件目录, 如果为空, 表示不记录到文件
  file: "{Y-m-d}.log" # 日志文件格式。默认为"{Y-m-d}.log"
  level: "all" # DEBU < INFO < NOTI < WARN < ERRO < CRIT，也支持ALL, DEV, PROD常见部署模式配置名称。level配置项字符串不区分大小写
  stStatus: 0 # 是否打印错误堆栈(1: enabled - default; 0: disabled). 如果开启, 使用g.Log().Error 将会打印错误堆栈
  ctxKeys: ["user", "mail"] # 自动打印Context的变量到日志中。默认为空
  stdout: true # 日志是否同时输出到终端。默认true
  stdoutColorDisabled: false # 关闭终端的颜色打印。默认false
  writerColorEnable: false # 日志文件是否带上颜色。默认false，表示不带颜色
  rotateExpire: "1d" # 多长时间切分一次日志
  rotateBackupExpire: "30d" # 删除超过多长时间的切分文件, 默认为0，表示不备份，切分则删除. 如果启用按时间备份, rotateBackupLimit 必须设置为一个相对较大的数
  rotateBackupLimit: 30 # 最多保留多少个切分文件, 但rotateBackupExpire的配置优先. 默认为0，表示不备份，切分则删除. 可以不设置rotateBackupExpire
  rotateCheckInterval: "1h" # 滚动切分的时间检测间隔，一般不需要设置。默认为1小时
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

1. 使用 Makefile 编译

```sh
# For gf2-demo-api
make build

# For gf2-demo-cli
make build.cli
```

2. 查看二进制文件版本信息

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

## 📜 References

- https://goframe.org/display/gf
- https://pkg.go.dev/github.com/gogf/gf/v2
