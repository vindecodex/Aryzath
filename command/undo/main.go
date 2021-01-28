package main

import (
	"Aryzath/command/undo/command"
	"Aryzath/command/undo/device"
)

func main() {
	do := command.Do{}
	printCommand := device.PrinterCommand{}

	printCommand.P.Content = "AbC"
	do.SetCommand(printCommand)

	printCommand.P.Content = "213"
	do.SetCommand(printCommand)

	printCommand.P.Content = "214"
	do.SetCommand(printCommand)

	do.It()   // outputs 214
	do.Undo() // outputs 213
	do.Undo() // outputs AbC
}
