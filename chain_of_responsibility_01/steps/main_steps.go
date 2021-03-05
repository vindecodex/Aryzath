package steps

type MainSteps struct {
	Car  *Car
	Next Steps
}

func NewMainSteps(car *Car) *MainSteps {
	return &MainSteps{Car: car}
}

func (m *MainSteps) Execute() {
	if m.Next != nil {
		m.Next.Execute()
	}
}

func (m *MainSteps) SetNext(steps Steps) {
	if m.Next != nil {
		m.Next.SetNext(steps)
	} else {
		m.Next = steps
	}
}
