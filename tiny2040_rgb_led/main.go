package main

import (
	"machine"
	"time"
)

var (
	pinRed   = machine.GP18
	pwmRed   = machine.PWM1
	pinGreen = machine.GP19
	pwmGreen = machine.PWM1
	pinBlue  = machine.GP20
	pwmBlue  = machine.PWM2
)

func setupPWM() {
	// Setup pins for PWM
	pinConfig := machine.PinConfig{Mode: machine.PinPWM}
	pinRed.Configure(pinConfig)
	pinGreen.Configure(pinConfig)
	pinBlue.Configure(pinConfig)
	// Setup PWM Groups
	period := uint64(1e9 / 500)
	pwmConfig := machine.PWMConfig{Period: period}
	pwmRed.Configure(pwmConfig)
	pwmGreen.Configure(pwmConfig)
	pwmBlue.Configure(pwmConfig)
}

func main() {
	// Initialize pins for PWM
	setupPWM()
	// Get PWM channels for pins
	chanRed, _ := pwmRed.Channel(pinRed)
	chanGreen, _ := pwmGreen.Channel(pinGreen)
	chanBlue, _ := pwmBlue.Channel(pinBlue)
	// Set up initial state values & loop forever
	redValue, greenValue, blueValue := 255, 0, 0
	// These are pointers to the above values that we will increment/decrement each phase
	incVal, decVal := &greenValue, &redValue
	for {
		// Loop over three transition phases
		for j := 0; j < 3; j++ {
			switch j {
			case 0:
				// increasing green and decreasing red
				incVal, decVal = &greenValue, &redValue
			case 1:
				// increasing blue and decreasing green
				incVal, decVal = &blueValue, &greenValue
			case 2:
				// increasing red and decreasing blue
				incVal, decVal = &redValue, &blueValue
			}
			// Perform the transition for this phase over 255 steps with 5 ms pause between steps
			for i := 0; i < 255; i++ {
				pwmRed.Set(chanRed, pwmRed.Top()*uint32(redValue)/255/2)
				pwmGreen.Set(chanGreen, pwmGreen.Top()*uint32(greenValue)/255/2)
				pwmBlue.Set(chanBlue, pwmBlue.Top()*uint32(blueValue)/255/2)
				*incVal++
				*decVal--
				time.Sleep(time.Millisecond * 5)
			}
		}
	}
}
