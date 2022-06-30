package bme

import (
	"fmt"
	m "machine"

	"tinygo.org/x/drivers/bme280"
)

var (
	B280 = bme280.New(m.I2C1)
)

// InitNewBme280 configures and returns a new bme280 object
// The I2C bus MUST be configured already
func InitNewBme280() {

	B280.Address = uint16(0x76)
	err := B280.Configure()
	if err != nil {
		println("Failed to configure bme280: ", err)
	}

	if B280.Connected() {
		println("initialized bme280")
	} else {
		println("Failed to connect to bme280")
	}
}

// tempCtoF Converts and returns Celcius in milidegrees to Fahrenheit degrees
func TempCtoF() (float32, error) {

	tempC, err := B280.ReadTemperature()
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read temperature, %v", err)
	}

	degree := float32(tempC / 1000)
	tempF := (degree * 1.8) + 32.0

	return tempF, nil
}

func HumidityPercent() (float32, error) {

	h, err := B280.ReadHumidity()
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read humidity, %v", err)
	}

	return float32(h) / 100.0, nil
}

func PressureInchHg() (float32, error) {

	p, err := B280.ReadPressure()
	if err != nil {
		return 0.0, fmt.Errorf("Failed to read pressure, %v", err)
	}

	return mPAtoinHG(p) / 100000.0, nil
}

func mPAtoinHG(mPA int32) float32 {

	PA := mPA / 1000
	inHG := float32(PA) * 3386.3886666667

	return inHG
}