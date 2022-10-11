//go:build !oled

package oled

import "tinygo.org/x/drivers/ssd1306"

const (
	Enabled = false
	disabledMessage = "OLED Display was disabled at build time."
) 
var Display = ssd1306.Device{}

func InitDisplay() error {

	println(disabledMessage)
	return nil
}

func WriteText(display *ssd1306.Device, lines []string) error {

	println(disabledMessage)
	return nil
}