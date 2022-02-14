# Lesson 11: Membrane Switch Module #

This example demonstrates how to read a keypress from a membrane keypad
like [this one](https://cdn.sparkfun.com/assets/f/f/a/5/0/DS-16038.pdf)
supplied with the elegoo kit.

The [TinyGo Drivers](https://pkg.go.dev/tinygo.org/x/drivers) project already
includes a driver for this keypad so we'll use that.

If you're curious to know how this works internally, check out the
[datasheet]((https://cdn.sparkfun.com/assets/f/f/a/5/0/DS-16038.pdf)) and the
[driver code](https://github.com/tinygo-org/drivers/blob/release/keypad4x4/keypad4x4.go).

## Keypad Config ##

Connect the keypad to the Pico as follows. Here I'm using the colours from
the picture in the datasheet, but obviously they are not important.

| Colour | Keypad Matrix | GPIO Pin (physical pin#) |
|-|-|-|
| Orange | Row 1 | GP26 (31) |
| Green | Row 2 | GP22 (29) |
| Purple | Row 3 | GP21 (27) |
| Blue| Row 4 | GP20 (26) |
| Yellow | Col 1 | GP19 (25) |
| Red | Col 2 | GP18 (24) |
| Turquioise | Col 3 | GP17 (22) |
| Pink | Col 4 | GP16 (21) |

## Feedback ##

Like lesson 10 (ultrasonic sensor module), the original example code for the Arduino Uno
used the serial port to print out which key was pressed. Since we don't automatically
have this feature out-of-the-box with the Pico, we'll use a passive buzzer as in lesson 7
to play a different tone for each key press.

Connect the passive buzzer +ve temrinal to GP12 (physical pin #16), and the -ve terminal
to Ground (physical pins 3,8,13,18,23,28,33,38)
