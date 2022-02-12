package main

import (
	"machine"
	"time"
)

func main() {
	// Configure the built in LED as a digital output
	pinLED := machine.LED
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// Now loop forever blinking the LED
	for {
		pinLED.High()
		time.Sleep(time.Second)
		pinLED.Low()
		time.Sleep(time.Second)
	}
}
