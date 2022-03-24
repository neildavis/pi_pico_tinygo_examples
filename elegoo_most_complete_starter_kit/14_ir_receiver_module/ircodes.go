package main

import "github.com/neildavis/tinygo_modules/passivebuzzer"

var irCmdButtons = map[uint16]passivebuzzer.Note{
	0xA2: passivebuzzer.NOTE_C4,
	0xE2: passivebuzzer.NOTE_D5,
	0x62: passivebuzzer.NOTE_E4,
	0x22: passivebuzzer.NOTE_F4,
	0x02: passivebuzzer.NOTE_G4,
	0xC2: passivebuzzer.NOTE_A4,
	0xE0: passivebuzzer.NOTE_B4,
	0xA8: passivebuzzer.NOTE_C5,
	0x90: passivebuzzer.NOTE_D5,
	0x98: passivebuzzer.NOTE_E5,
	0xB0: passivebuzzer.NOTE_F5,
	0x68: passivebuzzer.NOTE_G5,
	0x30: passivebuzzer.NOTE_A5,
	0x18: passivebuzzer.NOTE_B5,
	0x7A: passivebuzzer.NOTE_C6,
	0x10: passivebuzzer.NOTE_D6,
	0x38: passivebuzzer.NOTE_E6,
	0x5A: passivebuzzer.NOTE_F6,
	0x42: passivebuzzer.NOTE_G6,
	0x4A: passivebuzzer.NOTE_A6,
	0x52: passivebuzzer.NOTE_B6,
}
