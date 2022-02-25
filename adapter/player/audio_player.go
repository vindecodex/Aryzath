package player

import "github.com/vindecodex/Aryzath/adapter/file"

type AudioPlayer struct{}

func (AudioPlayer) PlayAudioFromFile(m file.Media) {
	m.PlayAudio()
}
