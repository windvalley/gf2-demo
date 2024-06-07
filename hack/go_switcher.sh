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

SHELL_RC_FILE=~/.zshrc
GO_INSTALL_PATH="/usr/local/opt" # /usr/local/opt or /opt/homebrew/opt
GO_PATH_LINK="$GO_INSTALL_PATH/go"
GO_BIN_PATH="$GO_PATH_LINK/bin"

show_go_versions() {
    echo "Current Go version: 
$(go version)

Go versions installed on this system:"

    go_list=$(ls -d $GO_INSTALL_PATH/go@*)
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

    if [[ ! -d "$GO_INSTALL_PATH/go@$version" ]]; then
        echo -e "Error: Go version not found -- $version\n"
        show_go_versions
        return 1
    fi

    # Switch to the specified Go version.
    unlink "$GO_PATH_LINK" &>/dev/null
    ln -sf "$GO_INSTALL_PATH/go@$version" "$GO_PATH_LINK"

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
    if ! grep -q "PATH=\"$GO_BIN_PATH:" "$SHELL_RC_FILE"; then
        echo "export PATH=\"$GO_BIN_PATH:\$PATH\"" >>"$SHELL_RC_FILE"
    fi

    echo "Warning: please reload your shell rc file, i.e. 'source $SHELL_RC_FILE'"

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
