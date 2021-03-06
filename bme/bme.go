package bme

import (
	"fmt"
	m "machine"

	"github.com/pico-weather/convert"

	"tinygo.org/x/drivers/bme280"
)

var (
	// B280 = bme280.New(m.I2C1)
	B280 = bme280.Device{}
)

// InitNewBme280 configures and returns a new bme280 object
// The I2C bus must be configured already
func InitNewBme280(i2cDev *m.I2C) {

	B280 = bme280.New(i2cDev)
	B280.Address = uint16(0x76)
	B280.Configure()

	if B280.Connected() {
		println("initialized bme280")
	} else {
		println("Failed to connect to bme280")
	}
}

func ReadTempF() (float32, error) {

	tempC, err := B280.ReadTemperature()
	// TODO convert to debug log
	// fmt.Printf("Temperature C as int32: %v", tempC)
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read temperature, %v\n", err)
	}
	return convert.Celsius2Fahrenheit(tempC), nil
}

func HumidityPercent() (float32, error) {

	h, err := B280.ReadHumidity()
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read humidity, %v\n", err)
	}

	return float32(h) / 100.0, nil
}

func PressureMilliBar() (float64, error) {

	p, err := B280.ReadPressure()
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read pressure, %v\n", err)
	}

	return convert.MilliPA2MilliBar(p), nil
}
