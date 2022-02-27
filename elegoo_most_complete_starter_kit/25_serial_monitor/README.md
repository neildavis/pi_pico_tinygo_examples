# Lesson 25: The Serial Monitor #

This lesson in the Elegoo kit uses the Arduino IDE serial monitor as an input device.
Building on the previous lesson, it takes input from the serial bus to switch the
LEDs on individually or clear them all.

Unfortunately, we don't have this capabilty with the Pico and TinyGo.
Whilst you could wire up eight buttons to control the LEDs, there's little point,
especially since the last lesson used the 7HC595 specifically to avoid using a large
number of the Pico's pins for input.

I did consider using a shift register in reverse (PISO - Parallel In Serial Out)
to multiplex eight button inputs into three pins, but the kit does not provide this
component, and I'm trying to keep to the supplied components with this project.

Perhaps in the future, we will gain USB serial comms capability and be able to
achieve this lesson's goal, but for now we'll just move on to lesson 26.
