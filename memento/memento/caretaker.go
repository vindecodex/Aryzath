package memento

type Caretaker struct {
	m []*Memento
}

func (c *Caretaker) Save(m *Memento) {
	c.m = append(c.m, m)
}

func (c *Caretaker) Restore(index int) *Memento {
	return c.m[index]
}
