package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/facade/device"
)

func main() {
	radio := device.Radio{}
	radio.TurnOn()
	fmt.Println(radio)
	radio.TurnOff()
	fmt.Println(radio)
}
