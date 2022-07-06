# Buttons using Interrupts #

This is a small example to demonstrate ~~[a bug](https://github.com/tinygo-org/tinygo/issues/2720)~~ 
using GPIO interrupts with 2 buttons to control an LED on/off state.

It was primarily made as the 'ideal' version of the elegoo kit lesson 5 (Digital Inputs) which uses
a loop to poll the button states, since at the time it was developed using TinyGo v0.22.0 which had
a [bug](https://github.com/tinygo-org/tinygo/issues/2720) preventing use of interrupts on more then one
GPIO pin simultaneously.

Thankfully this is now fixed and works correctly as of TinyGo v0.24.0
