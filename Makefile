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
	openocd -s ~/src/pico/openocd/tcl/ -f interface/picoprobe.cfg -f target/rp2040.cfg -c "program build/release verify reset exit"

serialterm: debugbuild
	minicom -D /dev/ttyACM0 -b 115200
