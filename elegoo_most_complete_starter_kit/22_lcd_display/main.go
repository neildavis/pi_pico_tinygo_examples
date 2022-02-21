package main

import (
	"fmt"
	"machine"
	"time"

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

func setupPins() {
	lcd, _ = hd44780.NewGPIO4Bit(pinLCDData[:], pinLCDEN, pinLCDRS, pinLCDRW)
	lcd.Configure(hd44780.Config{Width: lcdCols, Height: lcdRows})
}

func main() {
	setupPins()
	n := 1
	// Let's play a variant of fizzbuzz (https://en.wikipedia.org/wiki/Fizz_buzz)
	// The LCDs top line will display the number always.
	// The LCDs bottom line will show 'Fizz', 'Buzz' or 'Fizz Buzz' as appropriate.
	for {
		lcd.ClearDisplay()
		fizz := n%3 == 0
		buzz := n%5 == 0
		strTop := fmt.Sprintf("%v", n)
		strBot := ""
		if fizz && buzz {
			strBot = "Fizz Buzz"
		} else if fizz {
			strBot = "Fizz"
		} else if buzz {
			strBot = "Buzz"
		}
		n++
		lcd.SetCursor(0, 0)
		lcd.Write([]byte(strTop))
		lcd.Display()
		lcd.SetCursor(0, 1)
		lcd.Write([]byte(strBot))
		lcd.Display()
		time.Sleep(time.Second)
	}
}
