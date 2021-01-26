package main

import (
	"Aryzath/facade/device"
	"fmt"
)

func main() {
	radio := device.Radio{}
	radio.TurnOn()
	fmt.Println(radio)
	radio.TurnOff()
	fmt.Println(radio)
}
