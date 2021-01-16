# Go Temperature Monitor for AM2302 on Raspberry Pi

This small program will use the sensor AM2302 with [this go-dht library](https://github.com/MichaelS11/go-dht) to send data to [Adafruit](https://io.adafruit.com).

## How to Use it?

The easiest way is to have go installed in your Raspberry Pi. Then `go run main.go` from a tmux terminal. Remember that if you use Ubuntu or Ubuntu server you need to run it using sudo `sudo go run main.go`. It looks like that only sudo has access to the GPIO ports.

## Env Variables

The program needs a environment variable: `ADAFRUIT_IO_KEY`. This can be added into a file called `.env`. An example of the content of this file is:

```
ADAFRUIT_IO_KEY=XYzaBC
```

This is my first program enterly written on a Raspberry Pi (well, using ssh so I can still browse the internet) and using Vim! I used [this](https://tpaschalis.github.io/vim-go-setup/) guide to set up Vim for working with go. 

The overall program is based on [this](https://www.jeremymorgan.com/tutorials/raspberry-pi/how-to-iot-adafruit-raspberrypi/) project written in Python and [this](https://www.jeremymorgan.com/tutorials/go/get-temperature-raspberry-pi-go/) post explaining how to read the temperature from go. Thanks Jeremy Morgan!
