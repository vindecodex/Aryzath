package player

import "Aryzath/adapter/file"

type AudioPlayer struct{}

func (AudioPlayer) PlayAudioFromFile(m file.Media) {
	m.PlayAudio()
}
