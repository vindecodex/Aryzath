package strategy

import (
	"fmt"
)

type StrategyC struct{}

func (s *StrategyC) Solve(sq *SquareRoot) {
	val := sq.GetValue() * sq.GetValue()
	fmt.Println("Using Strategy C", val)
}
