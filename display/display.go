package display

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ili9341"
)

var (
	Display = ili9341.NewSPI(
		machine.SPI3,
		machine.LCD_DC,
		machine.LCD_SS_PIN,
		machine.LCD_RESET,
	)

	Backlight = machine.LCD_BACKLIGHT
)

func InitDisplay() (width, height int16) {
	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       machine.LCD_SCK_PIN,
		SDO:       machine.LCD_SDO_PIN,
		SDI:       machine.LCD_SDI_PIN,
		Frequency: 40000000,
	})

	Display.Configure(ili9341.Config{})
	Display.SetRotation(ili9341.Rotation270)
	Display.FillScreen(color.RGBA{0, 0, 0, 255})

	Backlight.Configure(machine.PinConfig{machine.PinOutput})
	Backlight.High()

	return Display.Size()
}
