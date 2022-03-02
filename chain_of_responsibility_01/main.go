package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/chain_of_responsibility_01/steps"
)

func main() {
	car := steps.NewCar(10)
	fmt.Println("current car speed: ", car.GetSpeed())

	mainSteps := steps.NewMainSteps(car)

	mainSteps.SetNext(steps.NewIncreaseSpeed(car))
	mainSteps.SetNext(steps.NewIncreaseSpeed(car))

	mainSteps.SetNext(steps.NewDoubleSpeed(car))

	mainSteps.Execute()

	fmt.Println("current car speed: ", car.GetSpeed())

}
