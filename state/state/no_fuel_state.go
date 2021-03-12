package state

import "fmt"

type NoFuelState struct {
	car *Car
}

func (s *NoFuelState) StartEngine() error {
	return fmt.Errorf("Can't Start Fuel is empty")
}

func (s *NoFuelState) StopEngine() error {
	fmt.Println("Engine already turned off")
	return nil
}

func (s *NoFuelState) Run() error {
	return fmt.Errorf("Engine Should be started")
}

func (s *NoFuelState) Break() {
	fmt.Println("Break is pressed")
}
