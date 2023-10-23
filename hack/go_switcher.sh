#!/usr/bin/env bash
# go_switcher.sh
#
# Go Multi-version Switching Management Tool.
#
# NOTE: 1. Only tested on macOS; 2. Use Oh My Zsh.
#
# How to install Go:
#   brew install go@1.17
#   brew install go@1.21
#
# How to install this script:
#   cp ./go_switcher.sh /usr/local/bin/gov
#
# Usage example:
#   $ gov -l
#   $ gov -s 1.21

GO_BIN_PATH="/usr/local/opt/go/bin"

show_go_versions() {
    echo "Current Go version: 
$(go version)

Go versions installed on this system:"

    # shellcheck disable=SC2125
    go_list=$(ls -d /usr/local/opt/go@*)
    for version in $go_list; do
        version=$(echo "$version" | grep -Eo "[0-9]+\.[0-9]+")
        echo "* $version"
    done
}

switch_go_version() {
    local version="$1"
    if ! echo "$version" | grep -qE '^[0-9]+\.[0-9]+$'; then
        echo -e "Error: invalid Go version specified -- $version\n"
        show_go_versions
        return 1
    fi

    if [[ ! -d "/usr/local/opt/go@$version" ]]; then
        echo -e "Error: Go version not found -- $version\n"
        show_go_versions
        return 1
    fi

    # Switch to the specified Go version.
    unlink /usr/local/opt/go
    ln -sf "/usr/local/opt/go@$version" /usr/local/opt/go

    echo "Switched to Go version: $version"
}

show_help() {
    echo "Go Multi-version Switching Management Tool.

Usage:

    -l            Show all installed Go versions
    -s <version>  Switch to the specified Go version
    -h            Show this help message

Example:

    $ $0 -l

    $ $0 -s 1.20"
}

# Make sure $GO_BIN_PATH is in $PATH.
if ! echo "$PATH" | grep -q "$GO_BIN_PATH"; then
    if ! grep -q "PATH=\"$GO_BIN_PATH:" ~/.zshrc; then
        echo "export PATH=\"$GO_BIN_PATH:\$PATH\"" >>~/.zshrc
    fi

    echo "Warning: please execute 'omz reload' or 'exec zsh' first"
    exit 1
fi

# Parse command line options.
while getopts "ls:h" option; do
    case $option in
    l)
        show_go_versions
        exit
        ;;
    s)
        switch_go_version "$OPTARG"
        exit
        ;;
    h)
        show_help
        exit
        ;;
    *) ;;

    esac
done

# Show help if no options were specified.
show_help

exit 0
