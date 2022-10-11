package main

import (
	"fmt"
	"time"

	m "machine"

	"github.com/s-fairchild/pico-weather/bme"
	"github.com/s-fairchild/pico-weather/bucket"
	"github.com/s-fairchild/pico-weather/clock"
	"github.com/s-fairchild/pico-weather/comms"
	"github.com/s-fairchild/pico-weather/components"
	"github.com/s-fairchild/pico-weather/oled"
	t "github.com/s-fairchild/pico-weather/types"
)

func main() {

	comms.InitI2c()
	comms.InitUART()
	components.LoadEnabled()

	rt, err := clock.InitRtc(m.GP14, m.GP15)
	if err != nil {
		println(err)
	}

	r := &t.SensorReadings{time.Time{}, &t.Bme280Readings{}, &t.TippingBucket{}}

	for true {

		textLines := []string{}
		r.Created = time.Now()
		clockStr := clock.Elapsed(r.Created)
		textLines = append(textLines, clockStr)

		var err error
		var tempUnit = "F"
		r.Bme280, err = bme.GetAllReadings(tempUnit)
		if err != nil {
			println(err)
		} else {
			tempStr := fmt.Sprintf("Temperature: %2.2f°%v\n", r.Bme280.Temp, tempUnit)
			println(tempStr)
			textLines = append(textLines, tempStr)

			humidityStr := fmt.Sprintf("Humidity: %2.2f%%\n", r.Bme280.Humidity)
			println(humidityStr)
			textLines = append(textLines, humidityStr)

			pressureStr := fmt.Sprintf("Pressure: %.2f mbar\n", r.Bme280.Pressure)
			println(pressureStr)
			textLines = append(textLines, pressureStr)
		}

		r.Rain.Amount = bucket.GetRain()
		rainStr := fmt.Sprintf("Rainfall: %.f\n", r.Rain.Amount)
		println(rainStr)
		if r.Rain.Amount != 0 {
			textLines = append(textLines, rainStr)
		}

		timeStr, err := clock.FormatTime(&rt)
		if err != nil {
			println(err)
		}
		fmt.Printf("RTC: %v\n", timeStr)

		textLines = append(textLines, timeStr)

		err = oled.WriteText(textLines)
		if err != nil {
			println(err)
		}

		comms.TxSerialData(r)

		println(clock.ExecutionTime(r.Created))
		time.Sleep(clock.TxInterval)
	}
}
