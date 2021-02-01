package file

import "fmt"

type Mp3 struct{}

func (Mp3) PlayAudio() {
	fmt.Println("Playing From Mp3 File")
}
