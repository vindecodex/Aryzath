package visitor

type Visitor interface {
	VisitBus(*Bus)
	VisitCar(*Car)
}
