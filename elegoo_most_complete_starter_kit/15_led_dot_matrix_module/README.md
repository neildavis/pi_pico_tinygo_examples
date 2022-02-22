# Lesson 15: MAX7219 LED Dot Matrix Module #

This example demonstrates use of the kit's LED dot matrix module incoroporating
a [MAX7219](https://datasheets.maximintegrated.com/en/ds/MAX7219-MAX7221.pdf) driver IC
to scroll some text across the display.

## Voltage & Power Levels ##

The LED Dot Matrix Module is designed for 5V use and can draw up to 2A. Some, but not
all USB connections can provide for 2A.

Using the formula P = VI, we can see that this means the module may use up to 10W power.
Consequently, if we connect this to 3.3V then to maintain the power requirement, it may
draw a current of up to 3A, which is more than available from most USB power supplies.
So we'll use 5V for the power supply (VCC,GND) and 3.3V for the SPI logic (DIN,CS,CLK)

Fortunately, the device is 'write-only' so we never need to read from it at 5V.

## Driver Support ##

[TinyGo drivers](https://github.com/tinygo-org/drivers) already provides a
[driver for the MAX7219](https://github.com/tinygo-org/drivers/tree/release/max72xx)
so we will make use of that to keep the code simpler. The internal workings of driving
the MAX7219 are not that complicated, so check out the code for the driver if you're
curious to see how it works.

## Connections ##

The connections shown assume the device is powered by USB that can provide up to 2A current.
If this doesn't work for you then you may consider using the kit's external power supply adapter.
Just ensure the grounds are connected together if using more than one power supply.

Since the MAX7219 is using the Serial Peripheral Interface
([SPI](https://learn.sparkfun.com/tutorials/serial-peripheral-interface-spi/all))
we need to choose the pins carefully. Unlike other examples, the pins here are not freely
interchangable. We need to ensure that the LED Matrix Module pins (DIN, CS, CLK) match up with
the corresponding pins (COPI/TX, CSn, SCK) of one of the Pico's SPI peripherals (SPI0/SPI1)

The connections below are using SPI1 periperal on the specified pins.

| Pin (physical pin # ) | LED Dot Matrix Module |
|-|-|
| SPI1 CSn (12) | CS |
| SPI1 SCK (14) | CLK |
| SPI1 COPI/SDO/TX (15) | DIN |
| VBUS (40) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND |
