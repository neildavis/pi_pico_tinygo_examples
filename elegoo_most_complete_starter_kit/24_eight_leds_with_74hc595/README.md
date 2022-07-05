# Lesson 24: Eight LEDs with 74HC595 #

This example demonstrates how to drive eight digital LEDs without the need to use eight output pins on the Pico.
Instead, it uses the [SN74HC595N](https://www.ti.com/lit/ds/symlink/sn74hc595.pdf)
[shift register](https://en.wikipedia.org/wiki/Shift_register) IC supplied with the Elegoo kit.

The SN74HC595N is an example of a
[Serial In Parallel Out (SIPO)](https://en.wikipedia.org/wiki/Shift_register#Serial-in_parallel-out_(SIPO))
shift register. Data bits are streamed in serially (bit-by-bit) using the 'data' and 'clock' pins
and shifted down (cascaded) from the most to least significant bits (MSB->LSB) in its internal register.
When the 'latch' pin is triggered, the internal register state is transferred onto the parallel output pins.
A '1' results in a HIGH signal, wherreas a '0' results in a LOW signal.

All this means we only need to use *three* of our Pico's pins as output to drive *eight* digital outputs.

## Driver Support ##

Although the [TinyGo drivers](https://github.com/tinygo-org/drivers) project includes
a [driver](https://github.com/tinygo-org/drivers/blob/release/shiftregister/shiftregister.go)
for SIPO shift register ICs like the 74HC595, the operations to drive it are not
complex and in this example we will write the code to drive the chip ourselves rather than
use the driver.

Subsequent lessons also make use of a 74HC595, and in those examples we will make use
of the driver, to show examples of both approaches.

## Connections ##

Note: For brevity of this table I have not split the resistors R0-R7 into separate columns
for the positive (+) and negative (-) terminals. The polarity doesn't matter anyway. Just be
aware that the resistor *Rn* is **always** connected in ***series*** **between** the 74HC595 pin
*Qn* and the LED *Ln*. (i.e) do **not connect in *parallel* directly beween *Qn* and *Ln***.

| Pico (pin #) | 74HC595 (pin #) | 220Ω Resistors Rn | LEDs Ln |
|-|-|-|-|
| GP11 (15) | SH_CP (11) | | |
| GP12 (16) | ST_CP (12) | | |
| GP13 (17) | DS (14) | | |
| 3V3 (OUT) (36) | VCC (16) **and** MR (10) | | | |
| | Q0 (15) | R0 | L0 (+) |
| | Q1 (1)  | R1 | L1 (+) |
| | Q2 (2)  | R2 | L2 (+) |
| | Q3 (3)  | R3 | L3 (+) |
| | Q4 (4)  | R4 | L4 (+) |
| | Q5 (5)  | R5 | L5 (+) |
| | Q6 (6)  | R6 | L6 (+) |
| | Q7 (7)  | R7 | L7 (+) |
| Ground (3,8,13,18,23,28,33,38) | GND (8) **and** OE (13 | | L0-L7 (-) |
||||
