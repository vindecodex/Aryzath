package player

import "github.com/vindecodex/Aryzath/adapter/file"

type AudioPlayer struct{}

func (AudioPlayer) PlayAudioFromFile(m file.Media) string {
	return m.PlayAudio()
}
