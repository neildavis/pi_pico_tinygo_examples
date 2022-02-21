package main

import (
	"machine"
	"time"
)

var (
	// Input buttons
	pinButtonOn  = machine.GP9
	pinButtonOff = machine.GP8
	// Output LED
	pinLED = machine.GP5
)

func setupPins() {
	// Set LED pin to digital output
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// We set both button pins as digital inputs and pulled high (so signal is LOW when pressed)
	pinConfig := machine.PinConfig{Mode: machine.PinInputPullup}
	pinButtonOn.Configure(pinConfig)
	pinButtonOff.Configure(pinConfig)
}

func main() {
	setupPins()
	// Note: machine.Pin.SetInterrupt is currently broken (as of TinyGo v0.22.0) for more than one pin so we need to poll instead :(
	for {
		// Poll for button states. Remember they are pulled HIGH so button is pressed when input is LOW
		if !pinButtonOn.Get() {
			pinLED.High()
		} else if !pinButtonOff.Get() {
			pinLED.Low()
		}
		time.Sleep(time.Millisecond * 10)
	}
}
