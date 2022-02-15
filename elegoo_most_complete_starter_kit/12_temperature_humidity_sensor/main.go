package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

var (
	// RGB LED 1 (Temp) pins {R,G,B}
	pinsTemp = [3]machine.Pin{machine.GP0, machine.GP1, machine.GP2}
	// RGB LED 2 (Humidity) pins {R,G,B}
	pinsHumidity = [3]machine.Pin{machine.GP3, machine.GP4, machine.GP5}
	// DHT 11
	pinDHT = machine.GP15
	dht11  dht.Device
)

const (
	// +/- thresholds for changing LED colour
	tempThreshold     = 0.2 // degrees celsius
	humidityThreshold = 2.0 // percent
)

func setupPins() {
	// Configure pins for the LEDs
	ledConfig := machine.PinConfig{Mode: machine.PinOutput}
	// - RGB LED 1 (Temp)
	for i := 0; i < len(pinsTemp); i++ {
		pinsTemp[i].Configure(ledConfig)
	}
	// - RGB LED 2 (Humidty)
	for i := 0; i < len(pinsHumidity); i++ {
		pinsHumidity[i].Configure(ledConfig)
	}
	// Configure DHT11 to update manually
	dht11 = dht.New(pinDHT, dht.DHT11)
	dht11.Configure(dht.UpdatePolicy{
		UpdateTime:          time.Second * 2,
		UpdateAutomatically: false})
}

func initialWait() {
	// We need to wait a second or two after power on before reading from the DHT11 to allow it to initialize
	// Show alternating red LEDs during this phase
	for i := 0; i < 2; i++ {
		ledRed(pinsTemp)
		ledOff(pinsHumidity)
		time.Sleep(time.Millisecond * 500)
		ledRed(pinsHumidity)
		ledOff(pinsTemp)
		time.Sleep(time.Millisecond * 500)
	}
	ledOff(pinsTemp)
	ledOff(pinsHumidity)
}

func calibrate() (temp, humidity float32) {
	// Take a few readings - 2s apart - and compute averages to return 'normal' values for temp & humidity
	temp, humidity = 0.0, 0.0
	for i := 0; i < 5; i++ {
		readMeasurements()
		newTemp, err := dht11.TemperatureFloat(dht.C)
		if err == nil {
			temp = (temp*float32(i) + newTemp) / float32(i+1)
		}
		newHumidity, err := dht11.HumidityFloat()
		if err == nil {
			humidity = (humidity*float32(i) + newHumidity) / float32(i+1)
		}
		// Show alternating yellow LEDs during this phase whilst we wait 2s for next reading
		for j := 0; j < 2; j++ {
			ledYellow(pinsTemp)
			ledOff(pinsHumidity)
			time.Sleep(time.Millisecond * 500)
			ledYellow(pinsHumidity)
			ledOff(pinsTemp)
			time.Sleep(time.Millisecond * 500)
		}
	}
	ledGreen(pinsTemp)
	ledGreen(pinsHumidity)
	return
}

func readMeasurements() {
	// Loop reading measurements until we don't get an error
	for {
		err := dht11.ReadMeasurements()
		if nil == err {
			return
		}
		time.Sleep(time.Microsecond * 10)
	}
}

// Convenience method to check a value against it's normal value within a given threshoold tolerance
// Sets LED colour Red if above threshold, Blue if below threshold and Green if within tolerance.
func checkMetric(value, norm, threshold float32, ledPins [3]machine.Pin) {
	if value < norm-threshold {
		ledBlue(ledPins)
	} else if value > norm+threshold {
		ledRed(ledPins)
	} else {
		ledGreen(ledPins)
	}
}

func monitor(normalTemp, normalHumidity float32) {
	for {
		// Read every two seconds
		time.Sleep(time.Second * 2)
		readMeasurements()
		temp, _ := dht11.TemperatureFloat(dht.C)
		checkMetric(temp, normalTemp, tempThreshold, pinsTemp)
		humidity, _ := dht11.HumidityFloat()
		checkMetric(humidity, normalHumidity, humidityThreshold, pinsHumidity)
	}
}

func main() {
	// Configure all peripherals
	setupPins()
	// Initial setup
	initialWait()
	// Calibrate to ambient conditions
	normalTemp, normalHumidity := calibrate()
	// Finally monitor for any chnages
	monitor(normalTemp, normalHumidity)
}
