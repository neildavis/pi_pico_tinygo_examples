package main

import (
	"machine"
	"time"

	"github.com/neildavis/tinygo_modules/tone"
	"tinygo.org/x/drivers/keypad4x4"
)

var (
	// Keypad row pins
	pinKPR1 = machine.GP26
	pinKPR2 = machine.GP22
	pinKPR3 = machine.GP21
	pinKPR4 = machine.GP20
	// Keypad column pins
	pinKPC1 = machine.GP19
	pinKPC2 = machine.GP18
	pinKPC3 = machine.GP17
	pinKPC4 = machine.GP16
	// Passive buzzer pin
	pinBuzzer = machine.GP12
	// Matching PWM group for passive buzzer pin
	pwmGroup = machine.PWM6
	// High level device drivers for keypad and passive buzzer
	keypad     keypad4x4.Device
	toneDevice tone.ToneDevice
	// 4x4 matrix of tones to play corresponding to keypad buttons
	tones = [4][4]tone.Note{
		{tone.NOTE_C4, tone.NOTE_D4, tone.NOTE_E4, tone.NOTE_F4},
		{tone.NOTE_G4, tone.NOTE_A4, tone.NOTE_B4, tone.NOTE_C5},
		{tone.NOTE_D5, tone.NOTE_E5, tone.NOTE_F5, tone.NOTE_G5},
		{tone.NOTE_A5, tone.NOTE_B5, tone.NOTE_C6, tone.NOTE_D6},
	}
	// Duration of tone to play
	duration = time.Millisecond * 500
)

func setupPins() {
	// Configure keypad
	keypad = keypad4x4.NewDevice(pinKPR1, pinKPR2, pinKPR3, pinKPR4,
		pinKPC1, pinKPC2, pinKPC3, pinKPC4)
	keypad.Configure()
	// Configure passive buzzer as tone device
	toneDevice = tone.NewDevice(pinBuzzer, pwmGroup)
}

func main() {
	setupPins()
	// Loop forever reading a keypress from the pad and playing the matching tone
	for {
		row, col := keypad.GetIndices()
		if row > -1 && col > -1 {
			toneDevice.Tone(tones[row][col], duration)
		}
	}
}
