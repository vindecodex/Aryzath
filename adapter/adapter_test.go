package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/adapter/adapter"
	"github.com/vindecodex/Aryzath/adapter/file"
	"github.com/vindecodex/Aryzath/adapter/player"
)

func TestAdapter(t *testing.T) {

	audioPlayer := player.AudioPlayer{}

	t.Run("Adapter test: mp3", func(t *testing.T) {

		mp3 := &file.Mp3{}
		playMp3 := audioPlayer.PlayAudioFromFile(mp3)

		want := "Playing From Mp3 File"

		if playMp3 != want {
			t.Errorf("got %q want %q", playMp3, want)
		}
	})

	t.Run("Adapter test: mp4", func(t *testing.T) {

		// use the adapter to support playing mp4 file types
		mp4 := &adapter.Mp4Adapter{
			&file.Mp4{},
		}

		playMp4 := audioPlayer.PlayAudioFromFile(mp4)

		want := "Playing From Mp4 File"

		if playMp4 != want {
			t.Errorf("got %q want %q", playMp4, want)
		}

	})

}
