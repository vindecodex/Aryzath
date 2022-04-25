package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/decorator/component"
	"github.com/vindecodex/Aryzath/decorator/decorator"
)

func TestDecorator(t *testing.T) {
	racingCar := &component.RacingCar{}

	racingCarWithSuperEngine := &decorator.SuperEngine{racingCar}
	racingCarWithSuperExhaust := &decorator.SuperExhaust{racingCar}
	allSuperRacingCar := &decorator.SuperExhaust{racingCarWithSuperEngine}

	t.Run("should contain default attributes", func(t *testing.T) {
		got := []string{"default_tire", "default_fluid", "default_pedals"}
		want := racingCar.Set()

		for i, attr := range want {
			if got[i] != attr {
				t.Errorf("got %q, want %q", got[i], attr)
			}
		}
	})

	t.Run("should contain super engine attributes", func(t *testing.T) {
		got := "super engine"
		want := racingCarWithSuperEngine.Set()
		lastAttribute := len(want) - 1

		if got != want[lastAttribute] {
			t.Errorf("got %q, want %q", got, want[lastAttribute])
		}
	})

	t.Run("should contain super exhaust attributes", func(t *testing.T) {
		got := "super exhaust"
		want := racingCarWithSuperExhaust.Set()
		lastAttribute := len(want) - 1

		if got != want[lastAttribute] {
			t.Errorf("got %q, want %q", got, want[lastAttribute])
		}
	})

	t.Run("should contain super engine and super exhaust", func(t *testing.T) {
		gotSuperExhaust := "super exhaust"
		gotSuperEngine := "super engine"

		wantAttributes := allSuperRacingCar.Set()
		attrLength := len(wantAttributes)

		if gotSuperEngine != wantAttributes[attrLength-2] {
			t.Errorf("got %q, want %q", gotSuperEngine, wantAttributes[attrLength-2])
		}

		if gotSuperExhaust != wantAttributes[attrLength-1] {
			t.Errorf("got %q, want %q", gotSuperExhaust, wantAttributes[attrLength-1])
		}

	})

}
