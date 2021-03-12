package strategy

type SquareRoot struct {
	Value float64
	Strat Strategy
}

func NewSquareRoot(val float64) *SquareRoot {
	return &SquareRoot{Value: val}
}

func (s *SquareRoot) Solve() {
	s.Strat.Solve(s)
}

func (s *SquareRoot) GetValue() float64 {
	return s.Value
}

func (s *SquareRoot) SetStrategy(strategy Strategy) {
	s.Strat = strategy
}
