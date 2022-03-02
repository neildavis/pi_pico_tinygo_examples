package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/easystepper"
)

var (
	// Rotary Encoder Pins
	pinRotSw  = machine.GP3
	pinRotDt  = machine.GP4
	pinRotClk = machine.GP5
	// Stepper Motor Pins
	pinStepper1 = machine.GP6
	pinStepper2 = machine.GP7
	pinStepper3 = machine.GP8
	pinStepper4 = machine.GP9
)

const (
	STEP_RPM       = 15   // 1/4 revolution per second for stepper motor
	NUM_STEPS      = 4096 // Num steps per full revolution of stepper motor
	NUM_ROT_INCS   = 20   // Number of increments for a full revolution of rotary encoder
	DEBOUNCE_DELAY = time.Millisecond * 4
	QUEUE_LEN      = 1000 // Max number of step operations to be queued
)

func setupPins() {
	// Setup Rotary Encoder
	pinRotDt.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	pinRotClk.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// Rest of setup is in stepperQueue since it has to be local to go routine
}

// stepperQueue is a go routine to setup the stepper and process a request of step requests from a channel queue
func stepperQueue(stepsQ chan int) {
	// We also need the rotary encoder switch in this go routine to process it
	pinRotSw.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	// Setup Stepper Motor
	stepper := easystepper.NewWithMode(pinStepper1, pinStepper3, pinStepper2, pinStepper4, NUM_STEPS, STEP_RPM, easystepper.EightStepMode)
	stepper.Configure()
	stepperPos := 0
	resetStepper := false
	// Process the queue forever
	for {
		// If the rotary encoder switch is pressed, we drain the queue and revert stepper to start pos
		if !pinRotSw.Get() {
			resetStepper = true
		}
		// process the next item in the steps queue
		if len(stepsQ) > 0 {
			steps := <-stepsQ
			if !resetStepper {
				// If no request to reset has been received, process the step request
				stepperPos = (stepperPos + steps) % NUM_STEPS
				stepper.Move(int32(steps))
			}
		} else if resetStepper {
			// Queue has been drained, and reset requested
			// Return stepper to initial position by shortest route
			if stepperPos < -NUM_STEPS/2 {
				stepperPos += NUM_STEPS
			} else if stepperPos > NUM_STEPS/2 {
				stepperPos -= NUM_STEPS
			}
			stepper.Move(int32(-stepperPos))
			stepperPos = 0
			resetStepper = false
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	setupPins()
	// Create a buffered channel of step movements as a queue
	stepsQ := make(chan int, QUEUE_LEN)
	go stepperQueue(stepsQ)
	lastRotCLK := pinRotClk.Get()
	for {
		// Read the rotary encoder clock pin
		rotCLK := pinRotClk.Get()
		// Process on rotary encoder CLK edge change
		if lastRotCLK != rotCLK {
			lastRotCLK = rotCLK
			if !rotCLK {
				// falling edge
				rotDT := pinRotDt.Get()
				if rotDT {
					// Clockwise
					stepsQ <- NUM_STEPS / NUM_ROT_INCS
				} else {
					// Anti-clockwise
					stepsQ <- -NUM_STEPS / NUM_ROT_INCS
				}
			}
		}
		time.Sleep(DEBOUNCE_DELAY)
	}
}
