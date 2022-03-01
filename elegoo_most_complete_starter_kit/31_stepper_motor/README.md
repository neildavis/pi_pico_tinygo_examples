# Lesson 31: Stepper Motor #

This example demonstrates driving a [28BJY-48](https://components101.com/motors/28byj-48-stepper-motor)
[stepper motor](https://en.wikipedia.org/wiki/Stepper_motor) using a
[breakout board](https://www.electronicoscaldas.com/datasheet/ULN2003A-PCB.pdf) for the
[ULN2003](https://www.ti.com/lit/ds/symlink/uln2003a.pdf) IC.

## Driver Support ##

The [TinyGo drivers](https://github.com/tinygo-org/drivers) project provides a generic driver
([easystepper](https://github.com/tinygo-org/drivers/blob/release/easystepper/easystepper.go)) for
driving stepper motors, which would work with our components.

However, as before I have decided not to use the driver immediately and instead drive the stepper
using our own code to better understand the workings. The links above do a good job of explaining
how to interact with the ULN2003 in order to move the stepper motor so I won't repeat that here.

Subsequent lessons also use a stepper motor, so we'll use the driver there to show and compare
both approaches.

## Connections ##

Just as with continous DC motors, **never** connect a stepper motor directly to the Pico's power
supply. Instead we use the MB102 external power supply module to provide 5V DC to the ULN2003
breakout board, and from there to the 28BJY-48 stepper motor. Just be sure to connect both power
supply grounds as shown below.

The connections between the ULN2003 breakout board and the 28BJY-48 stepper motor are taken
care of for us by the pre-soldered socket on the PCB and the crimped connector on the stepper
motor cable, which can only be connected one way, so I'll omit those connections in the table below.
Just be sure to connect your stepper motor to the breakout board!

| Pico (pin #) | ULN2003 PCB | MB102 Power Supply |
|-|-|-|
| GP6 (9)  | IN1 |
| GP7 (10) | IN2 |
| GP8 (11) | IN3 |
| GP9 (12) | IN4 |
| | Power + | 5V + |
| Ground (3,8,13,18,23,28,33,38) | Power - | 5V - |
