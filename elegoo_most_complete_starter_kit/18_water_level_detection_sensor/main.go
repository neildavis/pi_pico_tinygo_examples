package main

import (
	"machine"
	"time"
)

const (
	// Water level sensor calibration - these values were achieved empirically using 3.3V to power the WLS
	WLS_LVL_MAX = uint16(20000) // when the WLS tracks are fully submerged
	WLS_LVL_MIN = uint16(10000) // when the WLS tracks are just touching the liquid - NOT when completely dry and/or out of liquid
	// How long to hold the previous max level indication for
	MAX_LEVEL_HOLD_TIME = time.Second * 5
)

var (
	// Water Level
	adcWLS = machine.ADC{Pin: machine.ADC2}
)

func setupPins() {
	// Water Level
	adcWLS.Configure(machine.ADCConfig{})
	// LED Matrix
	setupLEDMatrix()
}

func updateLEDMatrix(lvlNow, lvlMax uint16) {
	clearLED()
	// We will use the lower 4 rows to show the current level
	// and the upper 4 rows to show the max level
	rowValNow := mapLevelToLEDRowValue(lvlNow)
	rowValMax := mapLevelToLEDRowValue(lvlMax)
	setLEDRow(0, rowValNow)
	setLEDRow(1, rowValNow)
	setLEDRow(2, rowValNow)
	setLEDRow(3, rowValNow)
	setLEDRow(4, rowValMax)
	setLEDRow(5, rowValMax)
	setLEDRow(6, rowValMax)
	setLEDRow(7, rowValMax)
}

func mapLevelToLEDRowValue(level uint16) byte {
	rowVal := 0b11111111 // 0xFF - Assume 'full' to start
	// Loop over 8 thresholds shifting rowVal right until level > threshold
	WLS_RANGE := WLS_LVL_MAX - WLS_LVL_MIN + 1
	i := uint16(0)
	for ; i < 9; i++ {
		threshold := WLS_LVL_MAX - WLS_RANGE*i/8
		if level > threshold {
			break
		}
	}

	return byte((rowVal >> i) & 0xFF)
}

func main() {
	setupPins()
	lvlMax := uint16(0) // Track the max level we have reached
	timeMax := time.Now()
	for {
		// Drop the max level after 30s
		now := time.Now()
		if now.Sub(timeMax) > MAX_LEVEL_HOLD_TIME {
			lvlMax = 0
			timeMax = now
		}
		// Take a reading from the water level sensor
		machine.InitADC()
		lvlNow := adcWLS.Get()
		// Update max
		if lvlNow >= lvlMax {
			lvlMax = lvlNow
			timeMax = now
		}
		updateLEDMatrix(lvlNow, lvlMax)
		time.Sleep(time.Millisecond * 500)
	}
}
