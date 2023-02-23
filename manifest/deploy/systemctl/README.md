# systemctl

## 根据实际情况修改 deploy.sh 部署相关变量

```sh
project_name="gf2-demo"
app_name="${project_name}-api"
systemd_service="${app_name}.service"

deploy_user="vagrant"
deploy_server="gf2-demo.sre.im" # 请提前配置发布机到目标服务器之间的ssh key信任
deploy_dir=/app/$project_name
```

## 在发布机执行部署脚本 deploy.sh

```sh
# 部署测试环境
$ manifest/deploy/systemctl/deploy.sh test

# 部署生产环境
$ manifest/deploy/systemctl/deploy.sh prod
```

## service 文件配置

```ini
[Unit]
Description=gf2-demo-api production environment
Documentation=
After=network.target
Wants=

[Service]
Type=simple
# 程序启动用到的环境变量.
Environment=GF_GERROR_BRIEF=true GF_GCFG_FILE=config.prod.yaml
# 程序所在目录.
WorkingDirectory=/app/gf2-demo
# 使用 sh -c "" 这种方式启动, 避免将stdout和stderr输出到/var/log/message中,
# 统一输出到指定的文件中.
ExecStart=/bin/sh -c '/app/gf2-demo/gf2-demo-api >> gf2-demo-api.log 2>&1'
# 发送SIGTERM信号给所有控制的进程.
KillMode=ctrl-group
# 只能通过 systemctl stop 来关闭服务, 直接kill进程服务还会自动重启.
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
```

## systemctl 常用操作命令

```sh
# gf2-demo-api.service 配置有变动的时候, 需要重新加载使生效
$ sudo systemctl daemon-reload

# 启动
$ sudo systemctl start gf2-demo-api

# 关闭: 发送 SIGTERM 信号给主(sh)和子进程(gf2-demo-api),
# gf2-demo-api程序可通过捕捉SIGTERM信号来实现优雅关闭.
$ sudo systemctl stop gf2-demo-api

# 重启: 先关闭(SIGTERM), 再启动
$ sudo systemctl restart gf2-demo-api
```
