package main

import (
	"machine"
	"time"
)

var (
	pinPIR = machine.GP6
	pinLED = machine.LED
)

func setupPins() {
	pinPIR.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func main() {
	// Configure pins for LED and PIR
	setupPins()
	// Turn the LED off
	pinLED.Low()
	for {
		// When PIR detects, the LED is on, and vice versa
		pinLED.Set(pinPIR.Get())
		time.Sleep(time.Millisecond * 100)
	}
}
