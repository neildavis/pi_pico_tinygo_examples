# Lesson 19: Real Time Clock Module ##

This lesson demonstrates setting & reading from the (battery backed-up) Real-Time Clock
(RTC) module supplied with the Elegoo kit.

## **IMPORTANT**: Part Numbers & Voltage Levels ##

The Elegoo kit has been supplied with *at least* two different RTC modules using different
IC chips. The documentation I have states that the RTC is a
[DS1307](https://datasheets.maximintegrated.com/en/ds/DS1307.pdf) but also references an
'older' version with a picture of a [DS3231](https://datasheets.maximintegrated.com/en/ds/DS3231.pdf).
The Arduino code example is still setup for the DS3231.

The difference is important, since the DS1307 operates at **5V**, whereas the DS3231 operates
at **3.3V**. Obviously the Pico is a 3.3V device. If you have the newer DS1307, you will need
to use a [bi-directional logic level converter](https://learn.sparkfun.com/tutorials/bi-directional-logic-level-converter-hookup-guide/all)
between the DS1307 & the Pico, as well as powering it from the *VBUS* (5V) pin instead of *3V3(OUT)*.

## Driver Support ##

The [TinyGo drivers](https://github.com/tinygo-org/drivers) project contains drivers for both the
[DS1307](https://github.com/tinygo-org/drivers/tree/release/ds1307) and the
[DS3231](https://github.com/tinygo-org/drivers/tree/release/ds3231).

I was 'lucky' enough to get the DS3231 with my kit, so the connections and example code target that device.
If you have the DS1307, you should change the connections (using a logic level converter) and modify the
code to use the appropriate driver.

## Feedback ##

Once again, the elegoo example code for the Arduino supplied with the kit simply uses
the USB serial bus to print out the time, date and temperature received from the RTC
module, which are then read using the serial monitor in the Arduino IDE.
This isn't possible out-of-the-box with our Pico and TinyGo so we'll need some alternative
method of displaying a time, date and temperature.

Whilst jumping forward to lesson #22
(LCD Display Module) would give a superior user experience, I decided to stick with what
we have learnt and used so far and make use of the LED Matrix Module from lesson 15. We'll
simply define LED codes for all the characters we need, and flash them up one-by-one.

Since the LED Matrix Module takes some code to configure and drive, and it's nothing
new to us here, that code is in a separate file `ledmatrix.go` in the `main` package.
Also the character definitions are broken out into their own file `ledchars.go`, also
in the `main` package

## Connections ##

| Pin (physical pin # ) | LED Dot Matrix Module | DS3231 RTC |
|-|-|-|
| SPI1 CSn (12) | CS |
| SPI1 SCK (14) | CLK |
| SPI1 COPI/SDO/TX (15) | DIN |
| I2C1 SDA (31) | | SDA |
| I2C1 SCL (32) | | SCL |
| 3V3(OUT) (36) | | VCC |
| VBUS (40) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND | GND |
