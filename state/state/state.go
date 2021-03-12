package state

type State interface {
	StartEngine() error
	StopEngine() error
	Run() error
	Break()
}
