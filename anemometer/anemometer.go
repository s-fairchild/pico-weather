//go:build anemometer

package anemometer

import (
	m "machine"
	"fmt"
)

var spins int
const Enabled = true
var Initialized bool

func Monitor() error {

	pin := m.GPIO17
	pin.Configure(m.PinConfig{
		Mode: m.PinInputPullup,
	})

	err := pin.SetInterrupt(m.PinLevelLow, func (p m.Pin)  {
		spins++
	})
	if err != nil {
		return fmt.Errorf("Failed to set interrupt on pin %v, error: %v", pin, err)
	}

	return nil
}