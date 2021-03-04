package steps

import "fmt"

type Step3 struct {
	next Steps
}

func (s3 *Step3) Execute(r *Record) {
	if r.StepThreeExecuted {
		fmt.Println("Proceeding to next step")
		if s3.next != nil {
			s3.next.Execute(r)
		}
		return
	}
	fmt.Println("Working With Step 3")
	r.StepThreeExecuted = true
	if s3.next != nil {
		s3.next.Execute(r)
	}
}

func (s3 *Step3) SetNext(steps Steps) {
	s3.next = steps
}
