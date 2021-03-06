package accel

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/lis3dh"
)

type AccelData struct {
	X, Y, Z int32
}

var (
	accel lis3dh.Device
)

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

func GetAccelChan(c chan<- AccelData, poll time.Duration) {
	for {
		c <- GetAccel()
		time.Sleep(poll)
	}
}
