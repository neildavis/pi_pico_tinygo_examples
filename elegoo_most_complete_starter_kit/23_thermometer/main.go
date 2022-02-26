package main

import (
	"fmt"
	"machine"
	"math"
	"time"
)

var (
	adcThermistor = machine.ADC{Pin: machine.ADC1}
)

// Constants that we will use in the beta equation to work out the temperature form the thermistor voltage divider
const (
	// Num samples to take average temp from thermistor
	SAMPLE_NUMBER = 10
	// Value for R2 that makes up the voltage divider with thermistor as R1
	BALANCE_RESISTOR = 9965.0
	// The maximum ADC value we can read at 3.3V
	MAX_ADC = float32(math.MaxUint16)
	// Thermistor Beta value from datasheet
	BETA = 3950.0
	// The 'room temp' for which the thermistor resistance is known from the datasheet (25°C)
	ROOM_TEMP = 298.15
	// The known resistance for the thermistor at its specified 'room temp'
	RESISTOR_ROOM_TEMP = 10000.0
)

func setupPins() {
	adcThermistor.Configure(machine.ADCConfig{})
	dht11Setup()
	lcdSetup()
}

func tempFromThermistor() float32 {
	// Take a number of samples from the analog read and average them out
	adcTotal := uint32(0) // do calc in 32-bits to avoid overflow
	for i := 0; i < SAMPLE_NUMBER; i++ {
		machine.InitADC()
		adcTotal += uint32(adcThermistor.Get())
		time.Sleep(time.Millisecond * 10)
	}
	adcAvg := float32(adcTotal / SAMPLE_NUMBER)
	// Calculate the resistance of the thermistor
	rThermistor := float64(BALANCE_RESISTOR * ((MAX_ADC / adcAvg) - 1))
	// Now use the beta equation to calculate the temperature (in Kelvin)
	tKelvin := float32((BETA * ROOM_TEMP) /
		(BETA + (ROOM_TEMP * math.Log(rThermistor/RESISTOR_ROOM_TEMP))))
	// Convert from Kelvin to Celsius
	return tKelvin - 273.15
}

func main() {
	// Configure all our devices
	setupPins()
	// Loop taking temp measurements and displaying on LCD
	for {
		// Read every two seconds - including initial wait for DHT11
		time.Sleep(time.Second * 2)
		// Read from DHT11
		dhtTemp, _, _ := dht11Read()
		// Read from thermistor
		thermistorTemp := tempFromThermistor()
		// Write values to LCD
		lcdWriteLines(
			fmt.Sprintf("DHT11: %.02f C", dhtTemp),
			fmt.Sprintf("Therm: %.02f C", thermistorTemp))
	}
}
