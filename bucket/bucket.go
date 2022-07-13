package bucket

import (
	m "machine"
	"fmt"
)

const (
	BucketSizeMm = 0.2794
	BucketSizeInch = 0.011
)

var Tips uint

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

// CalculateRain returns rainfall amount by multiplying the number of tips and size, which should be BucketSizeMm or BucketSizeInch.
// Tips are reset to 0 after.
func CalculateRain(size float32) (float32) {

	amount := float32(Tips) * size
	fmt.Printf("Rainfall in inches: %v\n", amount)
	Tips = 0
	return amount
}
