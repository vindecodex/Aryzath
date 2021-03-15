package visitor

type Vehicle interface {
	GetVehicleType()
	Accept(Visitor)
}
