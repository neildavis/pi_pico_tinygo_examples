package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/servo"
)

var (
	// We will use GP1 which belongs to PWM group 0
	pinPWM             = machine.GP1
	servoPWM servo.PWM = machine.PWM0
	theServo servo.Servo
)

func setupPins() {
	pinPWM.Configure(machine.PinConfig{Mode: machine.PinPWM})
	theServo, _ = servo.New(servoPWM, pinPWM)
}

func main() {
	// Setup the servo for PWM control
	setupPins()
	// Now we can control the servo by adjusting the PWM duty cycle in ms
	// According to the datasheet http://www.ee.ic.ac.uk/pcheung/teaching/DE1_EE/stores/sg90_datasheet.pdf:
	// Position "0" (1.5 ms pulse) is middle
	// "90" (~2ms pulse) is all the way to the right
	// "-90" (~1ms pulse) is all the way to the left.
	for {
		theServo.SetMicroseconds(1500) // move servo to middle position
		time.Sleep(time.Millisecond * 500)
		theServo.SetMicroseconds(1750) // move servo halfway to right
		time.Sleep(time.Millisecond * 500)
		theServo.SetMicroseconds(2000) // move servo fully to right
		time.Sleep(time.Millisecond * 500)
		theServo.SetMicroseconds(1500) // move servo to middle position
		time.Sleep(time.Millisecond * 500)
		theServo.SetMicroseconds(1250) // move servo halfway to left
		time.Sleep(time.Millisecond * 500)
		theServo.SetMicroseconds(1000) // move servo fully to left
		time.Sleep(time.Millisecond * 500)
		// Repeat forever
	}
}
