package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/command/command"
	"github.com/vindecodex/Aryzath/command/device"
)

func main() {
	lon := &device.LightOn{}  // Holds On Method
	lof := &device.LightOff{} // Holds Off Method

	do := command.Do{}

	do.SetCommand(lon)
	fmt.Println(do.It()) // do On Method

	do.SetCommand(lof)
	fmt.Println(do.It()) // do Off Method
}
