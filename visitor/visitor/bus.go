package visitor

import "fmt"

type Bus struct {
	Color string
}

func NewBus() *Bus {
	return &Bus{}
}

func (b *Bus) GetVehicleType() {
	fmt.Println("Bus")
}

func (b *Bus) Accept(v Visitor) {
	v.VisitBus(b)
}
