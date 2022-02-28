package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/shiftregister"
)

const (
	MSG_REPEAT_DELAY   = time.Second * 4 // Delay before repeating the message
	CHAR_ADVANCE_DELAY = time.Second / 2 // Delay between character scrolls
	MESSAGE            = "Hello TinyGo"  // Message to dsiplay
)

var (
	// Digit select pins
	pinDigits = []machine.Pin{machine.GP6, machine.GP7, machine.GP8, machine.GP9}
	// 74HC595 Shift Register
	pinClock       = machine.GP11 // SH_CP [SCK] on 74HC595
	pinLatch       = machine.GP12 // ST_CP [RCK] on 74HC595
	pinData        = machine.GP13 // DS [S1] on 74HC595
	shiftRegDevice shiftregister.Device
	// A channel we will use for the display loop
	ch chan []byte
	// Lookup tables for converting numbers & letters to 7-seg chars
	DIGITS = []byte{
		DIG_0, DIG_1, DIG_2, DIG_3, DIG_4,
		DIG_5, DIG_6, DIG_7, DIG_8, DIG_9}
	LETTERS = []byte{
		LET_A, LET_b, LET_C, LET_d, LET_E, LET_F,
		LET_g, LET_H, LET_I, LET_j, LET_K, LET_L,
		LET_M, LET_N, LET_O, LET_P, LET_q, LET_r,
		LET_S, LET_t, LET_U, LET_V, LET_W, LET_X,
		LET_Y, LET_Z,
	}
)

func setupPins() {
	// Digit select pins
	for _, pin := range pinDigits {
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		pin.High() // Deselect all digits to begin. We select a digit by lowering it to ground
	}
	// 74HC595 Shift Register
	shiftRegDevice = *shiftregister.New(shiftregister.EIGHT_BITS, pinLatch, pinClock, pinData)
	shiftRegDevice.Configure()
}

// Goroutine to continuously update the 4x 7-seg display
func fourSegDisplayLoop() {
	// Initial digit values
	digits := []byte{CLEAR, CLEAR, CLEAR, CLEAR}
	// We continually cycle over each digit
	for {
		for i := 0; i < len(digits); i++ {
			// Shift the digit in
			shiftRegDevice.WriteMask(uint32(digits[i]))
			// Select the digit by raising its common cathode pin to ground, allowing current to flow
			pinDigits[i].Low()
			// We need to wait a short amount of time to allow the digit to illuminate
			time.Sleep(time.Microsecond * 10)
			// Deselect the LAST digit for update by setting its common cathode pin to VCC, preventing current from flowing.
			pinDigits[i].High()
		}
		// See if we have updated digits from the channel
		if len(ch) > 0 {
			// We have a new set of digits waiting. Pull them from the channel and update
			newDigits := <-ch
			copy(digits, newDigits[:4])
		}
	}
}

func charTo7Seg(c byte) byte {
	mask := byte(0)
	switch {
	case c == ' ':
		// Space
		mask = CLEAR
	case c == '-':
		// Dash
		mask = DASH
	case c >= '0' && c <= '9':
		// Numeric digit
		mask = DIGITS[c-'0']
	case c >= 'A' && c <= 'Z':
		// Upper case letter
		mask = LETTERS[c-'A']
	case c >= 'a' && c <= 'z':
		// Lower case letter
		mask = LETTERS[c-'a']
	default:
		mask = byte(SEG_ALL) // undefined char
	}
	return mask
}

func displayText(text string) {
	// pad out the message so it clears the display
	bytes := []byte(text + "    ")
	// Start with an empty display and scroll chars in one-by-one from the right
	digits := [4]byte{CLEAR, CLEAR, CLEAR, CLEAR}
	for _, c := range bytes {
		// Bump each char to left
		digits[0] = digits[1]
		digits[1] = digits[2]
		digits[2] = digits[3]
		// Insert new char on RHS (digit 4)
		digits[3] = charTo7Seg(c)
		// Update display by sending new digits to the channel
		ch <- digits[:]
		// Wait before displaying next char
		time.Sleep(CHAR_ADVANCE_DELAY)
	}
}

func main() {
	setupPins()
	// Configure and start the display loop goroutine
	ch = make(chan []byte, 1)
	go fourSegDisplayLoop()
	// Display our message forever
	for {
		displayText(MESSAGE)
		time.Sleep(MSG_REPEAT_DELAY)
	}
}
