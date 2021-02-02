package device

type Mouse struct {
	Color string
}

func (c *Mouse) GetColor() string {
	return c.Color
}
