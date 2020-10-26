package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"
	"wio/buttons"
	"wio/msgbuf"
	"wio/serial"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/drivers/lis3dh"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

const (
	DisplayLines = 20
	LineHeight   = 10
)

var (
	display = ili9341.NewSPI(
		machine.SPI3,
		machine.LCD_DC,
		machine.LCD_SS_PIN,
		machine.LCD_RESET,
	)

	i2c = machine.I2C0

	backlight = machine.LCD_BACKLIGHT

	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	green = color.RGBA{0, 255, 0, 255}

	width, height int16
)

func initDisplay() (width, height int16) {

	machine.SPI3.Configure(machine.SPIConfig{
		SCK:       machine.LCD_SCK_PIN,
		SDO:       machine.LCD_SDO_PIN,
		SDI:       machine.LCD_SDI_PIN,
		Frequency: 80000000,
	})

	display.Configure(ili9341.Config{})
	display.SetRotation(ili9341.Rotation270)
	display.FillScreen(black)

	backlight.Configure(machine.PinConfig{machine.PinOutput})
	backlight.High()

	return display.Size()
}

func drawBuf(display *ili9341.Device, buf *msgbuf.Msgbuf) {
	display.FillRectangle(0, 0, width, height, black)
	for i, l := range buf.Get(DisplayLines) {
		tinyfont.WriteLine(display, &proggy.TinySZ8pt7b, 0, int16(i*LineHeight)+LineHeight, l, white)
	}
}

func main() {

	buf := msgbuf.NewMsgbuf(DisplayLines)

	buttonsChan := make(chan machine.Pin)
	go buttons.GetInput(buttonsChan, 100)

	uartChan := make(chan string)
	go serial.ReadLine(machine.UART0, uartChan)

	tickChan := make(chan int)
	go func() {
		for {
			time.Sleep(250 * time.Millisecond)
			tickChan <- 0
		}
	}()

	// LIS3DH
	i2c.Configure(machine.I2CConfig{SCL: machine.PIN_WIRE1_SCL, SDA: machine.PIN_WIRE1_SDA})
	accel := lis3dh.New(i2c)
	accel.Configure()

	width, height = initDisplay()

	i := 0
	for {
		select {
		case line := <-uartChan:
			println(">>", line, "<<")
		case <-tickChan:
			x, y, z, _ := accel.ReadAcceleration()
			i++
			println(accel.Connected(), x, y, x)
			buf.Add(fmt.Sprintf("(%d) X: %d Y: %d Z: %d", i, x, y, z))
			drawBuf(display, buf)
		}
	}
}
