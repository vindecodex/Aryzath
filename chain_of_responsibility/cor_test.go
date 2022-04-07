package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/chain_of_responsibility/steps"
)

func TestChainOfResponsibility(t *testing.T) {

	t.Run("record should all be true", func(t *testing.T) {

		record := &steps.Record{}

		step1 := &steps.Step1{}
		step2 := &steps.Step2{}
		step3 := &steps.Step3{}

		step1.SetNext(step2)
		step2.SetNext(step3)

		step1.Execute(record)

		want := true

		if record.StepOneExecuted != want {
			t.Errorf("got %t want %t", record.StepOneExecuted, want)
		}

		if record.StepTwoExecuted != want {
			t.Errorf("got %t want %t", record.StepTwoExecuted, want)
		}

		if record.StepThreeExecuted != want {
			t.Errorf("got %t want %t", record.StepThreeExecuted, want)
		}

	})
}
