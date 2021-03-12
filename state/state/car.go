package state

type Car struct {
	StartState   State
	RunningState State
	OffState     State
	NoFuel       State

	CurrentState State
	Fuel         int
}

func NewCar(fuel int) *Car {
	c := &Car{Fuel: fuel}

	startState := &StartState{c}
	runningState := &RunningState{c}
	offState := &OffState{c}
	noFuel := &NoFuelState{c}

	c.StartState = startState
	c.RunningState = runningState
	c.OffState = offState
	c.NoFuel = noFuel

	c.CurrentState = offState

	return c
}

func (c *Car) SetState(state State) {
	c.CurrentState = state
}

func (c *Car) StartEngine() error {
	return c.CurrentState.StartEngine()
}

func (c *Car) StopEngine() error {
	return c.CurrentState.StopEngine()
}

func (c *Car) Run() error {
	return c.CurrentState.Run()
}

func (c *Car) Break() {
	c.CurrentState.Break()
}
