package visitor

import "fmt"

type Car struct {
	Color string
}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) GetVehicleType() {
	fmt.Println("Car")
}

func (c *Car) Accept(v Visitor) {
	v.VisitCar(c)
}
