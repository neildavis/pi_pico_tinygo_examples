package main

import "machine"

var (
	pinLED    = machine.LED
	pinSwitch = machine.GP10
)

func setupPins() {
	// Configure LED as a digital output
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// Configure tilt ball switch as a digital input (pulled high)
	pinSwitch.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
}

func main() {
	// Configure our pins
	setupPins()
	// Loop forever
	for {
		//  When the switch opens, signal is pulled HIGH and LED is turned off
		pinLED.Set(!pinSwitch.Get())
	}
}
