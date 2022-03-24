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
does not include a driver for infra-red receivers like the AX-1838HS. I have written a
minimal IR receiver driver in [my fork](https://github.com/neildavis/drivers/tree/irremote/irremote)
of the drivers which implements the basic [NEC protocol](https://www.sbprojects.net/knowledge/ir/nec.php)
used by the Arduino IR handset.

We use the [`replace` directive](https://go.dev/ref/mod#go-mod-file-replace) in our `go.mod` fille to
substitute my fork for the `drivers` package. This way we don't need to modify our code if/when the
driver is included, we simply remove the `replace` directive from `go.mod`

## Connections ##

Connect the passive buzzer and IR receiver to the Pico as follows:

| Pin (physical pin#) | IR Receiver | Passive Buzzer |
|-|-|-|
| GP15 (20) | | + |
| GP3 (5) | Y | |
| 3V3(OUT) (36) | R | |
| GND (3,8,13,18,23,28,33,38) | G| - |
