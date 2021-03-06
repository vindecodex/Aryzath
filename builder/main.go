package main

import (
	"Aryzath/builder/builder"
	"fmt"
)

func main() {
	crystalBuilder, _ := builder.GetBuilder("crystal")
	rubyBuilder, _ := builder.GetBuilder("ruby")

	director := builder.NewDirector(crystalBuilder)

	crystalCyborg := director.BuildCyborg()
	fmt.Println("arms: ", crystalCyborg.LeftArm, " - ", crystalCyborg.RightArm)
	fmt.Println("legs: ", crystalCyborg.Legs)
	fmt.Println("procs:", crystalCyborg.Processor)

	director.SetBuilder(rubyBuilder)

	rubyCyborg := director.BuildCyborg()
	fmt.Println("arms: ", rubyCyborg.LeftArm, " - ", rubyCyborg.RightArm)
	fmt.Println("legs: ", rubyCyborg.Legs)
	fmt.Println("procs:", rubyCyborg.Processor)
}
