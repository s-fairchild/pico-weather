//go:build !anemometer

package anemometer

const Enabled = false
var Initialized bool

func Monitor() error {

	println("Anemometer was disabled at build time.")
	return nil
}