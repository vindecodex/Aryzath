package main

import (
	"Aryzath/adapter/adapter"
	"Aryzath/adapter/file"
	"Aryzath/adapter/player"
)

func main() {

	audioPlayer := player.AudioPlayer{}

	mp3 := &file.Mp3{}
	audioPlayer.PlayAudioFromFile(mp3)

	mp4 := &adapter.Mp4Adapter{
		Mp4: &file.Mp4{},
	}
	audioPlayer.PlayAudioFromFile(mp4)
}
