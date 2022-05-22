package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/flyweight/flyweight"
	"github.com/vindecodex/Aryzath/flyweight/game"
)

func TestFlyweight(t *testing.T) {
	t.Run("This will return team player of index n", func(t *testing.T) {
		game := &game.Game{}
		game.AddPlayer(flyweight.TeamTypeA)
		got := "Blue"
		want := game.ShowPlayer(0)

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}

	})
}
