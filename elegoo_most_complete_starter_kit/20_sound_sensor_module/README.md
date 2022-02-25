# Lesson 20: Sound Sensor Module #

This example demonstrates use of the Sound Sensor Module supplied with the Elegoo kit.
We read a digital input for a 'threshold' trigger which lights the Pico's integral LED.
We also read an analog sound level using the Pico's ADC, and show how that changes over
time.

## Feedback ##

The Elegoo example code for the Arduino writes out the raw analog input from the sensor
to the USB serial bus, where it can be seen in the Arduino IDE using the serial monitor.

Since we can't do this out-of-the-box with our Pico and TinyGo, we'll again make use of
the components we've used so far. In particular, we'll use the LED Matrix Module from
lesson 15 to display a crude sound level meter that shows time on the x-axis (rows)
and the analog signal intensity on the y-axis, updating at a rate of 16 samples/sec.

## Connections ##

| Pin (physical pin # ) | LED Dot Matrix Module | Sound Sensor Module |
|-|-|-|
| SPI1 CSn (12) | CS |
| SPI1 SCK (14) | CLK |
| SPI1 COPI/SDO/TX (15) | DIN |
| ADC0 (31) | | AO |
| GP27 (32) | | DO |
| 3V3(OUT) (36) | | + (VCC) |
| VBUS (40) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND | G (GND) |
