package observer

type Notifier interface {
	Register(observer Observer)
	Unregister(observerID int)
	Notify()
}
