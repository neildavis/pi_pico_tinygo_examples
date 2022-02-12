package main

import (
	"machine"
	"time"
)

var (
	// Notes in the melody
	melody = []int{NOTE_C5, NOTE_D5, NOTE_E5, NOTE_F5, NOTE_G5, NOTE_A5, NOTE_B5, NOTE_C6}
	// duration of each note
	duration = time.Millisecond * 500
	// Pin for passive buzzer pwm
	pinBuzzer = machine.GP12
	// PWM Group for buzzer
	pwmGroup = machine.PWM6
)

func setupPins() {
	// Configure the passive buzzer for PWM output
	pinBuzzer.Configure(machine.PinConfig{Mode: machine.PinPWM})
}

func main() {
	setupPins()
	// Loop forever coz you just can't get enough of this melody!
	for {
		// Play each note in turn
		for _, note := range melody {
			tone(note, duration)
		}
		time.Sleep(time.Second * 10)
	}
}
