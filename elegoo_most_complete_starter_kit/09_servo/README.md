# Lesson 9: Servo #

This example demonstrates the use of a servo.

A servo is a type of geared motor that can only rotate a maximum of 180 degrees.
It can be controlled by sending electrical pulses (PWM signals) from the Pi Pico
to tell the servo what position it should move to.

Note: The SG90 servo supplied with the elegoo kit is a 4.8V device and was supplied for
use with the Arduino which uses 5V signalling, whereas the Pi Pico uses 3.3V.
To avoid use of level shifting circuitry, we can power the servo from the Pico's
5V VBUS when powered by USB, or an external 5V supply perhaps using the power supply
module included with the kit. The PWM signalling should still work at 3.3V as long
as the ground is shared between the 3.3V and 5.5V supplies (it will be for the VBUS case)

According to the [SG90 datasheet](http://www.ee.ic.ac.uk/pcheung/teaching/DE1_EE/stores/sg90_datasheet.pdf)
the PWM signal should be within the following range:

| Pulse duration (ms) | Servo Position |
|-|-|
|1| Full Left|
|1.5| Middle|
|2| Full Right|

The [TinyGo Drivers](https://pkg.go.dev/tinygo.org/x/drivers) project already includes a driver for servos, so we'll
use that.

If you're curious exactly how this works, inspect the [driver code](https://github.com/tinygo-org/drivers/blob/release/servo/servo.go).


The servo has three wires and should be connected as follows:

| Wire Colour | Purpose | Pin | Physical Pin #
|-|-|-|-|
| Brown | Ground | Ground | 3,8,13,18,23,28,33,35,38 |
| Red | Vcc | 5v VBUS | 40|
| Orange | Signal | GP 1 | 2|
