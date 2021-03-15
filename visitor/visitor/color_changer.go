package visitor

import "fmt"

type ColorChanger struct{}

func (c *ColorChanger) VisitBus(bus *Bus) {
	fmt.Println("Color Changer for Bus")
	bus.Color = "RED"
}

func (c *ColorChanger) VisitCar(car *Car) {
	fmt.Println("Color Changer for Car")
	car.Color = "BLUE"
}
