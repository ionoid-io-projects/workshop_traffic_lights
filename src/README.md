# Introduction
Making leds blink 
Follow link below for pin connections 
https://projects.raspberrypi.org/en/projects/physical-computing/11

# How to
Compile led.go like this
```
go get github.com/stianeikeland/go-rpio
env GOOS=linux GOARCH=arm GOARM=6 go build trafficLight.go
```
Copy the generated file to your Raspberry Pi device and execute it with this command

```
env RED_PIN=12 YELLOW_PIN=13 GREEN_PIN=18 ./trafficLight
```

Happy blinking 