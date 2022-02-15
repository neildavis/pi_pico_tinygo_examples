package main

import "machine"

func ledOff(pins [3]machine.Pin) {
	pins[0].Low()
	pins[1].Low()
	pins[2].Low()
}

func ledRed(pins [3]machine.Pin) {
	pins[0].High()
	pins[1].Low()
	pins[2].Low()
}

func ledGreen(pins [3]machine.Pin) {
	pins[0].Low()
	pins[1].High()
	pins[2].Low()
}

func ledBlue(pins [3]machine.Pin) {
	pins[0].Low()
	pins[1].Low()
	pins[2].High()
}

func ledYellow(pins [3]machine.Pin) {
	pins[0].High()
	pins[1].High()
	pins[2].Low()
}
