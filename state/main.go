package main

import (
	"log"

	"github.com/vindecodex/Aryzath/state/state"
)

func main() {
	car := state.NewCar(50)

	err := car.StartEngine()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = car.Run()
	if err != nil {
		log.Fatalf(err.Error())
	}

	car.Break()

	err = car.StopEngine()
	if err != nil {
		log.Fatalf(err.Error())
	}

}
