package device

type Radio struct {
	volume  int
	station string
	status  string
}

/*
Facade is an abstraction of complex methods and combine them to
generate a simple easy to use function.

Now any client want to use Radio Class can turn on the radio
easily without setting up volume and station on there own.
*/

func (r *Radio) TurnOn() {
	r.setStatus("ON")
	r.SetVolume(10)
	r.SetStation("107.5")
}

func (r *Radio) TurnOff() {
	r.setStatus("OFF")
	r.SetVolume(0)
	r.SetStation("")
}

func (r *Radio) SetVolume(v int) {
	r.volume = v
}

func (r *Radio) SetStation(s string) {
	r.station = s
}

func (r *Radio) setStatus(s string) {
	r.status = s
}
