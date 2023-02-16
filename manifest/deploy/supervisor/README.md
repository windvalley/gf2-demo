# 部署流程

## 安装 supervisor

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

## 在发布机执行部署脚本

```sh
# 部署测试环境
$ manifest/deploy/supervisor/deploy.sh test

# 部署生产环境
$ manifest/deploy/supervisor/deploy.sh prod
```
