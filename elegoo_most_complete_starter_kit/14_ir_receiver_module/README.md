# Lesson 14: IR Receiver Module #

This lesson demonstrates use of the AX-1838HS infra-red receiver module and the Arduino
remote control handset supplied in the Elegoo kit.

## Feedback ##

Once again, the Elegoo example code uses the USB serial bus and monitor in the Arduino IDE
to simply print out the (mapped) button name when receiving an IR code from the handset.

Since this feature is not available out-of-the-box with TinyGo on the Pico, we'll reuse
what we've learnt before to use an alternative feedback mechanism. In this case, just like
lesson 11 (Membrane Switch Module), we'll generate a unique tone on our passive buzzer for
each distinct key press from the IR remote control handset.

## Driver support ##

At the time of writing the [TinyGo drivers](https://github.com/tinygo-org/drivers) project
did not include a driver for infra-red receivers like the AX-1838HS. So I wrote a
minimal IR receiver driver which implements the basic [NEC protocol](https://www.sbprojects.net/knowledge/ir/nec.php)
used by the Arduino IR handset.

Fortunately my [PR](https://github.com/tinygo-org/drivers/pull/383) has now been merged into release [v0.20.0](https://github.com/tinygo-org/drivers/releases/tag/v0.20.0) so we can now use it directly instead of redirecting to my fork.

## Connections ##

Connect the passive buzzer and IR receiver to the Pico as follows:

| Pin (physical pin#) | IR Receiver | Passive Buzzer |
|-|-|-|
| GP15 (20) | | + |
| GP3 (5) | Y | |
| 3V3(OUT) (36) | R | |
| GND (3,8,13,18,23,28,33,38) | G| - |
