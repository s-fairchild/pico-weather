#!/bin/bash

set -o pipefail -o nounset

goFilesChangeCheck() {

    binaryFile="$1"

    readarray -d '' goFiles < <(find . -name "*.go" -print0)

    local build
    build="false"
    for f in "${goFiles[@]}"; do
        if [ $f -nt $binaryFile ]; then
            build="true"
            break 2
        fi
    done
    echo $build
}

goFilesChangeCheck "$1"