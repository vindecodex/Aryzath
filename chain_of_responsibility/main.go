package main

import "Aryzath/chain_of_responsibility/steps"

func main() {

	record := &steps.Record{}

	step1 := &steps.Step1{}
	step2 := &steps.Step2{}
	step3 := &steps.Step3{}

	step1.SetNext(step2)
	step2.SetNext(step3)

	step1.Execute(record)
}
