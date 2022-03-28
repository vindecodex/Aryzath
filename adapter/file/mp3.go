package file

type Mp3 struct{}

func (Mp3) PlayAudio() string {
	return "Playing From Mp3 File"
}
