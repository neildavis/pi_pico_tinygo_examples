package main

import (
	"machine"
	"time"

	"github.com/neildavis/tinygo_modules/passivebuzzer"
	"tinygo.org/x/drivers/irremote"
)

var (
	// LED
	pinLED = machine.LED
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
	// LED
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
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

// Blink the LED
func blinkLED() {
	pinLED.High()
	time.Sleep(time.Millisecond * 10)
	pinLED.Low()
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
		blinkLED()
		note := irCmdButtons[cmd]
		if note != passivebuzzer.NOTE_NONE {
			buzzer.Note(note, duration)
		}
	}
}
