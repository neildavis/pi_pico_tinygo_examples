package main

import (
	"machine"
	"math"
	"time"

	"tinygo.org/x/drivers/shiftregister"
)

const (
	NUM_LEDS      = 8
	LED_MULTIPLER = 4 // Adjust depending on the ambient light
	DELAY         = time.Millisecond * 500
)

var (
	// Photocell / LDR
	adcLDR = machine.ADC{Pin: machine.ADC0}
	// 74HC595 Shift Register
	pinClock       = machine.GP11 // SH_CP [SCK] on 74HC595
	pinLatch       = machine.GP12 // ST_CP [RCK] on 74HC595
	pinData        = machine.GP13 // DS [S1] on 74HC595
	shiftRegDevice shiftregister.Device
	// We model our LEDs as a uint32 (max 32 LEDs), with each bit correpsonding to one LED
	leds uint32
)

func setupPins() {
	adcLDR.Configure(machine.ADCConfig{})
	shiftRegDevice = *shiftregister.New(shiftregister.EIGHT_BITS, pinLatch, pinClock, pinData)
	shiftRegDevice.Configure()
}

func main() {
	setupPins()
	for {
		// Read LDR
		machine.InitADC()
		ldrValue := adcLDR.Get()
		numLEDsLit := LED_MULTIPLER * int(ldrValue) * NUM_LEDS / math.MaxUint16
		if numLEDsLit > NUM_LEDS {
			numLEDsLit = NUM_LEDS
		}
		// Turn on LEDs
		leds = math.MaxUint32 >> (32 - numLEDsLit)
		shiftRegDevice.WriteMask(leds)
		time.Sleep(DELAY)
	}
}
