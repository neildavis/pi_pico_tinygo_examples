package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/dht"
)

var (
	// DHT 11
	pinDHT = machine.GP28
	dht11  dht.Device
)

func dht11Setup() {
	// Configure DHT11 to update automatically
	dht11 = dht.New(pinDHT, dht.DHT11)
	dht11.Configure(dht.UpdatePolicy{
		UpdateTime:          time.Second * 2,
		UpdateAutomatically: false})
}

func dht11Read() (temp, humidity float32, err error) {
	// Loop reading measurements until we don't get an error
	for i := 0; i < 10; i++ {
		err = dht11.ReadMeasurements()
		if nil == err {
			break
		}
		time.Sleep(time.Microsecond * 100)
	}
	temp, _ = dht11.TemperatureFloat(dht.C)
	humidity, _ = dht11.HumidityFloat()
	return
}
