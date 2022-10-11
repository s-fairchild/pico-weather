//go:build !bucket

package bucket

const Enabled bool = false

func GetRain() float64 {

	return 0.0
}

func Monitor() error {

	println("Bucket was disabled at build time.")
	return nil
}

func ResetRain() {

	Tips = 0
}