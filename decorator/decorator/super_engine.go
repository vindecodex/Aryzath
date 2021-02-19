package decorator

import (
	"Aryzath/decorator/component"
)

type SuperEngine struct {
	Car component.Car
}

func (s *SuperEngine) Set() []string {
	car_set := s.Car.Set()
	car_set = append(car_set, "super engine")
	return car_set
}
