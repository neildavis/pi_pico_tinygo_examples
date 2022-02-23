# Lesson 16: GY-521 Module #

This example demonstrates use of the GY-521 Inertia Measurement Unit (IMU) module
supplied with the Elegoo kit. This module combines both an accelerometer and a gyro
in the same package.

## Driver Support ##

The GY-521 module is actually a simple breakout board based around the
[InvenSense MPU-6050](https://invensense.tdk.com/wp-content/uploads/2015/02/MPU-6000-Datasheet1.pdf)
chip. The [TinyGo drivers](https://github.com/tinygo-org/drivers) package includes a
[driver](https://github.com/tinygo-org/drivers/tree/release/mpu6050) for communicating
with the MPU-6050 via I2C.

## Feedback ##

This is yet another example where the original example code from the Elegoo Arduino based kit
makes use of the USB serial bus and IDE monitor to output the values received from the device.
Without this capability out-of-the-box with the Pico, we need an alternative feedback mechanism.

Alas, since the MPU-6050 is a 6-DOF (Six Degrees of Freedom) sensor, it provides **six** values
for us to feedback in some way. Whilst I considered jumping forwards to lesson #22 (LCD Display),
for the purposes of this example I have chosen to provide feedback as follows:

### Accelerometer ###

The accelerometer measures inertia from movement in 3 axis (Ax,Ay,Az).
These values will be indicated by using the LED matrix display from the previous lesson 15.
The LED matrix will display a 2x2 lit 'ball' in the centre of the display which will move around
and vary in intensity in relation to the acclerometer values as shown in the following table:

| Measurement | LED Matrix |
|-|-|
| Ax | X-axis (rows) |
| Ay | Y-axis (columns) |
| Az | Intensity |

### Gyroscope ###

The gyroscope measures the current 'attitude' of the device in 3D space axis (Gx,Gy,Gz).
These values will be indicated by mapping each axis (Gx,Gy,Gz) to a colour component (R,G,B)
of a RGB LED, such that the perceived colour varies as the device's orientation changes.

| Measurement | RGB LED Component |
|-|-|
| Gx | Red |
| Gy | Green |
| Gz | Blue |

## Connections ##

Fortunately, the MPU-6050 is specified as a 3.3V device so we can power it from
the Pico's 3.3V(OUT) pin and don't need to worry about any conversion, voltage dividers etc.

The MPU-60505 uses the [I2C](https://learn.sparkfun.com/tutorials/i2c/all) bus to interface with
the Pico.
Similar to the SPI connection in the previous lesson (15: LED Matrix Module), we need to use
specific pairs of pins designated as a I2C peripherals (IC20/IC21) on the Pico. Here I am using
I2C1 on physical pins 31 and 32 (I2C1 SDA/I2C1 SCL respectively).

### GY-521 ###

| Pin (physical pin # ) | LED Dot Matrix Module |
|-|-|
| I2C1 SDA (31) | SDA |
| I2C1 SCL (32) | SCL |
| 3v3(OUT) (36) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND |

### RGB LED ###

| Pin (physical pin # ) | RGB LED |
|-|-|
| GP0 (1) | Red |
| GP1 (2) | Green |
| GP2 (4) | Blue |
| Ground (3,8,13,18,23,28,33,38) | Cathode |

### LED Matrix Module ###

| Pin (physical pin # ) | LED Dot Matrix Module |
|-|-|
| SPI1 CSn (12) | CS |
| SPI1 SCK (14) | CLK |
| SPI1 COPI/SDO/TX (15) | DIN |
| VBUS (40) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND |
