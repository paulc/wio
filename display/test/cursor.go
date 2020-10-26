package main

import (
	"fmt"
	"image/color"
	"wio/buttons"
	"wio/display"

	"machine"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

var (
	black         = color.RGBA{0, 0, 0, 255}
	white         = color.RGBA{255, 255, 255, 255}
	red           = color.RGBA{255, 0, 0, 255}
	blue          = color.RGBA{0, 0, 255, 255}
	green         = color.RGBA{0, 255, 0, 255}
	width, height int16
)

func min(x, y int16) int16 {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int16) int16 {
	if x > y {
		return x
	} else {
		return y
	}
}

func drawCursor(x, y int16, c color.RGBA) {
	display.Display.DrawFastHLine(max(x-10, 0), min(x+10, width-1), y, c)
	display.Display.DrawFastVLine(x, max(y-10, 0), min(y+10, height-1), c)
}

func main() {

	width, height = display.InitDisplay()

	buttonChan := make(chan machine.Pin)
	go buttons.GetInput(buttonChan, 20)

	x := int16(160)
	y := int16(120)

	for {
		select {
		case b := <-buttonChan:

			// Clear cursor
			drawCursor(x, y, black)

			switch b {
			case buttons.Up:
				y = max(y-1, 20)
			case buttons.Down:
				y = min(y+1, height)
			case buttons.Right:
				x = min(x+1, width)
			case buttons.Left:
				x = max(x-1, 0)
			case buttons.Press:
				x = 160
				y = 120
			}

			drawCursor(x, y, red)

			display.Display.FillRectangle(0, 0, width, 20, black)
			pos := fmt.Sprintf("Cursor: x=%d y=%d", x, y)
			tinyfont.WriteLine(display.Display, &proggy.TinySZ8pt7b, 10, 10, pos, white)
		}
	}

}
