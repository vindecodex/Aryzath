package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/decorator/component"
	"github.com/vindecodex/Aryzath/decorator/decorator"
)

func main() {
	racingCar := &component.RacingCar{}
	fmt.Println(racingCar.Set())

	racingCarWithSuperEngine := &decorator.SuperEngine{racingCar}
	fmt.Println(racingCarWithSuperEngine.Set())

	racingCarWithSuperExhaust := &decorator.SuperExhaust{racingCar}
	fmt.Println(racingCarWithSuperExhaust.Set())

	allSuperRacingCar := &decorator.SuperExhaust{racingCarWithSuperEngine}
	fmt.Println(allSuperRacingCar.Set())
}
