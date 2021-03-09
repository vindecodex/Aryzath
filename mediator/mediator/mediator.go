package mediator

type Mediator interface {
	CanUseRoad(Vehicle) bool
	NotifyForFreeRoad()
}
