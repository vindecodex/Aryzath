package main

import (
	"Aryzath/decorator/component"
	"Aryzath/decorator/decorator"
	"fmt"
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
