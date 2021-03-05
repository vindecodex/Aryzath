package steps

import "fmt"

type DoubleSpeed struct {
	MainSteps
}

func NewDoubleSpeed(c *Car) *DoubleSpeed {
	return &DoubleSpeed{MainSteps{Car: c}}
}

func (i *DoubleSpeed) Execute() {
	fmt.Println("Double the Speed of car")
	currentSpeed := i.MainSteps.Car.GetSpeed()
	i.MainSteps.Car.SetSpeed(currentSpeed * 2)
	i.MainSteps.Execute()
}
