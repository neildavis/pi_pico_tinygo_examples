# Lesson 28: Four Digital Seven Segment Display #

This example expands upon the previous lesson to utilize a
[*5461AS*](http://www.datasheetcafe.com/kyx-5461as-datasheet-7-segment-arduino/) 
'four seven segment' display to show some scrolling text.

## Common Cathodes ##

In the previous lesson (27: '74HC595 And Segment Display') we noted that the eight LEDs
making up the segment display shared a common cathode, exposed by pins GND1 & GND2,
which we connected to the Pico's GROUND via protective current limiting 220Ω Resistors.

With the *5461AS* things work a little differently:

* The '*SEG N*' pins work just like the single segment display, only now they apply to
**all four** of the digits.
* There are **four** common cathodes, one for each digit. These act as a 'digit selector'.
  Only one should be at GROUND at any particular time, and this corresponds to the active digit.
  Thus, only a single segment digit is ever powered/lit at any particular time.
* Therefore, rather than connecting the common cathodes directly to ground
  (via our current limiting 220Ω Resistors of course), we instead connect them to four
  digital pins of the Pico. Three are held HIGH to prevent current flowing, whilst the
  fourth is lowered to GROUND to select the active digit.
* The four simultaneous digit display is achieved through the principle of
  ['Persistence of Vision'](https://en.wikipedia.org/wiki/Persistence_of_vision).
  The four digits are continuously cycled rapidly in turn, one at a time, too fast for
  the human eye to detect, giving the optical illusion that all four digits are lit.

## Connections ##

Note: For brevity of this table I have not split the 220Ω resistors R1-R3 into separate columns
for the positive (+) and negative (-) terminals. The polarity doesn't matter anyway.
Just be aware that the resistor *Rn* is **always** connected in ***series*** **between** the 
*5461AS* (4x &-Seg display) pins *'DIG n'* and the Pico.
(i.e) **do not connect in *parallel* directly beween *5461AS* and the *Pico***.

| Pico (pin #) | 220Ω Resistors Rn | 74HC595 (pin #) | 4x 7-Seg Display (5461AS - pin #) | 
|-|-|-|-|
| GP6 (9) | R1 || DIG 1 (12) |
| GP7 (10)| R2 || DIG 2 (9) |
| GP8 (11) | R3 || DIG 3 (8) |
| GP9(12) | R4 || DIG 4 (6) |
| GP11 (15) |  | SH_CP (11) | |
| GP12 (16) |  | ST_CP (12) | |
| GP13 (17) |  | DS (14) | |
| 3V3 (OUT) (36) | | VCC (16) **and** MR (10) | | |
| |  | Q0 (15) | SEG A (11) |
| |  | Q1 (1)  | SEG B (7) |
| |  | Q2 (2)  | SEG C (4) |
| |  | Q3 (3)  | SEG D (2) |
| |  | Q4 (4)  | SEG E (1) |
| |  | Q5 (5)  | SEG F (10) |
| |  | Q6 (6)  | SEG G (5) |
| |  | Q7 (7)  | DP (3) |
| Ground (3,8,13,18,23,28,33,38) | | GND (8) **and** OE (13 | |
||||
