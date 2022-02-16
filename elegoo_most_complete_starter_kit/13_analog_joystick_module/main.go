package main

import (
	"machine"
	"math"
	"time"
)

// Since machine.pwmGroup is not exported, we create our own type to allow it to be passed around & stored
type PWM interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Set(channel uint8, value uint32)
	Top() uint32
}

var (
	// Joystick
	pinX             = machine.ADC{Pin: machine.ADC0} // analog X
	pinY             = machine.ADC{Pin: machine.ADC1} // analog Y
	pinZ             = machine.GP22                   // digital - push switch
	originX, originY int
	// LED {R,G,B} PWM
	pinsRGB = [3]machine.Pin{machine.GP0, machine.GP1, machine.GP2}
	pwmsRGB = [3]PWM{machine.PWM0, machine.PWM0, machine.PWM1}
)

func setupPins() {
	pinX.Configure(machine.ADCConfig{})
	pinY.Configure(machine.ADCConfig{})
	pinZ.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	for i := 0; i < len(pinsRGB); i++ {
		pinsRGB[i].Configure(machine.PinConfig{Mode: machine.PinPWM})
		pwmsRGB[i].Configure(machine.PWMConfig{Period: uint64(1e9 / 500)}) // 500 Hz
	}
}

// The raw values read from ADC for joystick x,y are 16-bit in range 0...math.MaxUint16
// We'll reduce the precision to 10 bit 0...1024
func joyPosRaw() (int, int) {
	machine.InitADC()
	return int(pinX.Get()) >> 6, int(pinY.Get()) >> 6
}

// Set the origin to the current joystick values
func joyCalibrate() {
	originX, originY = 0, 0
	// average over 10 readings
	for i := 0; i < 10; i++ {
		x, y := joyPosRaw()
		originX = (originX*i + x) / (i + 1)
		originY = (originY*i + y) / (i + 1)
		time.Sleep(time.Millisecond * 10)
	}
}

// Joystick position adjusted for calibrated origin
func joyPos() (int, int) {
	x, y := joyPosRaw()
	// adjust to origin
	x, y = x-originX, y-originY
	// Create a 'dead zone' by removing jitter close to origin (0,0)
	if math.Abs(float64(x)) < 10 {
		x = 0
	}
	if math.Abs(float64(y)) < 10 {
		y = 0
	}
	return x, y
}

func clockFaceAngle(x, y float64) float64 {
	angleRad := math.Pi/2.0 - math.Atan(y/x)
	if x < 0 {
		angleRad += math.Pi
	}
	return angleRad // radians
}

func rgbMultipliers(angleRad float64) [3]float64 {
	// Each colour component (R,G,B) contributes only in their own unique 2/3 (4*PI/3 rads) region of the circle
	var ret = [3]float64{0, 0, 0} // R,G,B
	// The contribution multiplier (0.0...1.0) depends on how 'far' the angle is from the maximum point for that colour
	var maxRads = [3]float64{0, 2 * math.Pi / 3, 4 * math.Pi / 3} // R, G, B
	for i := 0; i < 3; i++ {
		// Compute each multiplier in turn r,g,b
		diff := math.Abs(math.Mod(angleRad-maxRads[i], math.Pi*2)) // how far from max for this component is the angle?
		// Reduce the multiplier for this colour component linearly
		ret[i] = math.Max(0, 1.0-diff/(math.Pi*2/3))
	}
	return ret
}

func main() {
	setupPins()
	joyCalibrate()
	// Get the joysick position offset from origin (0,0)
	for {
		x, y := joyPos()
		if !pinZ.Get() {
			// Stick press down. Turn LED off
			x, y = 0, 0
		}
		// Calulate the clockwise angle in radians of the vector (12'o'clock = (0, Ymax) = 0/2*PI)
		angle := clockFaceAngle(float64(x), float64(y))
		// Get the RGB multipliers for each colour component. This determines the colour {R,G,B}
		rgbMultipliers := rgbMultipliers(angle)
		// Get the normalized magnitude of the vector ((0,0),(x,y)). This controls the brightness
		magnitude := math.Sqrt(float64(x*x+y*y)) / float64(originX)
		// Set the RGB LED colour/brightness with PWM
		for i := 0; i < len(pwmsRGB); i++ {
			multiplier := magnitude * rgbMultipliers[i]
			pwmChan, _ := pwmsRGB[i].Channel(pinsRGB[i])
			pwmsRGB[i].Set(pwmChan, uint32(float64(pwmsRGB[i].Top())*multiplier))
		}
		time.Sleep(time.Millisecond * 10)
	}
}
