# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
