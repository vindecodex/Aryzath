package adapter

import (
	"Aryzath/adapter/file"
	"fmt"
)

type Mp4Adapter struct {
	Mp4 *file.Mp4
}

func (mp4Adapter *Mp4Adapter) PlayAudio() {
	fmt.Println("Initializing Adapter...")
	mp4Adapter.Mp4.PlayVideo()
}
