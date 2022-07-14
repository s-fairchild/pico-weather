package bucket

import (
	m "machine"
	"fmt"
	c "github.com/pico-weather/convert"
)

var Tips uint

func GetRain() float64 {

	return c.CalculateRain(Tips, bucketSize)
}

// Monitor watches a tipping bucket button and populates the Tips field for each press
func Monitor() error {

	pin := m.GP18
	pin.Configure(m.PinConfig{m.PinInputPullup})

	// PinFalling is required because the tipping bucket has a slow transition from high to low, back to high
	err := pin.SetInterrupt(m.PinFalling, func(p m.Pin) {
		// TODO change to loglevel info
		// fmt.Printf("Button state changed to: %t", pin.Get())
		Tips++
	})
	if err != nil {
		return fmt.Errorf("Failed to set interrupt on pin %v, %v\n", string(pin), err)
	}
	
	return nil
}
