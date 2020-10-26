package main

import (
	"machine"
	"time"
	"wio/serial"
)

func main() {

	uartChan := make(chan string)
	go serial.ReadLine(machine.UART0, uartChan)

	tickChan := make(chan int)
	go func() {
		for {
			time.Sleep(time.Second)
			tickChan <- 0
		}
	}()

	for {
		select {
		case line := <-uartChan:
			println(">>", line, "<<")
		case <-tickChan:
			println("Tick")
		}
	}
}
