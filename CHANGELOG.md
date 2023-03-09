# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## v0.7.5

### Changed

- gf 框架版本升级到 `v2.3.3`
- 美化 README.md 文档格式

### Fixed

- 修复 `hack/change_project_name.sh` 脚本变更项目名称后 `make build` 失败的问题

## v0.7.4

### Changed

- 优化配置文件, 使服务启动过程中产生的日志和开发者打印的日志(`g.Log()`)使用同一个日志文件.
  问题描述请见: [gogf#2462](https://github.com/gogf/gf/issues/2462)

## v0.7.3

### Changed

- 优化 systemctl 部署服务方式
  - 优化服务文件 `manifest/deploy/systemctl/gf2-demo-api*.service`
  - 优化一键部署脚本 `manifest/deploy/systemctl/deploy.sh`
  - 增加说明文档 `manifest/deploy/systemctl/README.md`
- 优化 `supervisor` 部署服务方式
  - 优化说明文档 `manifest/deploy/supervisor/README.md`
- 优化 README.md 文档
  - 增加 `systemctl/supervisor/docker` 如何优雅关闭服务的说明, 不过需要 `GoFrameV2` 框架提供支持通过捕获 `SIGTERM` 信号实现优雅关闭的功能, 目前测试此功能还不可用(仅优雅重启可用).

## v0.7.2

### Changed

- 进一步优化 Dockerfile, 提高编译速度
- 优化 Makefile
  - 增加参数 OS 和 ARCH, 用于灵活 `make build`
  - 优化 `make image`, 制作 docker 镜像
  - 优化 `make image.push`, 制作 docker 镜像并推送到 docker 仓库
  - 增加 TAG 参数, 用于灵活 `make image`
  - 优化 make 帮助信息: `make` / `make help`
- 优化 README.md 文档

## v0.7.1

### Changed

- 优化 Dockerfile, 暴露编译过程, 使用两阶段构建, 依赖包可缓存, 编译速度快
- 优化 Makefile 文件
- GoFrame V2.3.1 -> V2.3.2

## v0.7.0

### Added

- 增加通过 `systemctl` 一键部署方式
- 增加通过 `supervisor` 一键部署方式
- 增加 `Dockerfile`

### Changed

- 进一步完善 README.md 文档

### Removed

- 删除了 `manifest/docker/`

## v0.6.1

### Added

- 增加获取资源列表接口的代码示例

### Changed

- 响应 json 增加`traceid`字段, 将`msg`字段修改为`message`:

```json
{
  "code": "string",
  "message": "string",
  "traceid": "string",
  "data": null
}
```

- 优化 `401` 和 `403` 的业务错误码描述
- 完善 README.md 文档, 增加详细的开发流程

## v0.6.0

### Added

- 增加关系数据库配置和 sql 日志配置
- 增加将各层串起来的简单 CRUD 示例接口供参考

### Changed

- 优化响应 json 的字段顺序: code -> msg -> data

## v0.5.0

### Added

- 优化工程目录结构, 使支持多个可执行命令
- 规范业务错误码, 中间件统一拦截响应, 规范响应格式
- 完善 HTTP 服务访问日志、HTTP 服务错误日志、开发者打印的日志、其他可执行命令的日志配置
- 多环境管理: 开发环境、测试环境、生产环境
- 编译的二进制文件可打印当前应用的版本信息
- 链路跟踪中间件, 默认使用客户端按规范传递的`X-Request-Id`
- 通过 Makefile 管理项目: `make run`, `make run.cli`, `make build`, `make build.cli` 等

## v0.0.0

### Added

### Changed

### Removed

### Fixed

### Security
