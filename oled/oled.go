//go:build oled

package oled

import (
	m "machine"

	"fmt"
	"image/color"

	"tinygo.org/x/drivers/ssd1306"

	"tinygo.org/x/tinyfont"
)

var (
	colors = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,	
	}
	Initialized bool
	display = ssd1306.Device{}
) 

const (
	sdaPin = m.GP20
	sclPin = m.GP21
	Enabled = true
)

func WriteText(lines []string) error {

	display.ClearDisplay()
	// y coordinate
	// Add a new line every 10 points down
	for i, y := 0, 10; i < len(lines); i, y = i+1, y+10 {
		tinyfont.WriteLine(&display, &tinyfont.Org01, 0, int16(y), lines[i], colors)
		if y > 64 {
			return fmt.Errorf("y pixel cannot be greater than 64, y is: %v", y)
		}
	}

	display.Display()
	return nil
}

func InitDisplay() error {

	i2c, err := configureI2c(sdaPin, sclPin)
	if err != nil {
		return err
	}

	display = ssd1306.NewI2C(i2c)
	c := ssd1306.Config{
		Address: 0x3C,
		Width: 128,
		Height: 64,
	}
	display.Configure(c)

	display.ClearDisplay()
	
	return err
}

func configureI2c(sda m.Pin, scl m.Pin) (*m.I2C, error) {

	c := m.I2CConfig{
		Frequency: m.TWI_FREQ_400KHZ,
		SDA: sda,
		SCL: scl,
	}
	err := m.I2C0.Configure(c)
	if err != nil {
		return m.I2C0, fmt.Errorf("error configuring I2C0 on pins SDA: %v, SCL: %v\n%v\n", sda, scl, err)
	}
	return m.I2C0, nil
}