package decorator

import "github.com/vindecodex/Aryzath/decorator/component"

type SuperExhaust struct {
	Car component.Car
}

func (s *SuperExhaust) Set() []string {
	car_set := s.Car.Set()
	car_set = append(car_set, "super exhaust")
	return car_set
}
