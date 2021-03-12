package strategy

import (
	"fmt"
)

type StrategyB struct{}

func (s *StrategyB) Solve(sq *SquareRoot) {
	val := sq.GetValue() * sq.GetValue()
	fmt.Println("Using Strategy B", val)
}
