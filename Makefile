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
	pkgsWithTests=("$$(find ./ -name '*_test.go' -printf "%h\n" | sort -ub)"); \
	scripts/runUnitTests.sh $$pkgsWithTests

# TODO add -no-debug and make a debug target
release: gotests
	# All tags: bsizeinches bucket oled bme anemometer
	if [[ ! -d build ]]; then \
		mkdir build ;\
	fi ;\
	build=$$(scripts/go_change_check.sh build/release); \
	if [ $$build == "true" ]; then \
		tags="bsizeinches bucket oled bme"; \
		tinygo build -target=pico -serial=uart -o build/release -tags "$$tags"; \
	fi

flash: release
	scripts/launch_openocd.sh build/release

terminal:
	if [ -a "/dev/ttyACM0" ]; then \
		minicom -D /dev/ttyACM0 -b 115200; \
	else \
		echo "No device /dev/ttyACM0 found"; \
	fi
