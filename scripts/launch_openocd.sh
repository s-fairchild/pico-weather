#!/bin/bash

launch() {
    openocd -s ~/src/pico/openocd/tcl/ -f interface/picoprobe.cfg -f target/rp2040.cfg -c "program build/release verify reset exit"
}

if [[ ! -d $HOME/src/pico/openocd/tcl/ ]]; then
    scripts/build_install_openocd.sh
fi
launch