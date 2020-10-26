package main

import (
	"fmt"
	"image/color"
	"time"
	"wio/display"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
)

func main() {
	width, height := display.InitDisplay()

	s1 := fmt.Sprintf("Hello: %dx%d", width, height)

	tinyfont.WriteLine(display.Display, &proggy.TinySZ8pt7b, 10, 10, s1, white)

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}
