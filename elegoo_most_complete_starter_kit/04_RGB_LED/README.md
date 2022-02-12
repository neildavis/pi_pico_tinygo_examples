# Lesson 4: RGB LED

This example controls an RGB LED via three separate PWM channels from a Pi Pico.

It starts with the LED in the Red colour state, then fades to Green, then fades to Blue and finally fades back to the initial Red colour. This cycles through most of the colours that can be achieved from the LED.

Connect The RGB LED legs (via suitable protective resistors - e.g. 220 ohms - for the RGB annodes) to the Pico's GPIO pins as follows:

| LED Leg | GPIO Pin | Physical Pin # |
|-|-|-|
| Red | GPIO0 | 1 |
| Green | GPIO2 | 4 |
| Blue | GPIO4 | 6 |
| Cathode | Ground | 3,8,13,18,23,28,33,38 |
