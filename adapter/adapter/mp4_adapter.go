package adapter

import (
	"Aryzath/adapter/file"
	"fmt"
)

type Mp4Adapter struct {
	Mp4 *file.Mp4
}

func (mp4Adapter *Mp4Adapter) PlayAudio() { // since Mp4 doesn't have PlayAudio we will be the one to implement that method
	fmt.Println("Initializing Adapter...")
	mp4Adapter.Mp4.PlayVideo()
}
