package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/abstract_factory/abstract_factory"
)

func TestAbstractFactory(t *testing.T) {

	// Talks to the factory which returns an interface
	razer, _ := abstract_factory.GetBrand("razer")
	steelseries, _ := abstract_factory.GetBrand("steelseries")

	rmouse := razer.CreateMouse()
	rkeyboard := razer.CreateKeyboard()

	smouse := steelseries.CreateMouse()
	skeyboard := steelseries.CreateKeyboard()

	t.Run("Abstract Factory test #1: Razer", func(t *testing.T) {

		rmouse_got := abstract_factory.GetMouseDetails(rmouse)
		rkeyboard_got := abstract_factory.GetKeyboardDetails(rkeyboard)

		want := "Green"

		if rmouse_got != want {
			t.Errorf("got %q want %q", rmouse_got, want)
		}

		if rkeyboard_got != want {
			t.Errorf("got %q want %q", rkeyboard_got, want)
		}

	})

	t.Run("Abstract Factory test #2: Steelseries", func(t *testing.T) {

		smouse_got := abstract_factory.GetMouseDetails(smouse)
		skeyboard_got := abstract_factory.GetKeyboardDetails(skeyboard)

		want := "Yellow"

		if smouse_got != want {
			t.Errorf("got %q want %q", smouse_got, want)
		}

		if skeyboard_got != want {
			t.Errorf("got %q want %q", skeyboard_got, want)
		}
	})

}
