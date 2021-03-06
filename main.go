package main

import (
	"fmt"
	"time"

	m "machine"

	an "github.com/pico-weather/anemometer"
	"github.com/pico-weather/bme"
	"github.com/pico-weather/bucket"
	t "github.com/pico-weather/types"
)

func main() {

	initUART()
	initI2c()
	bme.InitNewBme280(m.I2C1)
	err := bucket.Monitor()
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		println("Initialized and Monitoring Tipping Bucket.")
	}
	err = an.Monitor()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	var r = &t.SensorReadings{
		Created:  time.Time{},
		Bme280:   &t.Bme280Readings{},
		Rainfall: &t.TippingBucket{},
	}

	for true {

		r.Bme280.TempF, err = bme.ReadTempF()
		if err != nil {
			println(err)
		} else {
			fmt.Printf("tempF: %2.2f°F\n", r.Bme280.TempF)
		}

		r.Bme280.Humidity, err = bme.HumidityPercent()
		if err != nil {
			println(err)
		} else {
			fmt.Printf("humidity: %2.2f%%\n", r.Bme280.Humidity)
		}

		r.Bme280.Pressure, err = bme.PressureMilliBar()
		if err != nil {
			println(err)
		} else {
			fmt.Printf("pressure: %f mbar\n", r.Bme280.Pressure)
		}

		r.Rainfall.Inches = bucket.GetRain()
		println()
		if err != nil {
			fmt.Printf("%v\n", err)
		}

		// Time will always start at the beginning of Unix time until a RTC is added
		// I have a Sunfounder RTC PCF8563 I plan on testing
		r.Created = time.Now()
		txJSON(r)

		time.Sleep(t.TxInterval)
	}
}

// txJSON marshals and writes to UART1
func txJSON(r *t.SensorReadings) {

	txData, err := r.MarshalJSON()
	if err != nil {
		fmt.Printf("Error marshling JSON, %v\n", err)
	} else if len(txData) == 0 {
		fmt.Printf("txData has a length of %c, not transmitting.\n", len(txData))
	} else {
		fmt.Printf("Transmitting JSON over UART, length of: %v\n", len(txData))
		m.UART1.Write(txData)
	}
}

// initUART Configures the pico's uart0 as the stdout for text, and uart1 is configured for transmitting data
// to the receiving system
func initUART() {

	// Text console output
	err := m.UART0.Configure(m.UARTConfig{})
	if err != nil {
		fmt.Printf("Error configuring Serial UART0, %c\n", err)
	} else {
		println("initialized Serial UART0\n")
	}

	// Data output
	err = m.UART1.Configure(m.UARTConfig{
		TX: m.UART1_TX_PIN,
		RX: m.UART1_RX_PIN,
	})
	if err != nil {
		fmt.Printf("Error Configuring Serial UART1, %v\n", err)
	} else {
		println("initialized Serial UART1")
	}
	//m.UART1.SetFormat(1, 8, m.ParityNone)
}

func initI2c() {

	// I2C1 is used here to allow use of default UART0 pins
	err := m.I2C1.Configure(m.I2CConfig{})
	if err != nil {
		fmt.Printf("Failed to Configure I2C1, %v", err)
	} else {
		println("initialized I2C1")
	}
}
