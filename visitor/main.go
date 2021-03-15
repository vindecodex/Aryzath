package main

import "Aryzath/visitor/visitor"

func main() {
	car := visitor.NewCar()
	bus := visitor.NewBus()

	colorChanger := &visitor.ColorChanger{}

	car.Accept(colorChanger)
	bus.Accept(colorChanger)

	speedLimiter := &visitor.SpeedLimiter{}

	car.Accept(speedLimiter)
	bus.Accept(speedLimiter)
}
