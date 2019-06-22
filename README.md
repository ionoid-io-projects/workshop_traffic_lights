# Introduction
Making leds blink 

# Wiring
![traffic light](doc/img/trafficLight.jpg)

# Code

## 1- Power your raspberry

You can achive it with connecting it to your pc trought the Micro USB Port of the raspberry pi

![power](doc/img/1-min.jpg)

## 2- Connect to your raspberry pi
Using putty if you're on windows, Ssh if you're on a linux based os
Follow the following instruction if you dont know how to connect to raspberry pi
[Connect to raspberry pi using Putty](https://github.com/ionoid-io-projects/workshop/blob/master/doc/od-iot-raspbian-rpi-zero-windows.md#5-first-boot)

## 3- Download Traffic light binary file

Assuming you're connected with... copy and past this command
If you're using Raspberry zero
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_traffic_lights/master/bin/arm6/trafficLight
```

If you're using Raspberry 3 b
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_traffic_lights/master/bin/arm7/trafficLight
```
## make it executable
```
chmod +x trafficLight
```

## 4- execute binary to make led blink
```
./trafficLight
```

Congratulation.

Ressource
Follow link below for pin connections 
https://projects.raspberrypi.org/en/projects/physical-computing/11
