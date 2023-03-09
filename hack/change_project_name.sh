#!/usr/bin/env bash
# change_project_name.sh
#
# Change current demo project name to your own project name.
#
# e.g.:
#   ./change_project_name.sh your-project-name

set -e

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

cd "$SCRIPT_DIR" || exit 1

[[ -z "$1" ]] && {
    echo "Usage: 
    $0 <new-project-name>"

    exit 1
}

CURRENT_PROJECT_NAME=$(awk 'NR==1{print $2}' ../go.mod)
NEW_PROJECT_NAME=$1
SED="sed"

[[ $(uname) = "Darwin" ]] && SED=gsed

# shellcheck disable=SC2038
find ../ -type f -name "*.go" |
    xargs $SED -i "s#${CURRENT_PROJECT_NAME}#${NEW_PROJECT_NAME}#"

# shellcheck disable=SC2038
find ../ -type f -name "go.mod" -o -name "Makefile" \
    -o -name "*.md" -o -name "*.yaml" -o -name "*.sh" -o -name "*Dockerfile" |
    xargs $SED -i "s#${CURRENT_PROJECT_NAME}#${NEW_PROJECT_NAME}#"

mv ../cmd/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-api
mv ../cmd/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-cli

# shellcheck disable=SC2086
mv ../cmd/${NEW_PROJECT_NAME}-api/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-api.go
# shellcheck disable=SC2086
mv ../cmd/${NEW_PROJECT_NAME}-cli/{"${CURRENT_PROJECT_NAME}","${NEW_PROJECT_NAME}"}-cli.go

exit 0
