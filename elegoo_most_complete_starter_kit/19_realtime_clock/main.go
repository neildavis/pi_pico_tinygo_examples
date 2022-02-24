package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/ds3231"
)

var (
	pinSDA = machine.GP26 // I2C1 SDA
	pinSCL = machine.GP27 // I2C1 SCL
	i2c    = machine.I2C1
	rtc    ds3231.Device
)

func setupPins() {
	i2c.Configure(machine.I2CConfig{SDA: pinSDA, SCL: pinSCL})
	rtc = ds3231.New(i2c)
	rtc.Configure()
}

func main() {
	setupPins()
	setupLEDMatrix()
	//testChars()	// uncomment to watch a runthrough of all suppoerted chars on the LED matrix display
	valid := rtc.IsTimeValid()
	if !valid {
		// First time clock set. Adjust the date here to nearer the actual time
		date := time.Date(2022, 02, 24, 16, 35, 30, 0, time.UTC)
		rtc.SetTime(date)
	}

	running := rtc.IsRunning()
	if !running {
		rtc.SetRunning(true)
	}

	for {
		dt, _ := rtc.ReadTime()
		// Because of our primitive display, by the time we show the seconds it's ~16s after the date was read
		dt = dt.Add(time.Second * 16)
		// Date/Time will be most accurate as seconds are being displayed
		displayText(fmt.Sprintf("D %02d-%02d-%d T %02d:%02d:%02d ", dt.Day(), dt.Month(), dt.Year(), dt.Hour(), dt.Minute(), dt.Second()))
		temp, _ := rtc.ReadTemperature()
		displayText(fmt.Sprintf("C %.2f ", float32(temp)/1000))
	}
}
