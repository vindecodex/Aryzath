package steps

import "fmt"

type Step2 struct {
	next Steps
}

func (s2 *Step2) Execute(r *Record) {
	if r.StepTwoExecuted {
		fmt.Println("Proceeding to next step")
		s2.next.Execute(r)
		return
	}
	fmt.Println("Working With Step 2")
	r.StepTwoExecuted = true
	s2.next.Execute(r)
}

func (s2 *Step2) SetNext(steps Steps) {
	s2.next = steps
}
