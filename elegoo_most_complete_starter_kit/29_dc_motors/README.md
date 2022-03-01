# Lesson 29: DC Motors #

This example demonstrates how to use a [L293D](https://www.ti.com/lit/ds/symlink/l293d.pdf)
'*Quadruple Half-H Driver*' IC to drive a DC motor in either direction with variable speed.

## **IMPORTANT**: DC Motor Power Supply ##

You should **never** connect a motor directly to the Pico. When a motor is powered down it
creates an electrical feedback. With a small motor, like the one supplied with the kit, it
could damage the Pico. Larger motors could create sparks and/or fires.

In addition, DC motors are high inductive load devices and can draw more power than can be
provided from the Pico.

For these reasons, we always power the motor from an external power supply. The Elegoo kit
provides the [MB102](https://www.handsontec.com/dataspecs/mb102-ps.pdf) power supply module
for this purpose. The module itself can be powered in a couple of ways:

* An AC/DC mains transformer. The kit supplies such a transformer rated at 9V/1A DC output.
* A 9V battery. The kit supplies a battery and the necessary snap-on connector clip to interface
  with the MB102.

I recommend you use the AC mains adapter if possible.

Plug the MB102 into your breadboard, taking care with the orientation to ensure that its
positive and negative output terminals (marked +/- respectively) match up with the
positive and negative power rails on your breadboard (red+/blue- respectively).

Also ensure that the 'jumpers' on the MB102 are set correctly to power the rail you wish
to use for the motor at 5V (not 3.3V).

Finally, when using additional power supplies you should connect the Ground (-) sources
together as shown in the 'Connections' section below.

## Driver Support ##

The [TinyGo drivers](https://github.com/tinygo-org/drivers) project does provide a
[L293D driver](https://github.com/tinygo-org/drivers/blob/release/l293x/l293x.go).
However, the operation is very simple, consisting of two digital output pins to
control the motor direction, and one PWM output pin to control the speed.

For this example we will not use the driver and write the code ourselves.
Subsequent lessons also make use of the L293D, and we'll use the driver in those
examples to show and compare the two approaches.

## Connections ##

| Pico (pin #) | L293D (pin #) | MB102 (Power Supply) | DC Motor |
|-|-|-|-|
| GP28 (34)  | M1 Enable (1) | | |
| GP26 (31) | M1 Dir 0/1 (2) | | |
| GP27 (32) | M1 Dir 1/0 (7) | | |
| | M1 + (3) | | + |
| | M1 - (6) | | - |
| | Vmotor (8) **and** Vcc (16) | 5V +| |
| Ground (3,8,13,18,23,28,33,38) | Ground (4,5)| 5V - | |
