package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/hd44780"
	"tinygo.org/x/drivers/irremote"
)

var (
	// LCD
	pinLCDEN   = machine.GP16
	pinLCDRS   = machine.GP17
	pinLCDData = [4]machine.Pin{machine.GP18, machine.GP19, machine.GP20, machine.GP21}
	pinLCDRW   = machine.GP22
	lcd        hd44780.Device
	// LED
	pinLED = machine.LED
	// IR Receiver
	pinIRIn = machine.GP3
	ir      irremote.ReceiverDevice
	// A channel of IR commands
	ch chan irremote.Data
)

const (
	lcdCols = 16
	lcdRows = 2
)

func setupPins() {
	// LCD
	lcd, _ = hd44780.NewGPIO4Bit(pinLCDData[:], pinLCDEN, pinLCDRS, pinLCDRW)
	lcd.Configure(hd44780.Config{Width: lcdCols, Height: lcdRows})
	// LED
	pinLED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// Setup IR receiver
	ir = irremote.NewReceiver(pinIRIn)
	ir.Configure()
}

// Handle a callback from the IR receiver
func irCallback(data irremote.Data) {
	ch <- data
}

// Blink the LED
func blinkLED() {
	pinLED.High()
	time.Sleep(time.Millisecond * 10)
	pinLED.Low()
}

// Helper to output some text on the LCD
func lcdDisplayText(strTop, strBot string) {
	lcd.SetCursor(0, 0)
	lcd.Write([]byte(strTop))
	lcd.Display()
	lcd.SetCursor(0, 1)
	lcd.Write([]byte(strBot))
	lcd.Display()
	fmt.Println("LCD:", strTop, strBot, "\r")
}

func main() {
	setupPins()
	lcd.ClearDisplay()
	// Create a buffered channel of notes to play
	ch = make(chan irremote.Data, 100)
	// Register for IR callbacks
	ir.SetCommandHandler(irCallback)
	lcdDisplayText("Wait for NEC IR", fmt.Sprintf("on Pin %d", pinIRIn))
	for {
		// Read data from the channel
		data := <-ch
		// Blink the LED as feedback that we received something
		blinkLED()
		// Display IR data on LCD
		lcd.ClearDisplay()
		strRepeat := ""
		if data.Flags&irremote.DataFlagIsRepeat != 0 {
			strRepeat = "RPT"
		}
		strTop := fmt.Sprintf("Code: %08X", data.Code)
		strBot := fmt.Sprintf("A: %04X C:%02X %s", data.Address, data.Command, strRepeat)
		lcdDisplayText(strTop, strBot)
	}
}
