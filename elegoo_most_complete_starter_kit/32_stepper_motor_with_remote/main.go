package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/easystepper"
	"tinygo.org/x/drivers/irremote"
)

var (
	// IR Receiver
	pinIRIn = machine.GP3
	ir      irremote.ReceiverDevice
	// Stepper Motor
	pinStepper1 = machine.GP6
	pinStepper2 = machine.GP8
	pinStepper3 = machine.GP7
	pinStepper4 = machine.GP9
	stepper     *easystepper.Device
	// A channel of steps to move
	ch chan int32
)

const (
	STEP_RPM  = 15   // 1/4 revolution per second
	NUM_STEPS = 4096 // Num steps per full revolution.
)

func setupPins() {
	// Setup IR receiver
	ir = irremote.NewReceiver(pinIRIn)
	ir.Configure()
	// Setup Stepper Motor
	stepperConfig := easystepper.DeviceConfig{
		Pin1: pinStepper1, Pin2: pinStepper2, Pin3: pinStepper3, Pin4: pinStepper4,
		StepCount: NUM_STEPS, RPM: STEP_RPM, Mode: easystepper.ModeEight,
	}

	stepper, _ = easystepper.New(stepperConfig)
	stepper.Configure()
}

// Handle a callback from the IR receiver
func irCallback(data irremote.Data) {
	// Ignore repeats
	if data.Flags&irremote.DataFlagIsRepeat != 0 {
		return
	}
	switch data.Command {
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
	ir.SetCommandHandler(irCallback)
	for {
		// Read a step count from the channel
		steps := <-ch
		// Move the stepper
		stepper.Move(steps)
		time.Sleep(time.Millisecond * 10)
	}
}
