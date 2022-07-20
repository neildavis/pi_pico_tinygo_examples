package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/irremote"
)

var (
	pinIRSend = machine.GP7
	pwmIRSend = machine.PWM3
	irSend    irremote.SenderDevice
)

var irCmdButtons = []byte{
	0x45, 0x46, 0x47,
	0x44, 0x40, 0x43,
	0x07, 0x15, 0x09,
	0x16, 0x19, 0x0D,
	0x0C, 0x18, 0x5E,
	0x08, 0x1C, 0x5A,
	0x42, 0x52, 0x4A,
}

func setupPins() {
	irConfig := irremote.SenderConfig{
		Pin: pinIRSend,
		PWM: pwmIRSend,
		// Default 33% duty cycle for modulation
	}
	irSend = irremote.NewSender(irConfig)
	irSend.Configure()
}

func main() {
	setupPins()
	for {
		addr := uint16(0)
		for _, cmd := range irCmdButtons {
			irSend.SendNEC(addr, cmd, true)
			time.Sleep(time.Millisecond * 100)
			irSend.StopNECRepeats()
			time.Sleep(time.Millisecond * 500)
			addr += 1
		}
		time.Sleep(time.Second * 3)
	}
}
