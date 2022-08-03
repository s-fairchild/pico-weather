ONESHELL:
SHELL = /bin/bash

easyjson:
	if [[ types/types.go -nt types/types_easyjson.go ]]; then \
		easyjson -omit_empty -noformat types/; \
	fi

gochecks: easyjson
	go mod tidy
	go fmt .
	# go vet fails with tiny go currently

gotests: gochecks
	# testPkgs=$$(find -name *_test.go -printf "%h\n"); \
	pkgsWithTests=("$$(find ./ -name '*_test.go' -printf "%h\n" | sort -ub)"); \
	scripts/runUnitTests.sh $$pkgsWithTests

release: gotests
	build=$$(scripts/go_change_check.sh build/release); \
	if [ $$build == "true" ]; then \
		tags="-tags bsizeinches"; \
		tinygo build -target=pico -o build/release $$tags; \
	fi

flash: release
	scripts/launch_openocd.sh

terminal:
	if [ -a "/dev/ttyACM0" ]; then \
		minicom -D /dev/ttyACM0 -b 115200; \
	else \
		echo "No device /dev/ttyACM0 found"; \
	fi
