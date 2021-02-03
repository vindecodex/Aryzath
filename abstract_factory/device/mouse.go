package device

type Mouse struct {
	Color string
}

func (m *Mouse) GetColor() string {
	return m.Color
}
