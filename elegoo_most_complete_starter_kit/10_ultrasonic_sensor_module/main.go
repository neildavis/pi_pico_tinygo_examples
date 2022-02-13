package main

import (
	"machine"

	"tinygo.org/x/drivers/hcsr04"
)

var (
	pinLEDR    = machine.GP0
	pinLEDG    = machine.GP2
	pinLEDB    = machine.GP4
	pinUSMTrig = machine.GP14
	pinUSMEcho = machine.GP15
	usmDevice  hcsr04.Device
)

// Setup pins for RGB LED and configure USM device via driver
func setupPins() {
	pinOutConfig := machine.PinConfig{Mode: machine.PinOutput}
	pinLEDR.Configure(pinOutConfig)
	pinLEDG.Configure(pinOutConfig)
	pinLEDB.Configure(pinOutConfig)
	usmDevice = hcsr04.New(pinUSMTrig, pinUSMEcho)
	usmDevice.Configure()
}

// Set the RGB LED to Red
func ledRed() {
	pinLEDR.High()
	pinLEDG.Low()
	pinLEDB.Low()
}

// Set the RGB LED to Yellow
func ledYellow() {
	pinLEDR.High()
	pinLEDG.High()
	pinLEDB.Low()
}

// Set the RGB LED to Green
func ledGreen() {
	pinLEDR.Low()
	pinLEDG.High()
	pinLEDB.Low()
}

func main() {
	// Configure pins
	setupPins()
	// Loop forever
	for {
		distanceMm := usmDevice.ReadDistance()
		if distanceMm < 100 {
			ledRed()
		} else if distanceMm < 200 {
			ledYellow()
		} else {
			ledGreen()
		}
	}
}
