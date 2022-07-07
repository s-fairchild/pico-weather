#!/bin/bash

# repoCheck Finds the openocd repo path if this script is ran with sudo
homeCheck() {

    if [[ $UID -eq 0 ]]; then
        echo "/home/$SUDO_USER"
    else
        echo $HOME
    fi
}

home="$(homeCheck)"
buildPath="${home}/go/src/github.com/pico-weather/build/release"

launchOpenocd() {

    if [[ -v $1 ]] && [[ -z $1 ]]; then
        binaryFile=$1
    else
        binaryFile="${buildPath}"
    fi
    sudo openocd -s ~/src/pico/openocd/tcl/ -f interface/picoprobe.cfg -f target/rp2040.cfg -c "program ${binaryFile} verify reset exit"
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


repoPath="${home}/src/pico/openocd"
if [[ ! -d $repoPath ]]; then
    scripts/build_install_openocd.sh
fi

if launchOpenocd "$1"; then
    launchMinicom
fi