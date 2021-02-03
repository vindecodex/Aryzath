package device

type Keyboard struct {
	Color string
}

func (k *Keyboard) GetColor() string {
	return k.Color
}
