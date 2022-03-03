package main

import (
	"github.com/vindecodex/Aryzath/command/command"
	"github.com/vindecodex/Aryzath/command/device"
)

func main() {
	lon := &device.LightOn{}  // Holds On Method
	lof := &device.LightOff{} // Holds Off Method

	do := command.Do{}

	do.SetCommand(lon)
	do.It() // do On Method

	do.SetCommand(lof)
	do.It() // do Off Method
}
