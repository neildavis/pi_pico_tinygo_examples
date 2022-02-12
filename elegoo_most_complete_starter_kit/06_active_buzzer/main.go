package main

import (
	"machine"
	"time"
)

var (
	pinBuzzer = machine.GP12 // physical pin 16
)

func setupPins() {
	pinBuzzer.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func main() {
	// Setup Output pin
	setupPins()
	// Loop forever making an annoying two tone noise!
	for {
		// output the first frequency
		for i := 0; i < 80; i++ {
			pinBuzzer.High()
			time.Sleep(time.Millisecond)
			pinBuzzer.Low()
			time.Sleep(time.Millisecond)
		}
		// output the second frequency
		for i := 0; i < 100; i++ {
			pinBuzzer.High()
			time.Sleep(time.Millisecond * 2)
			pinBuzzer.Low()
			time.Sleep(time.Millisecond * 2)
		}
		// play it again Sam
	}
}
