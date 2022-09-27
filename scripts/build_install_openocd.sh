#!/bin/bash
# Build and install openocd for Fedora
# Tested on Fedora 36

main() {

    read -n1 -p "Install openocd built with picoprobe support? [Y/N]" install
    echo ""

    if [[ ! $install =~ ^[Yy]$ ]]; then
        abort
    fi

    installPkgs
    cloneRepo
    buildAndInstall

    echo -e "\nNOTE: Your user must be a member of the dialout group to access serial devices.\n"
    echo -e "Run this command to add user to dialout group: \n\tsudo usermod -aG dialout $USER"
}

buildAndInstall() {

    ./bootstrap || abort
    if ./configure --enable-ftdi --enable-sysfsgpio --enable-bcm2835gpio; then
        make -j"$(nproc)" || abort
        openocdBin="$(find "$HOME/src/pico/openocd/" -name openocd -type f)"

        read -n1 -p "Install openocd systemwide select S, or locally select L? [S/L]: " feedback
        echo ""
        if [[ $feedback =~ ^[Ss]$ ]]; then
            sudo make install;
        elif [[ $feedback =~ ^[Ll]$ ]]; then
            echo "Installing openocd in $HOME/bin"
            cp "$openocdBin" "$HOME/bin/"
            verifyInstall "$openocdBin"
        else
            echo "Unrecognized option $feedback."
            echo "openocd binary is located at $openocdBin"
        fi
    fi
}

# verifyInstall accepts a file to verify it was copied successfully
verifyInstall() {

    if [[ $1 ]]; then
        echo "Successfully install ${1}"
    else
        echo "Something went wrong, ${1} not found"
    fi
}

cloneRepo() {

    repoRoot="$HOME/src/pico"
    if [[ ! -d $HOME/src/pico ]]; then
        mkdir "$repoRoot"
    fi
    pushd $HOME/src/pico || abort

    if [[ -d openocd ]]; then
        pushd openocd || abort
        git pull origin rp2040 --recurse-submodules --depth=1 || abort
    else 
        pushd openocd || abort
        git clone https://github.com/raspberrypi/openocd.git --recursive --branch rp2040 --depth=1 || abort
    fi
}

installPkgs() {

    pkgs=( libftdi-devel
        automake
        autoconf
        texinfo
        libtool
        libusb1-devel
        git
        make
    )

    echo "Installing packages ${pkgs[*]}"
    sudo dnf in "${pkgs[@]}" -y
    sudo dnf groupinstall "Development Tools" -y
}

installLocally() {

    if [[ ! -d "$HOME/bin" ]]; then
        mkdir "$HOME/bin"
    fi
}

abort () {
    echo "Aborting..."
    exit 1
}

main