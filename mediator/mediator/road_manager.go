package mediator

type RoadManager struct {
	isRoadFree bool
	vehicle    []Vehicle
}

func NewRoadManager() *RoadManager {
	return &RoadManager{isRoadFree: true}
}

func (r *RoadManager) CanUseRoad(vehicle Vehicle) bool {
	if r.isRoadFree {
		r.isRoadFree = false
		return true
	}
	r.vehicle = append(r.vehicle, vehicle)
	return false
}

func (r *RoadManager) NotifyForFreeRoad() {
	if !r.isRoadFree {
		r.isRoadFree = true
	}
	if len(r.vehicle) > 0 {
		newRoadUser := r.vehicle[0]
		r.vehicle = r.vehicle[1:]
		newRoadUser.AllowToRoad()
	}
}
