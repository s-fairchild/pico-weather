#!/bin/bash
# Build and install openocd for Fedora
# Tested on Fedora 36

pkgs=(
    libftdi-devel
    automake
    autoconf
    texinfo
    libtool
    libusb1-devel
    git
    make
)

abort () {
    exit 1
}

echo "Installing packages ${pkgs[*]}"
sudo dnf in "${pkgs[@]}" -y
sudo dnf groupinstall "Development Tools" -y

if [[ -d $HOME/src ]]; then
    cd $HOME/src || abort
else
    cd /tmp || abort
fi

if [[ -d openocd ]]; then
    cd openocd || abort
    git pull origin rp2040 --recurse-submodules --depth=1 || abort
else 
    git clone https://github.com/raspberrypi/openocd.git --recursive --branch rp2040 --depth=1 || abort
fi

./bootstrap || abort
if ./configure --enable-ftdi --enable-sysfsgpio --enable-bcm2835gpio; then
    make -j"$(nproc)" || abort
    read -n1 -p "Install openocd systemwide?\n y/n: " feedback
    if [[ $feedback == 'y' ]]; then
        sudo make install;
    else
        echo "Not installing openocd. The binary is located at $(pwd)/src/openocd"
    fi
fi

echo "NOTE: required for non root serial device access."
echo -e "Run this command to add user to dialout group: \n\tsudo usermod -aG dialout $USER"