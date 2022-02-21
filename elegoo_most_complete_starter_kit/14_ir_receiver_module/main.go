package main

import (
	"machine"
	"time"

	"github.com/neildavis/tinygo_modules/passivebuzzer"
	"tinygo.org/x/drivers/irremote"
)

var (
	// IR Receiver
	pinIROut = machine.GP26
	ir       irremote.IRReceiverDevice
	// Passive buzzer
	pinBuzzer = machine.GP12
	pwmGroup  = machine.PWM6
	buzzer    passivebuzzer.Device
	// Duration of passivebuzzer to play
	duration = time.Millisecond * 500
	// A channel of notes to play
	ch chan passivebuzzer.Note
)

func setupPins() {
	// Setup IR receiver
	ir = irremote.New(pinIROut)
	ir.Configure()
	// Configure passive buzzer
	buzzer = passivebuzzer.New(pinBuzzer, pwmGroup)
	buzzer.Configure()
}

// Handle a callback from the IR receiver
func irCallback(code uint32, addr uint16, cmd uint8, repeat bool) {
	note := irCmdButtons[cmd]
	if note != 0 {
		// Send the note to the channel. Don't play here since we're in an interrupt handler
		ch <- note
	}
}

func main() {
	setupPins()
	// Create a buffered channel of notes to play
	ch = make(chan passivebuzzer.Note, 10)
	// Register for IR callbacks
	ir.Callback(irCallback)
	for {
		// Read a note from the channel and play it on the buzzer
		note := <-ch
		buzzer.Note(note, duration)
	}
}
