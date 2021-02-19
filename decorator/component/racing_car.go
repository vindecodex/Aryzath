package component

type RacingCar struct{}

func (r *RacingCar) Set() []string {
	default_set := make([]string, 0)
	default_set = append(default_set, "default_tire", "default_fluid", "default_pedals")
	return default_set
}
