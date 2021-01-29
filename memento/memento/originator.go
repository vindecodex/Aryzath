package memento

type Originator struct {
	State string
}

func (o *Originator) SetState(s string) {
	o.State = s
}

func (o *Originator) GetState() string {
	return o.State
}

func (o *Originator) CreateSavepoint() *Memento {
	return &Memento{
		State: o.State,
	}
}

func (o *Originator) RetrieveSavepoint(m *Memento) {
	o.State = m.State
}
