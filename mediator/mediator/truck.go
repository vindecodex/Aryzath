package mediator

import "fmt"

type Truck struct {
	mediator Mediator
}

func NewTruck(mediator Mediator) *Truck {
	return &Truck{mediator: mediator}
}

func (t *Truck) UseRoad() {
	if t.mediator.CanUseRoad(t) {
		fmt.Println("Truck is now Using the Road")
		return
	}
	fmt.Println("Truck is waiting someone uses the road")
}

func (t *Truck) UnuseRoad() {
	fmt.Println("Truck is done using the road")
	t.mediator.NotifyForFreeRoad()
}

func (t *Truck) AllowToRoad() {
	fmt.Println("Truck is now allowed to use the road")
	t.UseRoad()
}
