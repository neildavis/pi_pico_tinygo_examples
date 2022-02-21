# Lesson 5: Digital Inputs #

This example makes use of push buttons with digital inputs to turn an LED on and off.
Pressing one button will turn the LED on, whilst pressing the other button will turn the LED off.

Don;t forget to use a protective resistor, e.g. 220 ohms, in series witht he LED.

Note: This example uses a loop to poll the button state. A better approach would be to use interrupts
but this is currently broken for more than one input pin at the time of writing with TinyGo v0.22.0.

## Connections ##

| Pin (physical pin#) | 'On' Button | 'Off' Button | LED |
|-|-|-|-|
| GP5 (7) | | | + |
| GP8 (11) | | + | |
| GP9 (12) | + | | |
| Ground (3,8,13,18,23,28,33,38) | - | - | - |
