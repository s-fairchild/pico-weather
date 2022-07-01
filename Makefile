ONESHELL:
SHELL = /bin/bash

easyjson:
	easyjson -omit_empty -noformat types/

gochecks: easyjson
	go mod tidy
	go fmt .

debugbuild: gochecks
	tinygo build -target=pico -o build/debug -opt=1 -serial=uart

release: gochecks
	tinygo build -target=pico -o build/release -serial=uart

run: release
	scripts/launch_openocd.sh

terminal:
	if [ -a "/dev/ttyACM0" ]; then minicom -D /dev/ttyACM0 -b 115200; else echo "No device /dev/ttyACM0 found"; fi
