package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/easystepper"
	"tinygo.org/x/drivers/irremote"
)

var (
	// IR Receiver
	pinIROut = machine.GP3
	ir       irremote.IRReceiverDevice
	// Stepper Motor
	pinStepper1 = machine.GP6
	pinStepper2 = machine.GP7
	pinStepper3 = machine.GP8
	pinStepper4 = machine.GP9
	stepper     easystepper.Device
	// A channel of steps to move
	ch chan int32
)

const (
	STEP_RPM  = 15   // 1/4 revolution per second
	NUM_STEPS = 4096 // Num steps per full revolution.
)

func setupPins() {
	// Setup IR receiver
	ir = irremote.New(pinIROut)
	ir.Configure()
	// Setup Stepper Motor
	stepper = easystepper.NewWithMode(pinStepper1, pinStepper3, pinStepper2, pinStepper4, NUM_STEPS, STEP_RPM, easystepper.EightStepMode)
	stepper.Configure()
}

// Handle a callback from the IR receiver
func irCallback(code uint32, addr uint16, cmd uint8, repeat bool) {
	switch cmd {
	case 0x62:
		// VOL+ button pressed. Go forwards one revolution
		ch <- NUM_STEPS
	case 0xA8:
		// VOL- button pressed. Go backwards one revolution
		ch <- -NUM_STEPS
	default:
		break
	}
}

func main() {
	setupPins()
	// Create a buffered channel of steps to move
	ch = make(chan int32, 10)
	// Register for IR callbacks
	ir.Callback(irCallback)
	for {
		// Read a step count from the channel
		steps := <-ch
		stepper.Move(steps)
		time.Sleep(time.Millisecond * 10)
	}
}
