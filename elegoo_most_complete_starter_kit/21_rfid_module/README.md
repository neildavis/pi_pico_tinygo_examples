# Lesson 21: RC522 RFID Module #

This example makes use of the RC522 RFID Reader Module. The RC522 is usually supplied as a kit
and contains the [MFRC522](https://www.nxp.com/docs/en/data-sheet/MFRC522.pdf)
from NXP semi-conductors boken out on a board inlcuding an antenna and SPI/I2C interface pins,
as well as a 'fob' or card known as a '[Proximity IC Card](https://en.wikipedia.org/wiki/Proximity_card)'
or PICC.

## Status ##

This example is not currently implemented due to lack of driver support.
I will come back to this after completing other lessons.

## Driver Support ##

[TinyGo drivers](https://github.com/tinygo-org/drivers) does not currently include support for the
MFRC522. I plan to develop a minimal SPI based driver in due course.

## Connections ##

The RFC522 & MFRC522 support both SPI & I2C interfaces. The connections here use the SPI0
peripheral on the Pico.

| Pin (physical pin#) | RC522 RFID Module |
|-|-|
| GP3 (5) | RST |
| SPI0 CIPO/RX/SDI (6) | MISO |
| SPI0 CSn (7) | SDA |
| SPI0 SCK (9) | SCK |
| SPI0 COPI/TX/SDO (10) | MOSI |
| 3V3(OUT) (36) | 3.3V |
| Ground (3,8,13,18,23,28,33,38) | GND |
