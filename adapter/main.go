package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/adapter/adapter"
	"github.com/vindecodex/Aryzath/adapter/file"
	"github.com/vindecodex/Aryzath/adapter/player"
)

func main() {

	audioPlayer := player.AudioPlayer{}

	mp3 := &file.Mp3{}
	fmt.Println(audioPlayer.PlayAudioFromFile(mp3))

	mp4 := &adapter.Mp4Adapter{
		Mp4: &file.Mp4{},
	}
	fmt.Println(audioPlayer.PlayAudioFromFile(mp4))
}
