package visitor

import "fmt"

type SpeedLimiter struct{}

func (c *SpeedLimiter) VisitBus(bus *Bus) {
	fmt.Println("Speed Limiter for Bus")
}

func (c *SpeedLimiter) VisitCar(car *Car) {
	fmt.Println("Speed Limiter for Car")
}
