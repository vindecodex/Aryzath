package adapter

import (
	"fmt"

	"github.com/vindecodex/Aryzath/adapter/file"
)

type Mp4Adapter struct {
	Mp4 *file.Mp4
}

func (mp4Adapter *Mp4Adapter) PlayAudio() { // since Mp4 doesn't have PlayAudio we will be the one to implement that method
	fmt.Println("Initializing Adapter...")
	mp4Adapter.Mp4.PlayVideo()
}
