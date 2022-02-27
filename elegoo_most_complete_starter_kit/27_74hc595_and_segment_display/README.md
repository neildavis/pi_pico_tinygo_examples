# Lesson 27: 74HC595 And Segment Display #

This example builds on lesson 24 (Eight LEDs with 74HC595) by using a 'segment display' in place of
the eight LEDs to display the numbers 9-0 and a decimal point.

## Driver Support ##

Like lesson 26, and in contrast to lesson 24, we will use the shift register
[driver](https://github.com/tinygo-org/drivers/blob/release/shiftregister/shiftregister.go)
available from the [TinyGo drivers](https://github.com/tinygo-org/drivers) project to simplify
our code.

## Connections ##

The connections are similar to lesson 24, however the 'ground facing' side of the resistors
connect to the pins of the segment display as shown below instead of the LED pins.

Note: For brevity of this table I have not split the Resistors R0-R7 into separate columns
for the positive (+) and negative (-) terminals. The polarity doesn't matter anyway. Just be
aware that the resistor *Rn* is **always** connected in ***series*** **between** the 74HC595 pin
*Qn* and the segment display. (i.e) do **not connect in *parallel* directly beween *Qn* and the segment display**.

| Pico (pin #) | 74HC595 (pin #) | 220Ω Resistors Rn | Segment Display |
|-|-|-|-|
| GP11 (15) | SH_CP (11) | | |
| GP12 (16) | ST_CP (12) | | |
| GP13 (17) | DS (14) | | |
| ADC0 (31) | | | | - | + |
| 3V3 (OUT) (36) | VCC (16) **and** MR (10) | | |
| | Q0 (15) | R0 | A |
| | Q1 (1)  | R1 | B |
| | Q2 (2)  | R2 | C |
| | Q3 (3)  | R3 | D |
| | Q4 (4)  | R4 | E |
| | Q5 (5)  | R5 | F |
| | Q6 (6)  | R6 | G |
| | Q7 (7)  | R7 | DP |
| Ground (3,8,13,18,23,28,33,38) | GND (8) **and** OE (13 | | GND1 (3) **and** GND2 (8) |
||||
