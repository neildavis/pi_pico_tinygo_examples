# Lesson 17: HC-SR501 PIR Sensor #

This lesson demonstrates use of the HC-SR501 Passive Infra-Red (PIR) motion detector sensor.
It's just a simple example of a digital input (PIR) driving a digital output (LED).
We use the Pico's internal LED in this example to keep the connections minimal.

You can play around with the module's 'time delay adjust' and 'sensitivity adjust' as well
as the 'trigger selection' jumper to see how the sensor behaviour changes.

## Connections ##

The PIR is best powered from 5V, but its output is 3.3V so we don't need to worry about using
voltage dividers or logic level converters before connecting it directly to our Pico's input pins.

| Pin (physical pin#) | HC-SR501 |
|-|-|
| GP6 (9) | Output/Detect |
| VBUS (40) | VCC |
| Ground (3,8,13,18,23,28,33,38) | GND |
