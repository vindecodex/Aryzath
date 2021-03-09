package mediator

import "fmt"

type Car struct {
	mediator Mediator
}

func NewCar(mediator Mediator) *Car {
	return &Car{mediator: mediator}
}

func (c *Car) UseRoad() {
	if c.mediator.CanUseRoad(c) {
		fmt.Println("Car is now Using the Road")
		return
	}
	fmt.Println("Car is waiting someone uses the road")
}

func (c *Car) UnuseRoad() {
	fmt.Println("Car is done using the road")
	c.mediator.NotifyForFreeRoad()
}

func (c *Car) AllowToRoad() {
	fmt.Println("Car is now allowed to use the road")
	c.UseRoad()
}
