package accel

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/lis3dh"
)

type AccelData struct {
	x, y, z int32
}

var (
	accel lis3dh.Device
)

func init() {
	i2c := machine.I2C0
	i2c.Configure(machine.I2CConfig{SCL: machine.PIN_WIRE1_SCL, SDA: machine.PIN_WIRE1_SDA})
	accel = lis3dh.New(i2c)
	accel.Configure()
}

func GetAccel(c chan<- AccelData, poll time.Duration) {
	for {
		x, y, z, _ := accel.ReadAcceleration()
		c <- AccelData{x: x, y: y, z: z}
		time.Sleep(poll)
	}
}
