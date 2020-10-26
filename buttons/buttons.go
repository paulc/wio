package buttons

import (
	"machine"
	"time"
)

const (
	Button1 = machine.BUTTON_1
	Button2 = machine.BUTTON_2
	Button3 = machine.BUTTON_3
	Up      = machine.WIO_5S_UP
	Down    = machine.WIO_5S_DOWN
	Left    = machine.WIO_5S_LEFT
	Right   = machine.WIO_5S_RIGHT
	Press   = machine.WIO_5S_PRESS
)

var (
	AllInputs = []machine.Pin{Button1, Button2, Button3, Up, Down, Left, Right, Press}
)

func init() {
	for _, b := range AllInputs {
		b.Configure(machine.PinConfig{Mode: machine.PinInput})
	}
}

func PollInputs(inputs []machine.Pin) (r []machine.Pin) {
	for _, b := range inputs {
		if !b.Get() {
			r = append(r, b)
		}
	}
	return
}

func GetInput(c chan<- machine.Pin, poll time.Duration) {
	for {
		for _, b := range AllInputs {
			if !b.Get() {
				c <- b
			}
		}
		time.Sleep(time.Millisecond * poll)
	}
}
