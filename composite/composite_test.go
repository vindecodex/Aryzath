package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/composite/composite"
)

func TestComposite(t *testing.T) {

	phone := &composite.Product{Name: "IPhone"}
	keyboard := &composite.Product{Name: "Keyboard"}
	mouse := &composite.Product{Name: "Mouse"}

	box1 := &composite.Box{Name: "Box1"}

	box1.Add(phone)
	box1.Add(keyboard)
	box1.Add(mouse)

	box2 := &composite.Box{Name: "Box2"}
	box2.Add(box1)

	box3 := &composite.Box{Name: "Box3"}
	box3.Add(phone)
	box3.Add(keyboard)

	t.Run("Box2 compartment 0 should contain Box1", func(t *testing.T) {
		got := box1
		want := box2.Component[0]

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Box3 compartment 0 should contain phone", func(t *testing.T) {
		got := box3.Component[0]
		want := phone

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Box3 compartment 1 should contain keyboard", func(t *testing.T) {
		got := box3.Component[1]
		want := keyboard

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
