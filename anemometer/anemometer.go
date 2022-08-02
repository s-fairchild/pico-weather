package anemometer

import (
	m "machine"
	"fmt"
)

var spins int

func Monitor() error {

	pin := m.GPIO17
	pin.Configure(m.PinConfig{
		Mode: m.PinInputPullup,
	})

	err := pin.SetInterrupt(m.PinLevelLow, func (p m.Pin)  {
		spins++
		fmt.Printf("Spins happened! Spins count: %d", spins)
	})
	if err != nil {
		return fmt.Errorf("Failed to set interrupt on pin %v, error: %v", pin, err)
	}

	return nil
}