package state

import "fmt"

type RunningState struct {
	car *Car
}

func (s *RunningState) StartEngine() error {
	if s.car.Fuel > 0 {
		fmt.Println("Engine already Started")
		return nil
	}
	return fmt.Errorf("Can't Start Fuel is empty")
}

func (s *RunningState) StopEngine() error {
	return fmt.Errorf("Engine cannot be stop while running")
}

func (s *RunningState) Run() error {
	fmt.Println("Car is currently running")
	return nil
}

func (s *RunningState) Break() {
	fmt.Println("Break is pressed")
	s.car.SetState(s.car.StartState)
}
