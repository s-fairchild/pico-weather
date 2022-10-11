package components

import (
	m "machine"

	an "github.com/s-fairchild/pico-weather/anemometer"
	"github.com/s-fairchild/pico-weather/bme"
	"github.com/s-fairchild/pico-weather/bucket"
	"github.com/s-fairchild/pico-weather/oled"
)

func LoadEnabled() {

	var err error
	const flagCount = 4

	for i := 0; i<flagCount; i++ {
		switch {
		case bucket.Enabled == true && bucket.Initialized == false:
			err = bucket.Monitor()
			if err != nil {
				println(err)
			} else {
				println("Initialized and Monitoring Tipping Bucket.")
				bucket.Initialized = true
			}
		case an.Enabled == true && an.Initialized == false:
			err = an.Monitor()
			if err != nil {
				println(err)
			} else {
				an.Initialized = true
			}
		case oled.Enabled == true && oled.Initialized == false:
			err = oled.InitDisplay()
			if err != nil {
				println(err)
			} else {
				oled.Initialized = true
			}
		case bme.Enabled == true && bme.Initialized == false:
			err = bme.InitNewBme280(m.I2C1)
			if err != nil {
				println(err)
			} else {
				println("initialized bme280")
				bme.Initialized = true
			}
		}
	}
}
