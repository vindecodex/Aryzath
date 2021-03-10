package observer

type Observer interface {
	GetNotified(product string)
	GetId() int
}
