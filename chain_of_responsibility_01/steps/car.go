package steps

type Car struct {
	speed int
}

func NewCar(speed int) *Car {
	return &Car{speed}
}

func (c *Car) SetSpeed(speed int) {
	c.speed = speed
}

func (c *Car) GetSpeed() int {
	return c.speed
}
