package tone

import (
	"machine"
	"time"
)

// Note type is frequency
type Note uint64

// Frequencies for musical notes
var (
	NOTE_B0  Note = 31
	NOTE_C1  Note = 33
	NOTE_CS1 Note = 35
	NOTE_D1  Note = 37
	NOTE_DS1 Note = 39
	NOTE_E1  Note = 41
	NOTE_F1  Note = 44
	NOTE_FS1 Note = 46
	NOTE_G1  Note = 49
	NOTE_GS1 Note = 52
	NOTE_A1  Note = 55
	NOTE_AS1 Note = 58
	NOTE_B1  Note = 62
	NOTE_C2  Note = 65
	NOTE_CS2 Note = 69
	NOTE_D2  Note = 73
	NOTE_DS2 Note = 78
	NOTE_E2  Note = 82
	NOTE_F2  Note = 87
	NOTE_FS2 Note = 93
	NOTE_G2  Note = 98
	NOTE_GS2 Note = 104
	NOTE_A2  Note = 110
	NOTE_AS2 Note = 117
	NOTE_B2  Note = 123
	NOTE_C3  Note = 131
	NOTE_CS3 Note = 139
	NOTE_D3  Note = 147
	NOTE_DS3 Note = 156
	NOTE_E3  Note = 165
	NOTE_F3  Note = 175
	NOTE_FS3 Note = 185
	NOTE_G3  Note = 196
	NOTE_GS3 Note = 208
	NOTE_A3  Note = 220
	NOTE_AS3 Note = 233
	NOTE_B3  Note = 247
	NOTE_C4  Note = 262
	NOTE_CS4 Note = 277
	NOTE_D4  Note = 294
	NOTE_DS4 Note = 311
	NOTE_E4  Note = 330
	NOTE_F4  Note = 349
	NOTE_FS4 Note = 370
	NOTE_G4  Note = 392
	NOTE_GS4 Note = 415
	NOTE_A4  Note = 440
	NOTE_AS4 Note = 466
	NOTE_B4  Note = 494
	NOTE_C5  Note = 523
	NOTE_CS5 Note = 554
	NOTE_D5  Note = 587
	NOTE_DS5 Note = 622
	NOTE_E5  Note = 659
	NOTE_F5  Note = 698
	NOTE_FS5 Note = 740
	NOTE_G5  Note = 784
	NOTE_GS5 Note = 831
	NOTE_A5  Note = 880
	NOTE_AS5 Note = 932
	NOTE_B5  Note = 988
	NOTE_C6  Note = 1047
	NOTE_CS6 Note = 1109
	NOTE_D6  Note = 1175
	NOTE_DS6 Note = 1245
	NOTE_E6  Note = 1319
	NOTE_F6  Note = 1397
	NOTE_FS6 Note = 1480
	NOTE_G6  Note = 1568
	NOTE_GS6 Note = 1661
	NOTE_A6  Note = 1760
	NOTE_AS6 Note = 1865
	NOTE_B6  Note = 1976
	NOTE_C7  Note = 2093
	NOTE_CS7 Note = 2217
	NOTE_D7  Note = 2349
	NOTE_DS7 Note = 2489
	NOTE_E7  Note = 2637
	NOTE_F7  Note = 2794
	NOTE_FS7 Note = 2960
	NOTE_G7  Note = 3136
	NOTE_GS7 Note = 3322
	NOTE_A7  Note = 3520
	NOTE_AS7 Note = 3729
	NOTE_B7  Note = 3951
	NOTE_C8  Note = 4186
	NOTE_CS8 Note = 4435
	NOTE_D8  Note = 4699
	NOTE_DS8 Note = 4978
)

// Since machine.pwmGroup is not exported, we create our own type to allow it to be passed around & stored
type TonePWM interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Set(channel uint8, value uint32)
	Top() uint32
}

// Public API for the tone device
type ToneDevice interface {
	Tone(note Note, duration time.Duration)
	Configure()
}

// Internal device type. Conforms to the ToneDevice interface
type device struct {
	pin machine.Pin
	pwm TonePWM
}

// C'tor
func NewDevice(pin machine.Pin, pwm TonePWM) ToneDevice {
	return &device{pin, pwm}
}

// Configure device pins
func (device *device) Configure() {
	// Configure the passive buzzer for PWM output
	device.pin.Configure(machine.PinConfig{Mode: machine.PinPWM})

}

// Tone performs a 50% duty cycle PWM signal at the given frequency for the given duration
func (device *device) Tone(note Note, duration time.Duration) {
	device.pwm.Configure(machine.PWMConfig{Period: 1e9 / uint64(note)})
	pwmChan, _ := device.pwm.Channel(device.pin)
	device.pwm.Set(pwmChan, device.pwm.Top()/2)
	time.Sleep(duration)
	device.pwm.Set(pwmChan, 0)
}
