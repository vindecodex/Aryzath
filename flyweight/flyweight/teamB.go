package flyweight

type TeamB struct {
	Color string
}

func (t *TeamB) GetColor() string {
	return t.Color
}

func NewTeamB() *TeamB {
	return &TeamB{
		Color: "Red",
	}
}
