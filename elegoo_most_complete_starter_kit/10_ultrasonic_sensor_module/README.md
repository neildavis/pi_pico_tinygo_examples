#  Lesson 10: Ultrasonic Sensor Module #

This example demonstrates the use of a [HC-SR04](https://www.letscontrolit.com/wiki/index.php/HC-SR04) ultrasonic sensor module to measure distances.

The HC-SR04 module provides a range of 2cm-400cm non-contact measurement, with an aaccuracy that can reach as low as 3mm

The [TinyGo Drivers](https://pkg.go.dev/tinygo.org/x/drivers) project already includes a driver for the HC-SR04, so we'll
use that.

If you're curious how this works internally, check out the [HC-SR04 datasheet](https://cdn.sparkfun.com/datasheets/Sensors/Proximity/HCSR04.pdf) and inspect the [driver code](https://github.com/tinygo-org/drivers/blob/release/hcsr04/hcsr04.go).

## **IMPORTANT:** Voltage Levels ##

As descibed in the [linked page](https://www.letscontrolit.com/wiki/index.php/HC-SR04) above, the HC-SR04 is designed for 5V circuits. The Pi Pico runs on 3.3V so cannot be used in the same congfiguration as e.g. with a 5V Arduino Uno. 

As with the Servo example (lesson 9) we can workaround this to some extent, by connecting the HC-SR04 VCC input to the Pico's 5V VBUS pin (physical pin #40). Signalling ***to*** the device's 'trigger' pin using 3.3V should still work fine. However, the device's 'echo' pin used to read the measured distance will use the same voltage as provided to VCC. Without additional circuitry, we have two options:

1. Power the device using the Pico's 3.3V (OUT) supply (physical pin 36), and read the 'echo' at the same voltage. However, it has been [demonstrated](https://forums.raspberrypi.com/viewtopic.php?p=183386#p183386) that this can lead to wildly inaccurate results.
2. Power the device using the Pico's 5V VBUS supply (physical pin 40), and hope that reading the 'echo' at 5V will not damage the Pico. I have seen examples of this done, however, officially the Pico and RP2040 are **NOT 5V tolerant** so this should be avoided if you don't wish to fry your Pico!

Fortunately there is a simple solution using a minimal amount of extra circuitry - two resistors in series - to create a [voltage divider](https://learn.sparkfun.com/tutorials/voltage-dividers) which can lower the 5V from the HC-SR04's 'echo' pin to the 3.3V required by the Pico. I highly recommend you read this article and understand the concepts before assembling this project as I will **not** be responsible if you fry your pico board with 5V!

## Feedback ##
The original Arduino based version of this lesson simply printed out the measured distance value via the serial bus which could be seen in the Arduino IDE serial monitor. However, this isn't possible on the Pico with a simple USB connection (by default the UART is connected to GPIO rather than USB CDC)

Whilst there are solutions to this using e.g. another Pico as a 'Pico Probe', I decided to modify the example to drive an RGB LED colour based on the measured distance:

|Distance|Colour|
|-|-|
|<10cm|Red|
|10-20cm|Yellow|
|>20cm|Green|

## Circuit Setup ##

With the above information in mind, I opted to use 5V VBUS to power the HC-SR04 and use a voltage divider to drop the 'echo' pin from 5V to 3.3V before interfacing with the Pico.

I found the best values for the resistors supplied with the Elegoo kit were:

* R1 = 5.1K
* R2 = 10K

Feeding these values into the voltage divider equation:

 Vout = Vin * R2/(R1+R2)

leads to a divided voltage (Vout) = **3.31V** which is well within tolerance for the Pico.

## Connections ##

With the above in mind, here is how to connect the Pico, RGB LED and HC-SR04:

| Pico Pin (physical pin#) | HC-SR04 | RGB LED|
|-|-|-|
| VBUS (40) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND | Cathode |
| GP0 (1) | | Red |
| GP2 (4) | | Green |
| GP4 (6) | | Blue |
| GP14 (19) | Trig |
| GP15 (20) | Echo (**via R1/R2 - 5.1K/10K - 5V -> 3.3V voltage divider**)| 
