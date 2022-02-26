# Lesson 22: LCD Module #

This lesson makes use of the [LCM1602A](https://www.displayfuture.com/Display/datasheet/alphanumeric/LCM1602A.pdf)
(16x2) LCD Module to display some text. The program increments a counter every second and plays a variant of [Fizz Buzz](https://en.wikipedia.org/wiki/Fizz_buzz) displaying the results on the LCD display.

## Driver Support ##

TinyGo [drivers](https://github.com/tinygo-org/drivers) does not provide specific support for the LCM1602A.
However, our code will be interfacing directly with the module's internal
[SPLC780D](https://www.waveshare.com/datasheet/LCD_en_PDF/SPLC780D.pdf) controller/driver IC, which is pin
compatible with the Hitachi [HD44780](https://www.sparkfun.com/datasheets/LCD/HD44780.pdf) chip which *is*
provided for by [drivers](https://github.com/tinygo-org/drivers/tree/release/hd44780).

## **IMPORTANT:**  Voltage Levels ##

The LCM1602A is another one of those components designed for 5V devices like the Ardino Uno, rather than 3.3V
devices like the Pi Pico. So for best results we should power it at 5V for both the controller and the backlight.
However, since the databus is bi-directional, that could potentially **damage the Pi Pico** unless we include
other circuitry to compensate.
For example we could use a [bi-directional logic level converter](https://www.sparkfun.com/products/12009)
or there is a [hardware hack](https://www.codrey.com/electronic-circuits/hack-your-16x2-lcd/)
that can convert the device to work properly at 3.3V by boosting the backlight voltage.

Assuming neither of those options are available (none are included with the Elegoo kit), we have three options:

1. Try to run it at 3.3V
2. Use the driver in 'write-only' mode so that it never reads from the device at 5V.
3. Accept the risk and read from the device at 5V. This could potentially damage your Pi Pico beyond repair.

Looking at these in turn, in my experience, it *can* be run on 3.3V but the text is very feint and almost
impossible to read, so it's not a great option. The driver does include a 'write-only' mode which allows you
to keep the `RW` pin at ground and never read the `busy flag` from the device after writing instructions to it.
However, I was unable to get this to work at all despite [extensive efforts](https://github.com/tinygo-org/drivers/issues/380). That leaves us with option 3. Unhelpfully neither of the datasheets for the LCM1602A module itself
or the SPLC780D controller specify a *typical* voltage out for logic High (VOH). However the *minimum* VOH is
speciifed as **2.4V** which is verly low for a 5V device, so it's *possible* (but by no means guaranteed) that the
*typical* VOH could also be relatively low and within tolerance for the 3.3V Pico (VIH max ~= 3.6V)

### **DISCLAIMER** ###

I have not verified the typical VOH of the LCM1602A LCD module. However I have connected it *directly* to my Pico
without use of any level converters or hardware hacks and powered it from 5V without issues. This does not mean
you will get the same results. Ideally you should use a level converter, but proceed at your own risk! If in doubt,
run the device from 3.3V or use a level converter. I will not be held responsible if you damage your Pico by exposing
it to unsafe voltage levels.

## Connections ##

The LCD requires two power connections, and an additional potentiometer across the power rail to work as a contrast
adjustment.

| Pin (physical pin#) | LCD | Potentiometer (10K) |
|-|-|-|
| 16 (21) | EN | |
| 17 (22) | RS | |
| 18 (24) | D4 (**use a 5V<->3.3V level converter**) | |
| 19 (25) | D5 (**use a 5V<->3.3V level converter**) | |
| 20 (26) | D6 (**use a 5V<->3.3V level converter**) | |
| 21 (27) | D7 (**use a 5V<->3.3V level converter**) | |
| 22 (29) | RW | |
| VBUS 5V (40) / 3V3(OUT) (36) | VDD, A | A |
| Ground (3,8,13,18,23,28,33,38) | VSS, K| B |
| | V0 (contrast) | W (Wiper) |
