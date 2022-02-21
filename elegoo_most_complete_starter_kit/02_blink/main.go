package main

import (
	"machine"
	"time"
)

var (
	pinLED = machine.LED
)

func setupPins() {
	// Configure the built in LED as a digital output
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func main() {
	// Setup
	setupPins()
	// Now loop forever blinking the LED
	for {
		pinLED.High()
		time.Sleep(time.Second)
		pinLED.Low()
		time.Sleep(time.Second)
	}
}
