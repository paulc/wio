package main

import (
	"time"
	"wio/accel"
)

func main() {

	accelChan := make(chan accel.AccelData)
	go accel.GetAccel(accelChan, time.Millisecond*250)

	tickChan := make(chan int)
	go func() {
		for {
			time.Sleep(time.Second)
			tickChan <- 0
		}
	}()

	for {
		select {
		case <-tickChan:
			println("Tick")
		}
	}
}
