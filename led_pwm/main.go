package main

import (
	"machine"
	"time"
)

func main() {
	// Setup pin for PWM
	pin := machine.LED // LED is GPIO25 on Pico
	pin.Configure(machine.PinConfig{Mode: machine.PinPWM})
	// Setup PWM Group
	pwmGroup := machine.PWM4 // GPIO25 corresponds to PWM group 4 according to RP2040 datasheet
	var period uint64 = 1e9 / 500
	err := pwmGroup.Configure(machine.PWMConfig{Period: period})
	if err != nil {
		println(err.Error())
	}
	// Get PWM channel for pin
	pwmChan, err := pwmGroup.Channel(pin)
	if err != nil {
		println(err.Error())
	}
	// Loop forever. We can 'invert' the duty cycle to change from increasing to descreasing values instead of looping twice
	inverting := false
	for {
		pwmGroup.SetInverting(pwmChan, inverting)
		inverting = !inverting
		// Fade the LED from min to max (or vice versa if inverting) via PW
		for i := 0; i < 255; i++ {
			pwmGroup.Set(pwmChan, pwmGroup.Top()*uint32(i)/255)
			time.Sleep(time.Millisecond * 5)
		}
	}
}
