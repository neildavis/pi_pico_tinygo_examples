package main

import (
	"machine"
	"time"
)

var (
	pinsIn      = []machine.Pin{machine.GP6, machine.GP7, machine.GP8, machine.GP9}
	currentStep int // state machine 0..7
)

// Note: The 28BJY-48 has 1/64 stepping (5.625°) and 1/64 gearing

const (
	STEP_RPM   = 15                                   // 1/4 revolution per second
	NUM_STEPS  = 4096                                 // Num steps per full revolution.
	STEP_DELAY = time.Minute / (STEP_RPM * NUM_STEPS) // Delay between each step
)

func setupPins() {
	for _, pin := range pinsIn {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		pin.Low()
	}
}

func setPinsForStep(step int) {
	// We only deal with step states 0..7 as those are the 8 'step states'
	// we can choose to move the stepper one increment in either direction
	step = step % 8
	if step < 0 {
		step += 8
	}
	currentStep = step
	// Now energize just the necessary coils to move stepper to next step
	// HIGH (true) = coil energized
	pinStates := [4]bool{} // Initializes all to false (LOW - coil not energized)
	switch step {
	case 0:
		pinStates[0] = true
	case 1:
		pinStates[0] = true
		pinStates[1] = true
	case 2:
		pinStates[1] = true
	case 3:
		pinStates[1] = true
		pinStates[2] = true
	case 4:
		pinStates[2] = true
	case 5:
		pinStates[2] = true
		pinStates[3] = true
	case 6:
		pinStates[3] = true
	case 7:
		pinStates[3] = true
		pinStates[0] = true
	}
	for i, state := range pinStates {
		pinsIn[i].Set(state)
	}
}

// Turn off all coils
func stopStepper() {
	for _, pin := range pinsIn {
		pin.Low()
	}
}

// Move the stepper by numSteps. Negative numbers move backwards
func moveStepper(numSteps int) {
	forwards := !(numSteps < 0)
	if !forwards {
		numSteps = -numSteps
	}
	// Ensure we start from a consistent position
	setPinsForStep(currentStep)
	numSteps += currentStep
	// Step
	for i := currentStep; i < numSteps; i++ {
		time.Sleep(STEP_DELAY)
		if forwards {
			setPinsForStep(currentStep + 1)
		} else {
			setPinsForStep(currentStep - 1)
		}
	}
}

func main() {
	setupPins()
	for {
		// Go forwards one revolution
		moveStepper(NUM_STEPS)
		time.Sleep(time.Second)
		// Go backwards one revolution
		moveStepper(-NUM_STEPS)
		time.Sleep(time.Second)
	}
}
