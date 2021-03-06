package main

import (
	"machine"
	"time"
)

var (
	// Input buttons
	pinButtonOn  = machine.GPIO9
	pinButtonOff = machine.GPIO8
	// Output LED
	pinLED = machine.LED
)

func setupPins() {
	// Set LED pin to digital output
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pinLED.Set(false)
	// We set both button pins as digital input and pulled high (so signal is LOW grounded when pressed)
	pinConfig := machine.PinConfig{Mode: machine.PinInputPullup}
	pinButtonOn.Configure(pinConfig)
	pinButtonOff.Configure(pinConfig)
	// Rather than continuously loop checking pin level, we'll set an interrupt callback function
	pinButtonOn.SetInterrupt(machine.PinLevelLow, pinCallback)
	// Note: With TinyGo < v0.24.0 the next line fails with ErrNoPinChangeChannel since per pin interrupts are not supported
	// See machine.Pin.SetInterrupt() in machine_rp2040_gpio.go (pinCallbacks[core] is global, not per Pin)
	// Fixed in v0.24.0 https://github.com/tinygo-org/tinygo/commit/8a5ab5ab129b7acaa3a55f964866cac1207fb0f5
	pinButtonOff.SetInterrupt(machine.PinLevelLow, pinCallback)
}

// Callback function when a pin value falls to LOW (it's pulled HIGH so this means the button is pressed)
func pinCallback(pin machine.Pin) {
	// Change the value of the LED according to which button was pressed
	// Note: Currently broken as of TinyGo < 0.24.0 all interrupts send machine.Pin(0xff)
	// See machine.gpioHandleInterrupt() in machine_rp2040_gpio.go
	pinLED.Set(pin == pinButtonOn)
	// pinLED.Set(pin == machine.Pin(0xff)) // uncomment to make the 'on' button work on TinyGo < v0.24.0 ('off' button still broken due to above)
}

func main() {
	setupPins()
	for {
		time.Sleep(time.Millisecond * 100)
	}
}
