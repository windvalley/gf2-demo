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

- 通过命令行指定配置文件

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

### 响应业务码

### 服务日志

### 链路跟踪

### 版本管理

## 📜 References

- https://goframe.org/display/gf
- https://pkg.go.dev/github.com/gogf/gf/v2
