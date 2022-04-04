package device

type Light struct{}

func (l Light) On() string {
	return "ON"
}

func (l Light) Off() string {
	return "OFF"
}
