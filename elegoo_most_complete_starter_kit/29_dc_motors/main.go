package main

import (
	"machine"
	"math"
	"time"
)

var (
	pinDir1   = machine.GP26
	pinDir2   = machine.GP27
	pinEnable = machine.GP28 // PWM6A
	pwmEnable = machine.PWM6
)

const (
	DELAY = time.Second * 3
)

func setupPins() {
	pinDir1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinDir2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinEnable.Configure(machine.PinConfig{Mode: machine.PinPWM})
	pwmEnable.Configure(machine.PWMConfig{})
}

// setMotorDirAndSpeed sets motor direction and speed (in percent)
// Negative values reverse direction.
func setMotorDirAndSpeed(speed int) {
	if speed > 100 {
		speed = 100
	}
	if speed < -100 {
		speed = -100
	}
	pinDir1.Set(speed >= 0)
	pinDir2.Set(speed < 0)
	absSpeed := uint32(math.Abs(float64(speed)))
	ch, _ := pwmEnable.Channel(pinEnable)
	pwmEnable.Set(ch, pwmEnable.Top()*absSpeed/100)
}

func main() {
	setupPins()
	for {
		// Forward at 50%
		setMotorDirAndSpeed(50)
		time.Sleep(DELAY)
		// Forward at 75%
		setMotorDirAndSpeed(75)
		time.Sleep(DELAY)
		// Forward at 100%
		setMotorDirAndSpeed(100)
		time.Sleep(DELAY)
		// Stop
		setMotorDirAndSpeed(0)
		time.Sleep(DELAY)
		// Reverse at 50%
		setMotorDirAndSpeed(-50)
		time.Sleep(DELAY)
		// Reverse at 75%
		setMotorDirAndSpeed(-75)
		time.Sleep(DELAY)
		// Reverse at 100%
		setMotorDirAndSpeed(-100)
		time.Sleep(DELAY)
		// Stop
		setMotorDirAndSpeed(0)
		time.Sleep(DELAY)
	}
}
