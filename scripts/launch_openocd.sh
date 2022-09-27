#!/bin/bash

set -o errexit

main() {

    realHome="$(homeCheck)"
    userInput="$1"
    binaryFile="${userInput:="build/release"}"
    repoPath="${realHome}/src/pico/openocd"
    if [[ ! -d $repoPath ]]; then
        scripts/build_install_openocd.sh
    fi

    if launchOpenocd "$binaryFile"; then
        launchMinicom "$binaryFile"
    fi
}

# homeCheck Finds the openocd repo path if this script is ran with sudo
homeCheck() {

    if [[ $UID -eq 0 ]]; then
        echo "/home/$SUDO_USER"
    else
        echo "$HOME"
    fi
}

launchOpenocd() {

    sudo openocd -s "${realHome}/src/pico/openocd/tcl/" -f interface/picoprobe.cfg -f target/rp2040.cfg -c "program ${1} verify reset exit"
}

launchMinicom() {

    local termDev
    termDev="/dev/ttyACM0"
	if [[ -a $termDev ]]; then
        minicom -D $termDev -b 115200
    else
        echo "No device $termDev found"
    fi
}

main "$1"