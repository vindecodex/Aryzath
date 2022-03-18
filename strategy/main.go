package main

import "github.com/vindecodex/Aryzath/strategy/strategy"

func main() {
	squareRoot := strategy.NewSquareRoot(5)

	stratA := &strategy.StrategyA{}
	stratB := &strategy.StrategyB{}
	stratC := &strategy.StrategyC{}

	squareRoot.SetStrategy(stratA)
	squareRoot.Solve()

	squareRoot.SetStrategy(stratB)
	squareRoot.Solve()

	squareRoot.SetStrategy(stratC)
	squareRoot.Solve()

}
