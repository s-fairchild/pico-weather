#!/bin/bash
# Tinygo tests have to be ran from the relative package directory

set -o nounset \
    -o pipefail

# runTests accepts any number of Go package directories, enters them and runs `go test -v` after replacing _test.go with regex char *
runTests() {

    local -a testDirs
    testDirs=( "${@}" )

    for p in "${testDirs[@]}"; do
        if [[ -d $p ]]; then
            pushd "$p" || exit 1
            testFiles="$(find ./ -name "*_test.go" | sed s/'_test.go'/'*'/g)"
            go test -v $testFiles
            popd &> /dev/null || exit 1
        fi
    done
}

runTests "${@}"
