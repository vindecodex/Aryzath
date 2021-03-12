package state

import "fmt"

type StartState struct {
	car *Car
}

func (s *StartState) StartEngine() error {
	if s.car.Fuel > 0 {
		fmt.Println("Engine is Running")
		s.car.SetState(s)
		return nil
	}
	s.car.SetState(s.car.NoFuel)
	return fmt.Errorf("Can't Start Fuel is empty")
}

func (s *StartState) StopEngine() error {
	fmt.Println("Engine Stopped")
	s.car.SetState(s.car.OffState)
	return nil
}

func (s *StartState) Run() error {
	fmt.Println("Car is now running")
	s.car.SetState(s.car.RunningState)
	return nil
}

func (s *StartState) Break() {
	fmt.Println("Break is pressed")
	s.car.SetState(s)
}
