package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/max72xx"
)

var (
	pinDin = machine.SPI1_SDO_PIN // SPI1 COPI/SDO/TX
	pinCS  = machine.GP9          // SPI1 CSn - machine package does not define a default constant
	pinCLK = machine.SPI1_SCK_PIN // SPI1 SCK
	spi    = machine.SPI1
	matrix *max72xx.Device
	// We keep track of the led-status in this array
	status = [NUM_ROWS]byte{}
)

const (
	DELAY         = time.Millisecond * 50
	SPI_CPOL      = 1        // SPI Clock Polarity: CLK is HIGH when idle
	SPI_CPHA      = 0        // SPI Clock Phase: Data is read on 'falling edge', which is first edge when CPOL is 1
	SPI_FREQ      = 10000000 // 10 MHz is max frequency supportewd by MAX72xx
	SPI_DATA_BITS = 8        // 8-bit bus
	// Our LED Matrix is 8x8
	NUM_ROWS = 8
	NUM_COLS = 8
)

func setupPins() {
	spiMode := uint8((SPI_CPOL << 1) | SPI_CPHA)
	spiConfig := machine.SPIConfig{Frequency: SPI_FREQ, LSBFirst: false, Mode: spiMode, DataBits: SPI_DATA_BITS,
		SDO: pinDin, SCK: pinCLK}
	spi.Configure(spiConfig)
	matrix = max72xx.NewDevice(*spi, pinCS)
	matrix.Configure()
	// Configure the MAX72XX
	matrix.StopDisplayTest()
	matrix.SetScanLimit(8)
	matrix.SetDecodeMode(0)   // No decoding for matrix mode (used for 7-seg digits)
	matrix.StopShutdownMode() // The MAX72XX is in power-saving mode on startup
	setDisplayIntensity(8)
}

func setDisplayIntensity(intensity uint8) {
	matrix.WriteCommand(max72xx.REG_INTENSITY, intensity)
}

func clearDisplay() {
	for row := 0; row < NUM_ROWS; row++ {
		setRow(row, 0)
	}
}

func setLed(row, col int, state bool) {
	if row < 0 || row > NUM_ROWS-1 || col < 0 || col > NUM_COLS-1 {
		return
	}
	val := byte(0b10000000 >> col)
	if state {
		status[row] |= val
	} else {
		status[row] &^= val
	}
	setRow(row, status[row])
}

func setRow(row int, value byte) {
	if row < 0 || row > NUM_ROWS-1 {
		return
	}
	status[row] = value
	matrix.WriteCommand(byte(row+1), value)
}

func setColumn(col int, value byte) {
	if col < 0 || col > NUM_COLS-1 {
		return
	}
	for row := 0; row < NUM_ROWS; row++ {
		val := (value >> (NUM_ROWS - row - 1)) & 0x1
		setLed(row, col, val != 0)
	}
}

/*
 This method will display the characters for the
 word "TinyGo" one after the other on the matrix.
*/
func displayTinyGo() {
	// These are the data for the chars
	T := []byte{
		0b00000000,
		0b10000000,
		0b10000000,
		0b11111111,
		0b10000000,
		0b10000000,
		0b00000000,
	}
	i := []byte{
		0b00000000,
		0b00000000,
		0b00000001,
		0b10111111,
		0b00100001,
		0b00000000,
		0b00000000,
	}
	n := []byte{
		0b00000000,
		0b00011111,
		0b00100000,
		0b00100000,
		0b00010000,
		0b00111111,
		0b00000000,
	}
	y := []byte{
		0b00000000,
		0b11111110,
		0b00100001,
		0b00010001,
		0b00010001,
		0b11100010,
		0b00000000,
	}
	G := []byte{
		0b00000000,
		0b01001110,
		0b10001001,
		0b10000001,
		0b10000001,
		0b01111110,
		0b00000000,
	}
	o := []byte{
		0b00000000,
		0b00011100,
		0b00100010,
		0b00100010,
		0b00100010,
		0b00011100,
		0b00000000,
	}
	// Now display them one-by-one with a delay inbetween each char
	// T
	for row, val := range T {
		setRow(row, val)
	}
	time.Sleep(DELAY * 10)
	// i
	for row, val := range i {
		setRow(row, val)
	}
	time.Sleep(DELAY * 10)
	// n
	for row, val := range n {
		setRow(row, val)
	}
	time.Sleep(DELAY * 10)
	// y
	for row, val := range y {
		setRow(row, val)
	}
	time.Sleep(DELAY * 10)
	// G
	for row, val := range G {
		setRow(row, val)
	}
	time.Sleep(DELAY * 10)
	// o
	for row, val := range o {
		setRow(row, val)
	}
	time.Sleep(DELAY * 10)
}

/*
 This function lights up a some Leds in a row.
 The pattern will be repeated on every row.
 The pattern will blink along with the row-number.
 row number 4 (index==3) will blink 4 times etc.
*/
func displayRowsPattern() {
	for row := 0; row < NUM_ROWS; row++ {
		time.Sleep(DELAY)
		setRow(row, 0b10100000)
		time.Sleep(DELAY)
		setRow(row, 0)
		for i := 0; i < row; i++ {
			time.Sleep(DELAY)
			setRow(row, 0b10100000)
			time.Sleep(DELAY)
			setRow(row, 0)
		}
	}
}

/*
 This function lights up a some Leds in a column.
 The pattern will be repeated on every column.
 The pattern will blink along with the column-number.
 column number 4 (index==3) will blink 4 times etc.
*/
func displayColumnsPattern() {
	for col := 0; col < NUM_COLS; col++ {
		time.Sleep(DELAY)
		setColumn(col, 0b10100000)
		time.Sleep(DELAY)
		setColumn(col, 0)
		for i := 0; i < col; i++ {
			time.Sleep(DELAY)
			setColumn(col, 0b10100000)
			time.Sleep(DELAY)
			setColumn(col, 0)
		}
	}
}

/*
 This function will light up every Led on the matrix.
 The led will blink along with the row-number.
 row number 4 (index==3) will blink 4 times etc.
*/
func displaySinglePattern() {
	for row := 0; row < NUM_ROWS; row++ {
		for col := 0; col < NUM_COLS; col++ {
			time.Sleep(DELAY)
			setLed(row, col, true)
			time.Sleep(DELAY)
			for i := 0; i < col; i++ {
				setLed(row, col, false)
				time.Sleep(DELAY)
				setLed(row, col, true)
				time.Sleep(DELAY)
			}
		}
	}
}

// Entry point
func main() {
	// Configure all connections and SPI bus
	setupPins()
	for {
		clearDisplay()
		displayTinyGo()
		clearDisplay()
		displayRowsPattern()
		clearDisplay()
		displayColumnsPattern()
		clearDisplay()
		displaySinglePattern()
	}
}
