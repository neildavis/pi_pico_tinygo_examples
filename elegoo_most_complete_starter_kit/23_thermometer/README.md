# Lesson 23: Thermometer #

This example demonstrates the use of the [NTC thermistor](https://en.wikipedia.org/wiki/Thermistor)
supplied with the Elegoo kit. A thermistor is a resistor that varies its resistance based on its temperature.

When used as one half of a [voltage divider](https://learn.sparkfun.com/tutorials/voltage-dividers/all),
with a known resistance value for the other half (the 'balance' resistor), we can read Vout from the voltage
divider and calculate the ambient temperature.

## Calculation ##

Although there's nothing new or exciting for us from a circuit/component point of view (it's just another
analog pin read with the ADC), this example introduces us to the
'[`Beta` (β) equation](https://www.digikey.co.uk/en/maker/projects/how-to-measure-temperature-with-an-ntc-thermistor/4a4b326095f144029df7f2eca589ca54)'
which is necessary to convert the analog value we read from the thermistor voltage divider into a temperature.

From the 'ordering code' section of the [datasheet](http://focusens.com/data/upfile/1504/2015041509051812.pdf)
and our part number (`MF52-103F-3950`) we can determine:

1. The 'resistance value at 25°C': `103` = 10KΩ  (10 to the power 3)
2. The 'tolerance' of the resistor: `F` = ±1%
3. The '`Beta`' value: `β` = 3950

We will add a 10KΩ 'balance' resistor to complete the voltage divider, and then plug these values into the
beta equation to determine the correct temperature from our analog signal.

## Feedback ##

We now have two possible ways to read the temperature:

1. The DHT11 temperature & humidity sensor (see lesson 12).
2. The Thermistor.

We'll read both and use the LCD Display module from lesson 22 to display and compare the results.

Note: The Pico/RP2040 also includes an internal built-in temperature sensor as part of its ADC.
In other platforms (e.g. MicroPython) this is accessible as `ADC4`.
However, this is not (yet) exposed by the TinyGo `machine` package.

## Connections ##

### Thermistor / 10KΩ Divider (MF52-103F-3950) ###

| Pin (physical pin#) | R1 (Thermistor) | R2 (10KΩ 'Balance') |
|-|-|-|
| ADC1 (32) | - | + |
| 3V3(OUT) (36) | + | |
| Ground (3,8,13,18,23,28,33,38) | | - |
|||

### DHT11 ###

| Pin (physical pin#) | DHT11 |
|-|-|
| GP28 (34) | DATA |
| 3V3(OUT) (36)| VCC |
| Ground (3,8,13,18,23,28,33,38) | GND |
||

### LCD ###

The LCD requires two power connections, and an additional potentiometer across the power rail to work as a contrast
adjustment.

| Pin (physical pin#) | LCD | Potentiometer (10K) |
|-|-|-|
| 16 (21) | EN | |
| 17 (22) | RS | |
| 18 (24) | D4 (**use a 5V<->3.3V level converter**) | |
| 19 (25) | D5 (**use a 5V<->3.3V level converter**) | |
| 20 (26) | D6 (**use a 5V<->3.3V level converter**) | |
| 21 (27) | D7 (**use a 5V<->3.3V level converter**) | |
| 22 (29) | RW | |
| VBUS 5V (40) | VDD, A | A |
| Ground (3,8,13,18,23,28,33,38) | VSS, K| B |
| | V0 (contrast) | W (Wiper) |
| | |
