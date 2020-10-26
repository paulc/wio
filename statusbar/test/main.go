package main

import (
	"fmt"
	"time"
	"wio/display"
	"wio/statusbar"
)

func main() {
	display.InitDisplay()
	i := 0
	for {
		s := fmt.Sprintf("Status :: >> %d <<", i)
		statusbar.Status(s, statusbar.White, statusbar.Red)
		time.Sleep(time.Millisecond * 100)
		i++
	}
}
