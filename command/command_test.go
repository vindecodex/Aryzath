package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/command/command"
	"github.com/vindecodex/Aryzath/command/device"
)

func TestCommand(t *testing.T) {

	do := command.Do{}

	t.Run("should turn ON the light", func(t *testing.T) {

		lon := &device.LightOn{}
		do.SetCommand(lon)
		got := do.It()
		want := "ON"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("should turn OFF the light", func(t *testing.T) {

		lof := &device.LightOff{}
		do.SetCommand(lof)
		got := do.It()
		want := "OFF"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
