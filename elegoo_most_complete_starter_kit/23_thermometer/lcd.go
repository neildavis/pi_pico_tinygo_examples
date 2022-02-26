package main

import (
	"machine"

	"tinygo.org/x/drivers/hd44780"
)

var (
	pinLCDEN   = machine.GP16
	pinLCDRS   = machine.GP17
	pinLCDData = [4]machine.Pin{machine.GP18, machine.GP19, machine.GP20, machine.GP21}
	pinLCDRW   = machine.GP22
	lcd        hd44780.Device
)

const (
	lcdCols = 16
	lcdRows = 2
)

func lcdSetup() {
	lcd, _ = hd44780.NewGPIO4Bit(pinLCDData[:], pinLCDEN, pinLCDRS, pinLCDRW)
	lcd.Configure(hd44780.Config{Width: lcdCols, Height: lcdRows})
}

func lcdClear() {
	lcd.ClearDisplay()
}

func lcdWriteLine(line string) {
	lcdClear()
	lcd.SetCursor(0, 0)
	lcd.Write([]byte(line))
	lcd.Display()
}

func lcdWriteLines(line1, line2 string) {
	lcdClear()
	lcd.SetCursor(0, 0)
	lcd.Write([]byte(line1))
	lcd.Display()
	lcd.SetCursor(0, 1)
	lcd.Write([]byte(line2))
	lcd.Display()
}
