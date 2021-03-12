package strategy

import (
	"fmt"
)

type StrategyA struct{}

func (s *StrategyA) Solve(sq *SquareRoot) {
	val := sq.GetValue() * sq.GetValue()
	fmt.Println("Using Strategy A", val)
}
