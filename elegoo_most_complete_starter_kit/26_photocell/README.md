# Lesson 26: Photocell #

This example builds on lesson 24 (Eight LEDs with 74HC595) by adding a
'[photocell](https://en.wikipedia.org/wiki/Photoresistor)', also known as a Light Dependent Resistor (LDR).
The LDR varies its resistance based on the amount of light falling on it.

We combine the LDR with a 1KΩ resistor to form a
[voltage divider](https://learn.sparkfun.com/tutorials/voltage-dividers/all) between Ground and 3.3V.
Together, these work like a potentiometer that we can read using the Pico's ADC.

We vary the number of LEDs lit according to the ADC signal based on the light falling on the LDR.

## Driver Support ##

In contrast to lesson 24, this time we will use the
[driver](https://github.com/tinygo-org/drivers/blob/release/shiftregister/shiftregister.go)
available from the [TinyGo drivers](https://github.com/tinygo-org/drivers) project to simplify
our code.

## Connections ##

The connections are the same as for lesson 24, with the addition on the LDR/1KΩ voltage divider on ADC0.

Note: For brevity of this table I have not split the Resistors R0-R7 into separate columns
for the positive (+) and negative (-) terminals. The polarity doesn't matter anyway. Just be
aware that the resistor *Rn* is **always** connected in ***series*** **between** the 74HC595 pin
*Qn* and the LED *Ln*. (i.e) do **not connect in *parallel* directly beween *Qn* and *Ln***.

| Pico (pin #) | 74HC595 (pin #) | 220Ω Resistors Rn | LEDs Ln | LDR | 1KΩ
|-|-|-|-|-|-|
| GP11 (15) | SH_CP (11) | | |
| GP12 (16) | ST_CP (12) | | |
| GP13 (17) | DS (14) | | |
| ADC0 (31) | | | | - | + |
| 3V3 (OUT) (36) | VCC (16) **and** MR (10) | | | + | |
| | Q0 (15) | R0 | L0 (+) |
| | Q1 (1)  | R1 | L1 (+) |
| | Q2 (2)  | R2 | L2 (+) |
| | Q3 (3)  | R3 | L3 (+) |
| | Q4 (4)  | R4 | L4 (+) |
| | Q5 (5)  | R5 | L5 (+) |
| | Q6 (6)  | R6 | L6 (+) |
| | Q7 (7)  | R7 | L7 (+) |
| Ground (3,8,13,18,23,28,33,38) | GND (8) **and** OE (13 | | L0-L7 (-) | | - |
||||
