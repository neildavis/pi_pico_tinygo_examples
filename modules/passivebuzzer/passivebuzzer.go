package passivebuzzer

import (
	"machine"
	"time"
)

// Since machine.pwmGroup is not exported, we create our own type to allow it to be passed around & stored
type PWM interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Set(channel uint8, value uint32)
	Top() uint32
}

// Public API for the device
type Device interface {
	Configure()
	Frequency(freq uint64, duration time.Duration)
	Note(note Note, duration time.Duration)
}

// Internal device type. Conforms to the Device interface
type device struct {
	pin machine.Pin
	pwm PWM
}

// C'tor
func New(pin machine.Pin, pwm PWM) Device {
	return &device{pin, pwm}
}

// Configure device pins
func (device *device) Configure() {
	// Configure the passive buzzer for PWM output
	device.pin.Configure(machine.PinConfig{Mode: machine.PinPWM})
}

// Frequency performs a 50% duty cycle PWM signal at the given frequency for the given duration
func (device *device) Frequency(freq uint64, duration time.Duration) {
	device.pwm.Configure(machine.PWMConfig{Period: 1e9 / freq})
	pwmChan, _ := device.pwm.Channel(device.pin)
	device.pwm.Set(pwmChan, device.pwm.Top()/2)
	time.Sleep(duration)
	device.pwm.Set(pwmChan, 0)
}

// Convenience method to play a particular musical note for a given duration
func (device *device) Note(note Note, duration time.Duration) {
	device.Frequency(uint64(note), duration)
}
