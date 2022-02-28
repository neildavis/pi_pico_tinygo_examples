# Lesson 27: 74HC595 And Segment Display #

This example builds on lesson 24 (Eight LEDs with 74HC595) by using a 'segment display' in place of
the eight LEDs to display the numbers 9-0 and a decimal point.

## Driver Support ##

Like lesson 26, and in contrast to lesson 24, we will use the shift register
[driver](https://github.com/tinygo-org/drivers/blob/release/shiftregister/shiftregister.go)
available from the [TinyGo drivers](https://github.com/tinygo-org/drivers) project to simplify
our code.

## Connections ##

The connections are similar to lesson 24, only we're using a 'seven segment' display instead of
the eight LEDs. A seven segment display actually contains *eight* LEDs since it also includes a
decimal point (DP) in addition to the seven segments (A-G) making up the digit.

We can also take advantage of the 'common cathode' nature of the seven segment display to
reduce the number of 220Ω current limiting resistors required. There are two ground connection
pins on the seven segment display, each of which are shared with all eight of the LED cathodes.
Therefore, we only need two 220Ω current limiting resistors instead of the eight we used before.
One resistor each will connect between the seven seg ground pins and the Pico's ground.

| Pico (pin #) | 74HC595 (pin #) | Segment Display | R1 220Ω | R2 220Ω |
|-|-|-|-|-|
| GP11 (15) | SH_CP (11) | | | |
| GP12 (16) | ST_CP (12) | | | |
| GP13 (17) | DS (14) | | | |
| 3V3 (OUT) (36) | VCC (16) **and** MR (10) | | | |
| | Q0 (15) | A | | | |
| | Q1 (1)  | B | | | |
| | Q2 (2)  | C | | | |
| | Q3 (3)  | D | | | |
| | Q4 (4)  | E | | | |
| | Q5 (5)  | F | | | |
| | Q6 (6)  | G | | | |
| | Q7 (7)  | DP | | | |
| | | GND1 | + | |
| | | GND2 | | + |
| Ground (3,8,13,18,23,28,33,38) | GND (8) **and** OE (13) | | - | - |
