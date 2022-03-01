package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/l293x"
)

var (
	pinDir1   = machine.GP26
	pinDir2   = machine.GP27
	pinEnable = machine.GP28
	driver    l293x.Device
)

func setupPins() {
	driver = l293x.New(pinDir1, pinDir2, pinEnable)
	driver.Configure()
}

func main() {
	setupPins()
	for {
		// Back and forth
		for i := 0; i < 5; i++ {
			driver.Forward()
			time.Sleep(time.Second * 3)
			driver.Backward()
			time.Sleep(time.Second * 3)
		}
		driver.Stop()
		time.Sleep(time.Second * 3)
	}
}
