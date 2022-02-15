package main

import (
	"machine"
	"time"

	"github.com/neildavis/tinygo_modules/passivebuzzer"
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
	keypad keypad4x4.Device
	buzzer passivebuzzer.Device
	// 4x4 matrix of tones to play corresponding to keypad buttons
	tones = [4][4]passivebuzzer.Note{
		{passivebuzzer.NOTE_C4, passivebuzzer.NOTE_D4, passivebuzzer.NOTE_E4, passivebuzzer.NOTE_F4},
		{passivebuzzer.NOTE_G4, passivebuzzer.NOTE_A4, passivebuzzer.NOTE_B4, passivebuzzer.NOTE_C5},
		{passivebuzzer.NOTE_D5, passivebuzzer.NOTE_E5, passivebuzzer.NOTE_F5, passivebuzzer.NOTE_G5},
		{passivebuzzer.NOTE_A5, passivebuzzer.NOTE_B5, passivebuzzer.NOTE_C6, passivebuzzer.NOTE_D6},
	}
	// Duration of passivebuzzer to play
	duration = time.Millisecond * 500
)

func setupPins() {
	// Configure keypad
	keypad = keypad4x4.NewDevice(pinKPR1, pinKPR2, pinKPR3, pinKPR4,
		pinKPC1, pinKPC2, pinKPC3, pinKPC4)
	keypad.Configure()
	// Configure passive buzzer as passivebuzzer device
	buzzer = passivebuzzer.New(pinBuzzer, pwmGroup)
}

func main() {
	setupPins()
	// Loop forever reading a keypress from the pad and playing the matching passivebuzzer
	for {
		row, col := keypad.GetIndices()
		if row > -1 && col > -1 {
			buzzer.Note(tones[row][col], duration)
		}
	}
}
