package main

import (
	"fmt"
	"time"
	"wio/display"
	"wio/statusbar"

	"machine"

	"tinygo.org/x/drivers/lis3dh"
)

var (
	accel lis3dh.Device
)

type AccelData struct {
	X, Y, Z int32
}

func InitAccel() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{SCL: machine.PIN_WIRE1_SCL, SDA: machine.PIN_WIRE1_SDA})
	accel = lis3dh.New(i2c)
	accel.Configure()
}

func GetAccel() AccelData {
	x, y, z, _ := accel.ReadAcceleration()
	println(x, y, z)
	return AccelData{X: x, Y: y, Z: z}
}

func main() {

	InitAccel()
	display.InitDisplay()

	for {
		a := GetAccel()
		s := fmt.Sprintf("x: %-4d y: %-4d z: %-4d", a.X/10000, a.Y/10000, a.Z/10000)
		statusbar.Status(s, statusbar.White, statusbar.Red)
		time.Sleep(time.Millisecond * 50)
	}

}
