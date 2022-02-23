package main

import (
	"machine"
	"math"
	"time"

	"tinygo.org/x/drivers/max72xx"
	"tinygo.org/x/drivers/mpu6050"
)

// Since machine.pwmGroup is not exported, we create our own type to allow it to be passed around & stored
type PWM interface {
	Configure(config machine.PWMConfig) error
	Channel(pin machine.Pin) (channel uint8, err error)
	Set(channel uint8, value uint32)
	Top() uint32
}

var (
	// GY-521
	pinSDA = machine.GP26 // I2C1 SDA
	pinSCL = machine.GP27 // I2C1 SCL
	i2c    = machine.I2C1
	imu    mpu6050.Device
	// LED Matrix
	pinDin = machine.SPI1_SDO_PIN // SPI1 COPI/SDO/TX
	pinCS  = machine.GP9          // SPI1 CSn - machine package does not define a default constant
	pinCLK = machine.SPI1_SCK_PIN // SPI1 SCK
	spi    = machine.SPI1
	matrix *max72xx.Device
	status = [NUM_ROWS]byte{}
	// RGB LED {R,G,B}
	pinsRGB = [3]machine.Pin{machine.GP0, machine.GP1, machine.GP2}
	pwmsRGB = [3]PWM{machine.PWM0, machine.PWM0, machine.PWM1}
)

const (
	// General consts
	DELAY = time.Millisecond * 50
	// GY-521 / I2C consts
	I2C_FREQ = 400_000 // 400 KHz I2C is specified by MPU6050
	// LED Matrix / SPI consts
	SPI_CPOL      = 1        // SPI Clock Polarity: CLK is HIGH when idle
	SPI_CPHA      = 0        // SPI Clock Phase: Data is read on 'falling edge', which is first edge when CPOL is 1
	SPI_FREQ      = 10000000 // 10 MHz is max frequency supportewd by MAX72xx
	SPI_DATA_BITS = 8        // 8-bit bus
	// Our LED Matrix is 8x8
	NUM_ROWS = 8
	NUM_COLS = 8
)

func setupPins() {
	// Configure the GY-521 IMU
	i2c.Configure(machine.I2CConfig{Frequency: I2C_FREQ, SDA: pinSDA, SCL: pinSCL})
	imu = mpu6050.New(i2c)
	imu.Configure()
	// Configure the MAX72XX LED Matrix Display
	spiMode := uint8((SPI_CPOL << 1) | SPI_CPHA)
	spiConfig := machine.SPIConfig{Frequency: SPI_FREQ, LSBFirst: false, Mode: spiMode, DataBits: SPI_DATA_BITS,
		SDO: pinDin, SCK: pinCLK}
	spi.Configure(spiConfig)
	matrix = max72xx.NewDevice(*spi, pinCS)
	matrix.Configure()
	matrix.StopDisplayTest()
	matrix.SetScanLimit(8)
	matrix.SetDecodeMode(0)   // No decoding for matrix mode (used for 7-seg digits)
	matrix.StopShutdownMode() // The MAX72XX is in power-saving mode on startup
	// Configure the RGB LED
	for i := 0; i < len(pinsRGB); i++ {
		pinsRGB[i].Configure(machine.PinConfig{Mode: machine.PinPWM})
		pwmsRGB[i].Configure(machine.PWMConfig{Period: uint64(1e9 / 500)}) // 500 Hz
	}
}

func setDisplayIntensity(intensity uint8) {
	matrix.WriteCommand(max72xx.REG_INTENSITY, intensity)
}

func clearDisplay() {
	for row := 0; row < NUM_ROWS; row++ {
		setRow(row, 0)
	}
}

func setRow(row int, value byte) {
	if row < 0 || row > NUM_ROWS-1 {
		return
	}
	status[row] = value
	matrix.WriteCommand(byte(row+1), value)
}

func drawBall(x, y int, i byte) {
	if x > 6 {
		x = 6
	}
	if y > 6 {
		y = 6
	}
	if i > 0xf {
		i = 0xf
	}
	rowPattern := byte(0b00000011) << y
	setRow(x, rowPattern)
	setRow(x+1, rowPattern)
	setDisplayIntensity(i)
}

func mapAccelerationToPositionAndIntensity(ax, ay, az int32) (x, y int, i byte) {
	// With a 2x2 'ball' on a 8x8 matrix, we can only move 3 units in each direction from center
	// So we'll 'floor'' the input x,y to 1G and cap at +/- 3G
	fx := math.Floor(float64(ax) / 1_000_000)
	fy := math.Floor(float64(ay) / 1_000_000)
	fz := math.Floor(float64(az) / 1_000_000)
	x = int(math.Min(3, math.Max(-3, fx)) + 3.0)
	y = int(math.Min(3, math.Max(-3, fy)) + 3.0)
	// Intensity range is 0x0...0xf. We treat 0G as 0x8
	i = byte(math.Floor((math.Min(3, math.Max(-3, fz)) + 3.0) * 16.0 / 7.0))
	return
}

func mapRotationToRGB(rx, ry, rz int32) [3]uint32 { // [r,g,b]
	// We'll use a scale factor with a max of 360 degrees per second, mapped to PWM Top
	const maxRot float64 = 360_000_000.0
	rot := [3]int32{rx, ry, rz}
	ret := [3](uint32){}
	for i, v := range rot {
		ret[i] = uint32(math.Min(maxRot, math.Abs(float64(v))) / maxRot * float64(pwmsRGB[i].Top()))
	}
	return ret
}

func main() {
	setupPins()
	setDisplayIntensity(8)
	for {
		clearDisplay()
		// Read Acceleration in µg (micro-gravity)
		accX, accY, accZ := imu.ReadAcceleration()
		// Map acceleration to LED display pos/intensity
		posX, posY, intensity := mapAccelerationToPositionAndIntensity(accX, accY, accZ)
		drawBall(posX, posY, intensity)
		// Read Rotation in µ°/s (micro-degrees/sec)
		rotX, rotY, rotZ := imu.ReadRotation()
		// Map rotation to RGB component intensity
		rgb := mapRotationToRGB(rotX, rotY, rotZ)
		// Set RGB LED
		for i, v := range rgb {
			ch, _ := pwmsRGB[i].Channel(pinsRGB[i])
			pwmsRGB[i].Set(ch, v)
		}
		// Wait and repeat
		time.Sleep(time.Millisecond * 10)
	}
}
