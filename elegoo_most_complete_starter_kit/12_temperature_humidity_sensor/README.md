# Lesson 12: DHT11 Temperature and Humidity Sensor #

This example demonstrates the use of a DHT11 temperature and humidity sensor.

## **IMPORTANT:** Voltage Levels ##

The DHT11 may or may not work with the 3.3V supplied from the Pi Pico.
According to [Mouser](https://www.mouser.com/datasheet/2/758/DHT11-Technical-Data-Sheet-Translated-Version-1143054.pdf):
"DHT11’s power supply is 3-5.5V" but [Adafruit also say](https://learn.adafruit.com/dht?view=all)
"Sometimes 3.3V power isn't enough in which case try 5V power."

The Arduino Uno R3 supplied with the Elegoo kit operates on 5V signalling so is fine to use the DHT11 directly at 5V.
The Pi Pico uses 3.3V for signalling on its pins. If the DHT11 is powered via 5V (e.g. from the Pico VBUS pin)
then the DATA pin used for reading data will also use 5V signalling which could damage the Pico.

Unlike lesson 10 (HC-SR04 ultrasonic sensor) we cannot even use a voltage divider to drop the DATA pin from
5V to 3.3V since the pin is used for both sending data TO **and** receiving data FROM the device.
The voltage divider works in only one direction.

A potential solution may be to connect the DATA pin twice, once directly and the other via a voltage divider, and
use diodes to ensure the current takes the correct path for reading from & writing to the device.
However this has not been tested and the connections shown below use 3.3V for both power & signalling.

## Feedback ##

Again, the elegoo kit used the Arduino USB serial bus & monitor to simply print out the temperature & humidity
readings received from the DHT11. Since this isn't availble 'out-of-the-box' with the pico, we have to be
creative in how we communicate the readings.

This project uses two RGB LEDs (you can use 3 single colour LEDs instead of either/both RGB LEDs if you wish)
to indicate changes to the ambient temperature and humidity.
The code will take a few initial readings of the environment and take the average of these readings to mean
'normal', indicated by both LEDs showing green.
If the temperature or humidity raise above a certain threshold the corresponding LED will turn red.
Likewise if the temperature or humidity drop below a certain threshold the corresponding LED will turn blue.

Try breathing on the sensor, which should cause both temperature and humidty to increase and turn the LEDs red.
Alternatively try fanning the sensor to reduce its temperature and turn one LED blue.

## Connections ##

| Pin (physical pin#) | RGB LED 1 (Temp) | RGB LED 2 (Humidity) | DHT11 |
|-|-|-|-|
| GP0 (1) | Red | | |
| GP1 (2) | Green | |
| GP2 (4) | Blue | |
| GP3 (5) | | Red | |
| GP4 (6) | | Green | |
| GP5 (7) | | Blue |
| GP15 (20) | | | DATA |
| Ground (3,8,13,18,23,28,33,38) | Cathode | Cathode | GND |
| 3V3(OUT) (36)| | | VCC |
