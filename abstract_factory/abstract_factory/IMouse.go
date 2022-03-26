package abstract_factory

type IMouse interface {
	GetColor() string
}

func GetMouseDetails(mouse IMouse) string {
	return mouse.GetColor()
}
