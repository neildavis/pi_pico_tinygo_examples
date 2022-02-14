package main

import (
	"machine"
	"time"

	"github.com/neildavis/tinygo_modules/tone"
)

var (
	// Notes in the melody
	melody = []tone.Note{tone.NOTE_C5, tone.NOTE_D5, tone.NOTE_E5, tone.NOTE_F5, tone.NOTE_G5, tone.NOTE_A5, tone.NOTE_B5, tone.NOTE_C6}
	// duration of each note
	duration = time.Millisecond * 500
	// Pin for passive buzzer pwm
	pinBuzzer = machine.GP12
	// PWM Group for buzzer
	pwmGroup = machine.PWM6
	// Tone device
	toneDevice tone.ToneDevice
)

func setupPins() {
	toneDevice = tone.NewDevice(pinBuzzer, pwmGroup)
	toneDevice.Configure()
}

func main() {
	setupPins()
	// Loop forever coz you just can't get enough of this melody!
	for {
		// Play each note in turn
		for _, note := range melody {
			toneDevice.Tone(note, duration)
		}
		time.Sleep(time.Second * 10)
	}
}
