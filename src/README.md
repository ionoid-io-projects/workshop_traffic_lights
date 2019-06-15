# Introduction
Making leds blink 
Follow link below for pin connections 
https://projects.raspberrypi.org/en/projects/physical-computing/11

# How to
Compile led.go like this
```
go get github.com/stianeikeland/go-rpio
env GOOS=linux GOARCH=arm GOARM=6 go build TrafficLight.go
```
Copy the generated file to your raspberry pi device and execute it with this command

```
./TrafficLight
```

Happy blinking 