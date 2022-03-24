package main

import (
	"machine"
	"time"

	"github.com/neildavis/tinygo_modules/passivebuzzer"
	"tinygo.org/x/drivers/irremote"
)

var (
	// IR Receiver
	pinIRIn = machine.GP3
	ir      irremote.ReceiverDevice
	// Passive buzzer
	pinBuzzer = machine.GP15
	pwmGroup  = machine.PWM7
	buzzer    passivebuzzer.Device
	// Duration of passivebuzzer to play
	duration = time.Millisecond * 25
	// A channel of IR commands
	ch chan uint16
)

func setupPins() {
	// Setup IR receiver
	ir = irremote.NewReceiver(pinIRIn)
	ir.Configure()
	// Configure passive buzzer
	buzzer = passivebuzzer.New(pinBuzzer, pwmGroup)
	buzzer.Configure()
}

// Handle a callback from the IR receiver
func irCallback(data irremote.Data) {
	ch <- data.Command
}

func main() {
	setupPins()
	// Create a buffered channel of notes to play
	ch = make(chan uint16, 10000)
	// Register for IR callbacks
	ir.SetCommandHandler(irCallback)
	for {
		// Read a note from the channel and play it on the buzzer
		cmd := <-ch
		note := irCmdButtons[cmd]
		if note != passivebuzzer.NOTE_NONE {
			buzzer.Note(note, duration)
		}
	}
}
