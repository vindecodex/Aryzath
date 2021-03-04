package steps

type Steps interface {
	Execute(*Record)
	SetNext(Steps)
}
