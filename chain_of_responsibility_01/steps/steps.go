package steps

type Steps interface {
	Execute()
	SetNext(Steps)
}
