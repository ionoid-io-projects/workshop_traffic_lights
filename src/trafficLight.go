package main

import (
	"fmt"
	"time"
  	"os"
  	"os/signal"
  	"syscall"

	"github.com/stianeikeland/go-rpio"
)

func main() {

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	// Get the pin for each of the lights
	redPin := rpio.Pin(25)
	yellowPin := rpio.Pin(8)
	greenPin := rpio.Pin(7)

	// Set the pins to output mode
	redPin.Output()
	yellowPin.Output()
	greenPin.Output()

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		redPin.Low()
    	yellowPin.Low()
    	greenPin.Low()
    	os.Exit(0)
    }()

    defer rpio.Close()

  	// Turn lights off to start.
  	redPin.Low()
  	yellowPin.Low()
  	greenPin.Low()

	// A while true loop.
	for {
		// Red
		redPin.High()
		time.Sleep(time.Second * 3)

		// Red and yellow
		yellowPin.High()
		time.Sleep(time.Second)

		// Green
		redPin.Low()
		yellowPin.Low()
		greenPin.High()
		time.Sleep(time.Second * 5)

		// Yellow
		greenPin.Low()
		yellowPin.High()
		time.Sleep(time.Second * 2)

		// Yellow off
		yellowPin.Low()
	}
}