//go:build bme

package bme

import (
	"fmt"
	m "machine"

	"github.com/s-fairchild/pico-weather/convert"
	t "github.com/s-fairchild/pico-weather/types"

	"tinygo.org/x/drivers/bme280"
)

var BmeDev = bme280.Device{}
var Initialized bool
const Enabled = true

func GetAllReadings(unit string) (*t.Bme280Readings, error) {

	r := &t.Bme280Readings{}
	var err error
	r.Pressure, err = PressureMilliBar()
	if err != nil {
		return r, err
	}

	r.Temp, err = ReadTempF(unit)
	if err != nil {
		return r, err
	} 

	r.Humidity, err = HumidityPercent()
	if err != nil {
		return r, err
	}

	return r, nil
}

// InitNewBme280 configures and returns a new bme280 object
// The I2C bus must be configured already
func InitNewBme280(i2cDev *m.I2C) error {

	BmeDev = bme280.New(i2cDev)
	BmeDev.Address = uint16(0x76)
	BmeDev.Configure()

	if BmeDev.Connected() {
		return nil
	} else {
		return fmt.Errorf("Failed to connect to bme280")
	}
}

func ReadTempF(unit string) (float32, error) {

	temp, err := BmeDev.ReadTemperature()
	temp = temp / 1000
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read temperature, %v\n", err)
	}
	if unit == "f" || unit == "F" {
		return convert.Celsius2Fahrenheit(temp), nil
	} else if unit == "c" || unit == "C" {
		return float32(temp), nil
	}
	return 0.0, fmt.Errorf("Something went wrong, Unable to calculate temperature")
}

func HumidityPercent() (float32, error) {

	h, err := BmeDev.ReadHumidity()
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read humidity, %v\n", err)
	}

	return float32(h) / 100.0, nil
}

func PressureMilliBar() (float64, error) {

	p, err := BmeDev.ReadPressure()
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read pressure, %v\n", err)
	}

	return convert.MilliPA2MilliBar(p), nil
}
