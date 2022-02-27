package main

import (
	"machine"
	"time"
)

const (
	NUM_LEDS        = 8
	SHIFT_REG_FREQ  = 100_000_000 // 100 MHz max frequency of 74HC595
	SHIFT_REG_DELAY = time.Second / SHIFT_REG_FREQ
	DELAY           = time.Millisecond * 500
)

// We model our LEDs as an array of Boolean on/off values
// Whilst a single byte mask would be more efficient for production purposes,
// this will keep the code simpler for this simple example
type LEDS [NUM_LEDS]bool

var (
	pinClock = machine.GP11 // SH_CP [SCK] on 74HC595
	pinLatch = machine.GP12 // ST_CP [RCK] on 74HC595
	pinData  = machine.GP13 // DS [S1] on 74HC595
	leds     LEDS
)

func setupPins() {
	pinOutputConfig := machine.PinConfig{Mode: machine.PinOutput}
	pinClock.Configure(pinOutputConfig)
	pinLatch.Configure(pinOutputConfig)
	pinData.Configure(pinOutputConfig)
	pinLatch.High()
}

func updateShiftRegister() {
	// Set the latch pin LOW to begin shifting data from DS into 74HC595 internal registers
	pinLatch.Low()

	// We are assuming the SIPO shift register has NUM_LEDS outputs
	for i := 0; i < NUM_LEDS; i++ {
		pinClock.Low()              // Set the clock LOW when providing data bit
		pinData.Set(leds[i])        // Shift data in LSB first
		pinClock.High()             // Reset clock to HIGH to shift bit into 74HC595
		time.Sleep(SHIFT_REG_DELAY) // Throttle data bits to shift register's max frequency
	}

	// Transition latch pin HIGH to transfer 74HC595 internal registers to parallel output (LEDs)
	pinLatch.High()
}

func main() {
	setupPins()
	for {
		// Turn all LEDs off
		leds = LEDS{} // Initializes to all false
		updateShiftRegister()
		time.Sleep(DELAY)
		// Now turn them on one-by-one with a short delay between
		for i := 0; i < NUM_LEDS; i++ {
			leds[NUM_LEDS-i-1] = true
			updateShiftRegister()
			time.Sleep(DELAY)
		}
	}
}
