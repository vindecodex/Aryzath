package file

import "fmt"

type Mp4 struct{}

func (Mp4) PlayVideo() {
	fmt.Println("Playing From Mp4 File")
}
