package main

import (
	"machine"

	"tinygo.org/x/drivers/max72xx"
)

const (
	// Our LED Matrix is 8x8
	LED_NUM_ROWS = 8
	LED_NUM_COLS = 8
	// LED Matrix SPI config
	SPI_CPOL      = 1        // SPI Clock Polarity: CLK is HIGH when idle
	SPI_CPHA      = 0        // SPI Clock Phase: Data is read on 'falling edge', which is first edge when CPOL is 1
	SPI_FREQ      = 10000000 // 10 MHz is max frequency supportewd by MAX72xx
	SPI_DATA_BITS = 8        // 8-bit bus
)

var (
	// LED Matrix
	pinDin = machine.SPI1_SDO_PIN // SPI1 COPI/SDO/TX
	pinCS  = machine.GP9          // SPI1 CSn - machine package does not define a default constant
	pinCLK = machine.SPI1_SCK_PIN // SPI1 SCK
	spi    = machine.SPI1
	matrix *max72xx.Device
)

func setupLEDMatrix() {
	spiMode := uint8((SPI_CPOL << 1) | SPI_CPHA)
	spiConfig := machine.SPIConfig{Frequency: SPI_FREQ, LSBFirst: false, Mode: spiMode, DataBits: SPI_DATA_BITS,
		SDO: pinDin, SCK: pinCLK}
	spi.Configure(spiConfig)
	matrix = max72xx.NewDevice(*spi, pinCS)
	matrix.Configure()
	matrix.SetScanLimit(8)
	matrix.SetDecodeMode(0) // No decoding for matrix mode (used for 7-seg digits)
	matrix.SetIntensity(8)
	matrix.StopShutdownMode() // The MAX72XX is in power-saving mode on startup
}

func clearLED() {
	for row := 0; row < LED_NUM_ROWS; row++ {
		setLEDRow(row, 0)
	}
}

func setLEDRow(row int, value byte) {
	if row < 0 || row > LED_NUM_ROWS-1 {
		return
	}
	matrix.WriteCommand(byte(row+1), value)
}

func setLEDCheckPattern() {
	row0 := 0b01010101
	row1 := 0b10101010
	for i := 0; i < LED_NUM_ROWS; i++ {
		if i%2 == 0 {
			setLEDRow(i, byte(row0))
		} else {
			setLEDRow(i, byte(row1))
		}
	}
}
