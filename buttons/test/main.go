package main

import (
	"machine"

	"time"
	"wio/buttons"
	"wio/serial"
)

func main() {

	buttonsChan := make(chan machine.Pin)
	go buttons.GetInput(buttonsChan, 100)

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
		case b := <-buttonsChan:
			println(">>", b)
			for _, x := range buttons.PollInputs(buttons.AllInputs) {
				println("-->", x)
			}
		case <-tickChan:
			println("Tick")
		}
	}
}
