package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/facade/device"
)

func TestFacade(t *testing.T) {
	t.Run("This will set radio volume to 5", func(t *testing.T) {
		got := 5
		radio := device.Radio{}

		radio.TurnOn()
		radio.SetVolume(got)
		radio.SetStation("99.5")

		want := radio.GetVolume()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
