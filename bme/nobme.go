//go:build !bme

package bme

import (
	t "github.com/s-fairchild/pico-weather/types"
	m "machine"
)

const (
	Enabled = true
	disabledMessage = "BME Disabled at build time."
)
var Initialized bool

func GetAllReadings(unit string) (*t.Bme280Readings, error) {

	println(disabledMessage)
	r := t.Bme280Readings{}
	return &r, nil
}

func InitNewBme280(i2cDev *m.I2C) {

	println(disabledMessage)
}