package flyweight

type TeamA struct {
	Color string
}

func (t *TeamA) GetColor() string {
	return t.Color
}

func NewTeamA() *TeamA {
	return &TeamA{
		Color: "Blue",
	}
}
