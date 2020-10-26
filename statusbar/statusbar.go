package statusbar

import (
	"image/color"
	"wio/display"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

var (
	Black = color.RGBA{0, 0, 0, 255}
	White = color.RGBA{255, 255, 255, 255}
	Red   = color.RGBA{255, 0, 0, 255}
	Blue  = color.RGBA{0, 0, 255, 255}
	Green = color.RGBA{0, 255, 0, 255}
)

func Status(msg string, fg color.RGBA, bg color.RGBA) {
	display.Display.FillRectangle(0, 0, 320, 16, bg)
	tinyfont.WriteLine(display.Display, &proggy.TinySZ8pt7b, 10, 10, msg, fg)
}
