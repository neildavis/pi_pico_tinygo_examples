package main

import (
	"machine"
	"time"
)

var (
	// Sound Sensor
	adcSndAO = machine.ADC{Pin: machine.ADC0}
	pinSndDO = machine.GP27
	// LED
	pinLED = machine.LED
)

func setupPins() {
	// Sound Sensor
	adcSndAO.Configure(machine.ADCConfig{})
	pinSndDO.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	// LED
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// LED Matrix
	setupLEDMatrix()
}

func calibrateSoundAO() uint16 {
	// calibrate the 'background noise' level for the sound sensor
	// Show some feedback on LED matric whilst calibrating
	setLEDCheckPattern()
	// Wait some time for sound module to settle
	time.Sleep(time.Millisecond * 500)
	// take 50 readings over 1s and take average value
	baseVal := uint32(0)
	for i := uint32(0); i < 50; i++ {
		machine.InitADC()
		baseVal = (baseVal*i + uint32(adcSndAO.Get())) / (i + 1)
		time.Sleep(time.Second / 50)
	}
	return uint16(baseVal)
}

func main() {
	setupPins()
	clearLED()
	minAO, maxAO := uint16(0xffff), uint16(0)
	// Calibrate sound sensor for background noise
	baseAO := calibrateSoundAO()
	sm := SoundMeter{baseLevel: baseAO, maxLevel: baseAO + 0x1000}
	for {
		// Read analog output AO from sound sensor
		machine.InitADC()
		analogValue := adcSndAO.Get()
		if analogValue > maxAO {
			maxAO = analogValue
		}
		if analogValue < minAO {
			minAO = analogValue
		}
		// Turn the LED on if digital threshold HIGH
		pinLED.Set(pinSndDO.Get())
		// Update sound meter
		sm.AddSample(analogValue)
		sm.Display()
		// Sleep a small amount of time
		time.Sleep(time.Second / 16) // 16 samples / sec
	}
}
