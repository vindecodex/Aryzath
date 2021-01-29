package main

import (
	"Aryzath/memento/memento"
	"testing"
)

func TestMemento(t *testing.T) {
	orig := memento.Originator{}
	ctaker := memento.Caretaker{}

	orig.SetState("1")
	orig.SetState("2")
	ctaker.Save(orig.CreateSavepoint())

	orig.SetState("3")
	ctaker.Save(orig.CreateSavepoint())

	t.Run("Memento test #1", func(t *testing.T) {

		orig.RetrieveSavepoint(ctaker.Restore(0))
		got := orig.GetState()

		want := "2"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Memento test #2", func(t *testing.T) {

		orig.RetrieveSavepoint(ctaker.Restore(1))
		got := orig.GetState()

		want := "3"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
