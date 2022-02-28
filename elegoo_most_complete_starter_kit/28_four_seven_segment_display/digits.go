package main

const (
	// Individual Segments that can be OR'ed together to form recognizable characters
	SEG_A   = 0b10000000
	SEG_B   = 0b01000000
	SEG_C   = 0b00100000
	SEG_D   = 0b00010000
	SEG_E   = 0b00001000
	SEG_F   = 0b00000100
	SEG_G   = 0b00000010
	SEG_DP  = 0b00000001
	SEG_ALL = 0b11111111
	// Special chars
	CLEAR     = 0
	FULL_STOP = SEG_DP
	DASH      = SEG_G
	// Digits 0-9
	DIG_0 = SEG_A | SEG_B | SEG_C | SEG_D | SEG_E | SEG_F
	DIG_1 = SEG_B | SEG_C
	DIG_2 = SEG_A | SEG_B | SEG_D | SEG_E | SEG_G
	DIG_3 = SEG_A | SEG_B | SEG_C | SEG_D | SEG_G
	DIG_4 = SEG_B | SEG_C | SEG_F | SEG_G
	DIG_5 = SEG_A | SEG_C | SEG_D | SEG_F | SEG_G
	DIG_6 = SEG_A | SEG_C | SEG_D | SEG_E | SEG_F | SEG_G
	DIG_7 = SEG_A | SEG_B | SEG_C
	DIG_8 = SEG_A | SEG_B | SEG_C | SEG_D | SEG_E | SEG_F | SEG_G
	DIG_9 = SEG_A | SEG_B | SEG_C | SEG_D | SEG_F | SEG_G
	// Letters - not all letters are possible on 7-seg
	LET_A = SEG_A | SEG_B | SEG_C | SEG_E | SEG_F | SEG_G
	LET_b = SEG_C | SEG_D | SEG_E | SEG_F | SEG_G
	LET_C = SEG_A | SEG_D | SEG_E | SEG_F
	LET_d = SEG_B | SEG_C | SEG_D | SEG_E | SEG_G
	LET_E = SEG_A | SEG_D | SEG_E | SEG_F | SEG_G
	LET_F = SEG_A | SEG_E | SEG_F | SEG_G
	LET_g = DIG_9
	LET_H = SEG_B | SEG_C | SEG_E | SEG_F | SEG_G
	LET_I = DIG_1
	LET_j = SEG_B | SEG_C | SEG_D
	LET_K = SEG_ALL // K/k is not possible
	LET_L = SEG_D | SEG_E | SEG_F
	LET_M = SEG_ALL // M/m is not possible
	LET_N = SEG_A | SEG_B | SEG_C | SEG_E | SEG_F
	LET_O = DIG_0
	LET_P = SEG_A | SEG_B | SEG_E | SEG_F | SEG_G
	LET_q = SEG_A | SEG_B | SEG_C | SEG_F | SEG_G
	LET_r = SEG_A | SEG_E | SEG_F
	LET_S = DIG_5
	LET_t = SEG_D | SEG_E | SEG_F | SEG_G
	LET_U = SEG_B | SEG_C | SEG_D | SEG_E | SEG_F
	LET_V = SEG_ALL // V/v is not possible
	LET_W = SEG_ALL // W/w is not possible
	LET_X = SEG_ALL // X/x is not possible
	LET_Y = SEG_B | SEG_C | SEG_D | SEG_F | SEG_G
	LET_Z = DIG_2
)
