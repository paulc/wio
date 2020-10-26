package serial

import (
	"machine"
	"time"
)

func GetChar(uart machine.USBCDC) (c byte) {
	for uart.Buffered() == 0 {
		time.Sleep(time.Millisecond * 10)
	}
	c, _ = uart.ReadByte()
	return
}

func ReadLine(uart machine.USBCDC, c chan<- string) {
	for {
		line := make([]byte, 0, 80)
		for c := GetChar(uart); c != '\n'; c = GetChar(uart) {
			if c == '\r' {
				continue
			}
			line = append(line, c)
		}
		c <- string(line)
	}
}
