package steps

import "fmt"

type Step1 struct {
	next Steps
}

func (s1 *Step1) Execute(r *Record) {
	if r.StepOneExecuted {
		fmt.Println("Proceeding to next step")
		s1.next.Execute(r)
		return
	}
	fmt.Println("Working With Step 1")
	r.StepOneExecuted = true
	s1.next.Execute(r)
}

func (s1 *Step1) SetNext(steps Steps) {
	s1.next = steps
}
