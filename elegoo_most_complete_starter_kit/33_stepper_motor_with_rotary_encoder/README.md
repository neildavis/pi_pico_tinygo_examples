# Lesson 33: Controlling Stepper Motor With Rotary Encoder #

This example builds on the previous lesson (32: Stepper Motor With Remote) by replacing the
IR remote control with a
[rotary encoder](https://en.wikipedia.org/wiki/Rotary_encoder) to control the
[28BJY-48](https://components101.com/motors/28byj-48-stepper-motor)
[stepper motor](https://en.wikipedia.org/wiki/Stepper_motor).

We'll once again use the [breakout board](https://www.electronicoscaldas.com/datasheet/ULN2003A-PCB.pdf)
for the [ULN2003](https://www.ti.com/lit/ds/symlink/uln2003a.pdf) IC.

## Driver Support ##

As in the previous example, we'll redirect the [TinyGo drivers](https://github.com/tinygo-org/drivers)
module to use [my fork](https://github.com/neildavis/drivers) in `go.mod` until the driver in the
upstream repository supports the 8-step mode operation that we require.

## Connections ##

The connections are the same as for the previous example, but with the substitution of the rotary encoder
for the IR receiver.

Just as with continous DC motors, **never** connect a stepper motor directly to the Pico's power
supply. Instead we use the MB102 external power supply module to provide 5V DC to the ULN2003
breakout board, and from there to the 28BJY-48 stepper motor. Just be sure to connect both power
supply grounds as shown below.

The connections between the ULN2003 breakout board and the 28BJY-48 stepper motor are taken
care of for us by the pre-soldered socket on the PCB and the crimped connector on the stepper
motor cable, which can only be connected one way, so I'll omit those connections in the table below.
Just be sure to connect your stepper motor to the breakout board!

| Pico (pin #) | Rotary Encoder | ULN2003 PCB | MB102 Power Supply |
|-|-|-|-|
| GP3 (5) | SW | | |
| GP6 (9) | | IN1 | |
| GP7 (10) | | IN2 | |
| GP8 (11) | | IN3 | |
| GP9 (12) | | IN4 | |
| 3V3(OUT) (36) | + | | |
| | | Power + | 5V + |
| Ground (3,8,13,18,23,28,33,38) | G | Power - | 5V - |
