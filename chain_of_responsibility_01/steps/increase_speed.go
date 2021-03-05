package steps

import "fmt"

type IncreaseSpeed struct {
	MainSteps
}

func NewIncreaseSpeed(c *Car) *IncreaseSpeed {
	return &IncreaseSpeed{MainSteps{Car: c}}
}

func (i *IncreaseSpeed) Execute() {
	fmt.Println("Increasing car speed by 1")
	currentSpeed := i.MainSteps.Car.GetSpeed()
	i.MainSteps.Car.SetSpeed(currentSpeed + 1)
	i.MainSteps.Execute()
}
