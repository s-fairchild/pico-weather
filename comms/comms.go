package comms

import (
	m "machine"

	t "github.com/s-fairchild/pico-weather/types"
)

// InitUART Configures the pico's uart0 as the stdout for text, and uart1 is configured for transmitting data
// to the receiving system
func InitUART() error {

	// Text console output
	err := m.UART0.Configure(m.UARTConfig{})
	if err != nil {
		println("error configuring Serial UART0: ")
		return err
	}
	println("Initialized Serial UART0")

	// Data output
	err = m.UART1.Configure(m.UARTConfig{TX: m.UART1_TX_PIN, RX: m.UART1_RX_PIN})
	if err != nil {
		println("error Configuring Serial UART1: ")
		return err
	}
	println("Initialized Serial UART1")
	return nil
}

// InitI2c initializes the default UART1 pins
func InitI2c() {

	err := m.I2C1.Configure(m.I2CConfig{})
	if err != nil {
		println("Failed to Configure I2C1: ", err)
	} else {
		println("initialized I2C1")
	}
}

// txSerialData marshals and writes to UART1
func TxSerialData(r *t.SensorReadings) {

	txData, err := r.MarshalJSON()
	if err != nil {
		println("Error marshaling JSON", err)
	} else if len(txData) == 0 {
		println("txData has a length of ", len(txData), "not transmitting")
	} else {
		println("Transmitting JSON over UART, length of: ", len(txData))
		m.UART1.Write(txData)
	}
}
