package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/chain_of_responsibility_01/steps"
)

func TestChainOfResponsibility01(t *testing.T) {

	t.Run("car speed should increase by 2", func(t *testing.T) {

		car := steps.NewCar(10)
		mainSteps := steps.NewMainSteps(car)

		carDefaultSpeed := car.GetSpeed()

		mainSteps.SetNext(steps.NewIncreaseSpeed(car))
		mainSteps.SetNext(steps.NewIncreaseSpeed(car))
		mainSteps.Execute()

		got := car.GetSpeed()
		want := carDefaultSpeed + 2

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("car should double its speed", func(t *testing.T) {

		car := steps.NewCar(10)
		mainSteps := steps.NewMainSteps(car)

		carDefaultSpeed := car.GetSpeed()

		mainSteps.SetNext(steps.NewDoubleSpeed(car))
		mainSteps.Execute()

		got := car.GetSpeed()
		want := carDefaultSpeed * 2

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
