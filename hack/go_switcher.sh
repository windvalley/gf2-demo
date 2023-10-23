#!/usr/bin/env bash
# go_switcher.sh
#
# Switch between Go versions.
#
# NOTE: Only tested on macOS.
#
# How to install Go: brew install go@1.20

show_go_versions() {
    echo "Current Go version: 
$(go version)

Go versions installed on this system:"

    # shellcheck disable=SC2125
    go_list=/usr/local/opt/go@*
    for version in $go_list; do
        echo "* $(basename "$version")"
    done
}

switch_go_version() {
    local version="$1"
    if echo "$version" | grep -q 'go@'; then
        version=$(echo "$version" | cut -d'@' -f2)
    fi

    if [[ -z "$version" ]]; then
        echo "Error: no Go version specified"
        return 1
    fi

    if [[ ! -d "/usr/local/opt/go@$version" ]]; then
        echo "Error: Go version $version not found"
        return 1
    fi

    echo "$PATH" | grep -q /usr/local/opt/go/bin || {
        export PATH=/usr/local/opt/go/bin:$PATH
        echo "export PATH=/usr/local/opt/go/bin:\$PATH" >>~/.zshrc
    }

    unlink /usr/local/opt/go
    ln -sf "/usr/local/opt/go@$version" /usr/local/opt/go

    echo "Switched to Go version: $version"
}

show_help() {
    echo "Usage:

    -l            Show all installed Go versions
    -s <version>  Switch to the specified Go version
    -h            Show this help message

Example:

    $ $0 -l

    $ $0 -s 1.20"
}

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

show_help

exit 0
