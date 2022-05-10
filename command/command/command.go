package command

type Command interface {
	Execute() string
}
