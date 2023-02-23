#!/usr/bin/env bash
# deploy.sh
#

set -e

env=$1

[[ "$env" != "test" ]] && [[ "$env" != "prod" ]] || [[ "$#" != 1 ]] && {
    echo "Usage: $0 <test|prod>"
    exit 1
}

SED=sed
[[ $(uname) = "Darwin" ]] && SED=gsed

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

project_name="gf2-demo"
app_name="${project_name}-api"
systemd_service="${app_name}.service"

deploy_user="vagrant"
deploy_server="gf2-demo.sre.im" # 请提前配置发布机到目标服务器之间的ssh key信任
deploy_dir=/app/$project_name

[[ "$env" = "test" ]] && {
    systemd_service="${app_name}_test.service"
    deploy_dir=${deploy_dir}_test
}

cd "$script_dir" || exit 1

directory=$(awk -F= '/WorkingDirectory/{print $2}' $systemd_service)
$SED -i "s#$directory#$deploy_dir#" $systemd_service

# ******** 部署systemd_service

# shellcheck disable=SC2029
ssh $deploy_user@$deploy_server "
    sudo mkdir -p $deploy_dir;
    sudo chown -R $deploy_user. $deploy_dir;
    rm -f $deploy_dir/$app_name
"

scp $systemd_service $deploy_user@$deploy_server:$deploy_dir

# shellcheck disable=SC2029
ssh $deploy_user@$deploy_server "
    sudo \cp $deploy_dir/$systemd_service /usr/lib/systemd/system/;
    sudo systemctl daemon-reload;
    sudo systemctl enable $systemd_service;
"

# ******** 部署项目

cd ../../../ || exit 1

make build

scp bin/linux_amd64/${app_name} $deploy_user@$deploy_server:$deploy_dir

# shellcheck disable=SC2029
ssh $deploy_user@$deploy_server "sudo systemctl restart ${systemd_service}"

exit 0
