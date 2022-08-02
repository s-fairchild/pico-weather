#!/bin/bash
# Tinygo tests have to be ran from the relative package directory

set -o nounset \
    -o pipefail

# runTests accepts any number of Go package directories, enters them and runs `go test -v`
runTests() {

    for package in "${@}"; do
        if [[ -d $package ]]; then
            cd "${package}" || exit 1
            go test -v
        else
            echo "Package ${package} doesn't appear to exist, not running unit tests for ${package}"
        fi
    done
}

IFS=' ' read -r -a arr <<< "${@}"
runTests "${arr[@]}"