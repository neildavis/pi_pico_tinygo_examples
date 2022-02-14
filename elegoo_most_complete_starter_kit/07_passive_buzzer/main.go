package main

import (
	"machine"
	"time"

	"github.com/neildavis/tinygo_modules/passivebuzzer"
)

var (
	// Notes in the melody
	melody = []passivebuzzer.Note{passivebuzzer.NOTE_C5, passivebuzzer.NOTE_D5, passivebuzzer.NOTE_E5, passivebuzzer.NOTE_F5, passivebuzzer.NOTE_G5, passivebuzzer.NOTE_A5, passivebuzzer.NOTE_B5, passivebuzzer.NOTE_C6}
	// duration of each note
	duration = time.Millisecond * 500
	// Pin for passive buzzer pwm
	pinBuzzer = machine.GP12
	// PWM Group for buzzer
	pwmGroup = machine.PWM6
	// Buzzer device
	buzzer passivebuzzer.Device
)

func setupPins() {
	buzzer = passivebuzzer.New(pinBuzzer, pwmGroup)
	buzzer.Configure()
}

func main() {
	setupPins()
	// Loop forever coz you just can't get enough of this melody!
	for {
		// Play each note in turn
		for _, note := range melody {
			buzzer.Note(note, duration)
		}
		time.Sleep(time.Second * 10)
	}
}
