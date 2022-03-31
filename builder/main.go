package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/builder/builder"
)

func main() {
	crystalBuilder, _ := builder.GetBuilder("crystal")
	rubyBuilder, _ := builder.GetBuilder("ruby")

	director := builder.NewDirector(crystalBuilder)

	crystalCyborg := director.BuildCyborg()
	fmt.Println("arms: ", crystalCyborg.LeftArm, " and ", crystalCyborg.RightArm)
	fmt.Println("legs: ", crystalCyborg.Legs)
	fmt.Println("procs:", crystalCyborg.Processor)

	director.SetBuilder(rubyBuilder)

	rubyCyborg := director.BuildCyborg()
	fmt.Println("arms: ", rubyCyborg.LeftArm, " and ", rubyCyborg.RightArm)
	fmt.Println("legs: ", rubyCyborg.Legs)
	fmt.Println("procs:", rubyCyborg.Processor)
}
