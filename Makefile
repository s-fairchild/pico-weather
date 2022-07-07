ONESHELL:
SHELL = /bin/bash

easyjson:
	if [[ types/types.go -nt types/types_easyjson.go ]]; then \
		easyjson -omit_empty -noformat types/; \
	fi

gochecks: easyjson
	go mod tidy
	go fmt .

release: gochecks
	build=$$(scripts/go_change_check.sh build/release); \
	if [ $$build == "true" ]; then \
		tinygo build -target=pico -o build/release; \
	fi

flash: release
	scripts/launch_openocd.sh

terminal:
	if [ -a "/dev/ttyACM0" ]; then \
		minicom -D /dev/ttyACM0 -b 115200; \
	else \
		echo "No device /dev/ttyACM0 found"; \
	fi
