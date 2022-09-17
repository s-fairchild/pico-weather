package comms

import (
	"fmt"
	"log"
	m "machine"
)

// InitUART configures the text output and encoded data output UARTs based on the build tag 
func InitUART() {

	var oppositeUart int8
	if dataUart0 {
		oppositeUart = initDataUart(m.UART0)
	} else {
		oppositeUart = initDataUart(m.UART1)
	}

	initTextUart(oppositeUart)
}

func initDataUart(u *m.UART) int8 {

	var (
		err error
		uartNum int8
	)
	if u == m.UART0 {
		err = m.UART0.Configure(m.UARTConfig{})
		uartNum = 0
	} else if u == m.UART1 {
		err = m.UART1.Configure(m.UARTConfig{})
		uartNum = 1
	}

	if err != nil {
		log.Fatalf("Error configuring text UART%v, %v\n", uartNum, err)
	} else {
		fmt.Printf("initialized text UART%v\n", uartNum)
	}
	return uartNum
}

func initTextUart(u int8) {

	var err error
	if u == 0 {
		err = m.UART1.Configure(m.UARTConfig{})
	} else if u == 1 {
		err = m.UART0.Configure(m.UARTConfig{})
	}

	if err != nil {
		fmt.Printf("Error Configuring Serial %v, %v\n", u, err)
	} else {
		println("initialized Serial %v", u)
	}
}