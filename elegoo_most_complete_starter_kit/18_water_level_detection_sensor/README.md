# Lesson 18: Water Level Detection Sensor Module #

This example demonstrates use of the RB-02S048 Liquid Level Detection Sensor Module
supplied with the elegoo kit to measure the level of water when it's (partially)
submerged.

Really there's nothing new here, it's just another example of reading an analog value
through the Pico's ADCs. Here I've used ADC2 on physical pin #34.

## **IMPORTANT**: Voltage Levels ##

The RB-02S048 module is designed to be powered at 5V, and provide a typical ~4.2V
output on it's pin. Reading this directly on our 3.3V Pico could damage it.

Our options here are:

1. Run the module at 5V and read directly. My testing indicated that the actual output
level didn't reach 3.3V so this is likely OK. However YMMV, and I don't recommend it.
If you do want to go down this route, test your own component with a multimeter first
to verify the voltage levels you are seeing.

2. Run the module at 5V and use a (e.g. 5.1K/10K)
[voltage divider](https://learn.sparkfun.com/tutorials/voltage-dividers/all)
or a 'Logic Level Converter' to drop the output voltage before interfacing with the Pico ADC.
Since I wasn't seeing high output voltages anyway I didn't test this method.

3. Run the module at 3.3V and read directly. It's possible the module will not behave 100%
correclty, and in particular its sensitivity may be affected. However this is the route
I chose for this example since our chosen feedback method is quite coarse grained anyway
(only 9 distinct levels) so sensitivity/accuracy is not that important.

## Feedback ##

Once again, the elegoo example code for the Arduino supplied with the kit simply uses
the USB serial bus to print out the values received from the sensor, which are then
read using the serial monitor in the Arduino IDE.

This is neither appealing or possible out-of-the-box with our Pico and TinyGo so we'll
do something more interesting! We'll make use of the LED Matrix Module from lesson 15.
We'll split the matrix display down the middle to show two 4x8 columns to represent the
water level on a scale of 0 (no liquid dectected) to 8 (high level detected).
One column will show the level in real time, and the other will keep track of the maximum
level we have seen for a period of a few seconds before resetting.

Since the LED Matrix Module takes some code to configure and drive, and it's nothing
new to us here, that code is in a separate file `ledmatrix.go` in the `main` package.

## Connections ##

| Pin (physical pin # ) | LED Dot Matrix Module | RB-02S048 Liquid Level Sensor|
|-|-|-|
| SPI1 CSn (12) | CS |
| SPI1 SCK (14) | CLK |
| SPI1 COPI/SDO/TX (15) | DIN |
| ADC2 (34) | | S (Output) |
| 3V3(OUT) (36) | | + (VCC) |
| VBUS (40) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND | - (GND) |
