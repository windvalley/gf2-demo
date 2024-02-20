#!/usr/bin/env bash
# change_project_name.sh
#
# Change current demo project name to your own project name.

set -e

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

cd "$SCRIPT_DIR" || exit 1

print_usage() {
    echo "Usage: 
    $0 <new-project>

    Example: 

    $ $0 github.com/windvalley/yournewprojectname
    "
}

[[ -z "$1" ]] && {
    print_usage

    exit 1
}

# i.e.: github.com/windvalley/gf2-demo
CURRENT_PROJECT=$(awk 'NR==1{print $2}' ../go.mod)
# i.e.: gf2-demo
CURRENT_PROJECT_NAME=${CURRENT_PROJECT##*/}

NEW_PROJECT=$1
NEW_PROJECT_NAME=${NEW_PROJECT##*/}

[[ "${NEW_PROJECT}" == "${CURRENT_PROJECT}" ]] && {
    printf "new project name must be different from current project name\n\n"
    print_usage
    exit 1
}

SED="sed"
[[ $(uname) = "Darwin" ]] && SED=gsed

# shellcheck disable=SC2038
find ../ -type f -name "*.go" -o -name "go.mod" |
    xargs $SED -i "s#${CURRENT_PROJECT}#${NEW_PROJECT}#"
# shellcheck disable=SC2038
find ../internal/cmd/ -type f -name "*.go" |
    xargs $SED -i "s#${CURRENT_PROJECT_NAME}#${NEW_PROJECT_NAME}#"

# shellcheck disable=SC2038
find ../ -type f -name "Makefile" -o -name "*.yaml" -o -name "*Dockerfile" \
    -o -name "*.sh" -o -name "*.ini" -o -name "*.service" -o -name "*.md" -o -name ".gitignore" |
    xargs $SED -i "s#${CURRENT_PROJECT_NAME}#${NEW_PROJECT_NAME}#g"

[[ "${NEW_PROJECT_NAME}" != "${CURRENT_PROJECT_NAME}" ]] && {
    mv ../cmd/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-api
    mv ../cmd/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-cli

    # shellcheck disable=SC2086
    mv ../cmd/${NEW_PROJECT_NAME}-api/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-api.go
    # shellcheck disable=SC2086
    mv ../cmd/${NEW_PROJECT_NAME}-cli/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-cli.go

    mv ../manifest/deploy/supervisor/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-api.ini
    mv ../manifest/deploy/supervisor/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-api_test.ini

    mv ../manifest/deploy/systemctl/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-api.service
    mv ../manifest/deploy/systemctl/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-api_test.service
}

exit 0
