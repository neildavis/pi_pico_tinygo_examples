# Lesson 13: Analog Joystick Module #

This example demonstrates use of the Pico's 'Analog-to-Digital Converter' (ADC) inputs to read the position of an analog joystick.

## Feedback ##

Again, the Elegoo kit used the Arduino serial bus and IDE monitor to simply print out the state of
the joystick inputs. This example uses an RGB LED. The direction of the joystick will determine the
colour of the LED, like a simple [colour wheel](https://en.wikipedia.org/wiki/Color_wheel).
How far the joystick is pushed determines the intensity of the LED. Pressing the stick down will
turn off the LED whichever position it is in, until released.

## Connections ##

| Pin (physical pin#) | Joystick | RGB LED |
|-|-|-|
| GP0 (1) | | Red |
| GP1 (2) | | Green |
| GP2 (4) | | Blue |
| GP22 (29) | SW(itch) |
| ADC0 (31) | X/Horiz |
| ADC1 (32) | Y/Vert |
| 3V3(OUT) (36) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND | Cathode |
