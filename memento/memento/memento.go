package memento

type Memento struct {
	State string
}

func (m *Memento) GetState() string {
	return m.State
}
