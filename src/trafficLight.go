package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strconv"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {

	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	// Get the pin for each of the lights
	RED_PIN, err := strconv.Atoi(os.Getenv("RED_PIN"))
	if err != nil {
		RED_PIN = 25
	}

	YELLOW_PIN, err := strconv.Atoi(os.Getenv("YELLOW_PIN"))
	if err != nil {
		YELLOW_PIN = 8
	}	

	GREEN_PIN, err := strconv.Atoi(os.Getenv("GREEN_PIN"))
	if err != nil {
		GREEN_PIN = 7
	}

	redPin := rpio.Pin(RED_PIN)
	yellowPin := rpio.Pin(YELLOW_PIN)
	greenPin := rpio.Pin(GREEN_PIN)

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
