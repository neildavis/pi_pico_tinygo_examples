# Lesson 32: Controlling Stepper Motor With Remote #

This example builds on the previous lesson (31: Stepper Motor) by driving the
[28BJY-48](https://components101.com/motors/28byj-48-stepper-motor)
[stepper motor](https://en.wikipedia.org/wiki/Stepper_motor) from a distance using an IR
remote control.
We'll once again use the [breakout board](https://www.electronicoscaldas.com/datasheet/ULN2003A-PCB.pdf)
for the [ULN2003](https://www.ti.com/lit/ds/symlink/uln2003a.pdf) IC.

## Driver Support ##

In contrast to the previous example, this time we'll use the
[easystepper](https://github.com/tinygo-org/drivers/blob/release/easystepper/easystepper.go)
driver provided by the [TinyGo drivers](https://github.com/tinygo-org/drivers) project to
simplify our code.

Interestingly, at the time of writing (TinyGo drivers v0.19.0) the driver used only a
[*4-step* model](https://github.com/tinygo-org/drivers/blob/v0.19.0/easystepper/easystepper.go#L121-L149)
for stepping the motor, where two coils are always energised (in the sequence 12-23-34-41),
whereas we previously used an *8-step* model, alternating between individual and pairs of coils energizing
(in the sequence 1-12-2-23-3-34-4-41).

As it turns out, the 4-step model does not work with our 28BJY-48 & ULN2003!
I [reported](https://github.com/tinygo-org/drivers/issues/393) this as an issue, and submitted a
[pull request](https://github.com/tinygo-org/drivers/pull/394) to resolve it.
Thankfully this PR has been merged and included in release
[v0.21.0](https://github.com/tinygo-org/drivers/releases/tag/v0.21.0) of the drivers.

## Connections ##

The connections are the same as for the previous example, but with the addition of the IR receiver
like we used back in lesson 14 (IR Receiver Module)

Just as with continous DC motors, **never** connect a stepper motor directly to the Pico's power
supply. Instead we use the MB102 external power supply module to provide 5V DC to the ULN2003
breakout board, and from there to the 28BJY-48 stepper motor. Just be sure to connect both power
supply grounds as shown below.

The connections between the ULN2003 breakout board and the 28BJY-48 stepper motor are taken
care of for us by the pre-soldered socket on the PCB and the crimped connector on the stepper
motor cable, which can only be connected one way, so I'll omit those connections in the table below.
Just be sure to connect your stepper motor to the breakout board!

| Pico (pin #) | IR Receiver | ULN2003 PCB | MB102 Power Supply |
|-|-|-|-|
| GP3 (5) | Y | | |
| GP6 (9) | | IN1 | |
| GP7 (10) | | IN2 | |
| GP8 (11) | | IN3 | |
| GP9 (12) | | IN4 | |
| 3V3(OUT) (36) | R | | |
| | | Power + | 5V + |
| Ground (3,8,13,18,23,28,33,38) | G | Power - | 5V - |
