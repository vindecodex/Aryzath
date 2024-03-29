package main

import "github.com/vindecodex/Aryzath/mediator/mediator"

func main() {
	roadManager := mediator.NewRoadManager()

	car := mediator.NewCar(roadManager)
	truck := mediator.NewTruck(roadManager)

	car.UseRoad()
	truck.UseRoad()
	car.UnuseRoad()
}
