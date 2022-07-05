# Lesson 29: DC Motors #

This example builds on the previous example and demonstrates how to use a
[relay](https://en.wikipedia.org/wiki/Relay)
in conjunction with the [L293D](https://www.ti.com/lit/ds/symlink/l293d.pdf) to
drive the motor.

The relay supplied with the Elegoo kit is a 
[Songle SRD-05VDC-SL-C](http://www.songlerelay.com/Public/Uploads/20161104/581c81ac16e36.pdf)
This product code means the coil voltage is rated at *5V DC*, it is of *sealed* type
construction, the coil draws *0.36 Watts* and the contact is the *conversion 'C'* type
(meaning the contactor switches between two ouput contacts instead of a simple
Normally Open (NO) or Normally Closed (NC) type with a single output contact)

In this example we will use only one output contact of the relay to provide VCC to the
motor, and connect the other terminal of the motor directly to ground.

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

In contrast to the previous lesson (29: DC Motors), in this example we will use the
[L293D driver](https://github.com/tinygo-org/drivers/blob/release/l293x/l293x.go) from the
[TinyGo drivers](https://github.com/tinygo-org/drivers) project.

Note: The driver supports both digital and PWM modes. Since we have a relay between the
L293D and the DC motor, it doesn't make sense to use PWM in this case, so we will use the
digital version and the motor control will be just on/off at full speed.

## Connections ##

| Pico (pin #) | L293D (pin #) | MB102 (Power Supply) | Relay | DC Motor |
|-|-|-|-|-|
| GP28 (34)  | M1 Enable (1) | | | |
| GP26 (31) | M1 Dir 0/1 (2) | | | |
| GP27 (32) | M1 Dir 1/0 (7) | | | |
| | M1 + (3) | | Coil + | |
| | M1 - (6) | | Coil - | |
| | | | Contactor Out (either) | + |
| | | 5V+ | Contactor In | |
| | Vmotor (8) **and** Vcc (16) | 5V +| |
| Ground (3,8,13,18,23,28,33,38) | Ground (4,5)| 5V - | | - |
