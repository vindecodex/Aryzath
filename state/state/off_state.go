package state

import "fmt"

type OffState struct {
	car *Car
}

func (s *OffState) StartEngine() error {
	if s.car.Fuel > 0 {
		fmt.Println("Engine is Running")
		s.car.SetState(s.car.StartState)
		return nil
	}
	return fmt.Errorf("Can't Start Fuel is empty")
}

func (s *OffState) StopEngine() error {
	fmt.Println("Engine already turned off")
	return nil
}

func (s *OffState) Run() error {
	return fmt.Errorf("Engine Should be started")
}

func (s *OffState) Break() {
	fmt.Println("Break is pressed")
}
