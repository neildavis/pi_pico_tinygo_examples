package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/shiftregister"
)

const (
	DELAY = time.Second
)

var (
	// 74HC595 Shift Register
	pinClock       = machine.GP11 // SH_CP [SCK] on 74HC595
	pinLatch       = machine.GP12 // ST_CP [RCK] on 74HC595
	pinData        = machine.GP13 // DS [S1] on 74HC595
	shiftRegDevice shiftregister.Device
	// We model our LEDs as a uint32 (max 32 LEDs), with each bit correpsonding to one LED
	leds uint32
)

func setupPins() {
	shiftRegDevice = *shiftregister.New(shiftregister.EIGHT_BITS, pinLatch, pinClock, pinData)
	shiftRegDevice.Configure()
}

func main() {
	setupPins()
	for {
		// count from 9 to 0 and end with dp
		sevenSegDigits := []uint32{
			DIG_9,
			DIG_8,
			DIG_7,
			DIG_6,
			DIG_5,
			DIG_4,
			DIG_3,
			DIG_2,
			DIG_1,
			DIG_0,
			DEC_PT,
		}
		for i := 0; i < len(sevenSegDigits); i++ {
			time.Sleep(DELAY)
			shiftRegDevice.WriteMask(sevenSegDigits[i])
		}
		time.Sleep(DELAY * 3)
	}
}
