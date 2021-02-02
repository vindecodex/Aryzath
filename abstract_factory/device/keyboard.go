package device

type Keyboard struct {
	Color string
}

func (c *Keyboard) GetColor() string {
	return c.Color
}
