package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/bridge/abstraction_heirarchy"
	"github.com/vindecodex/Aryzath/bridge/implementation_heirarchy"
)

func TestAdapter(t *testing.T) {

	eclipse := implementation_heirarchy.Eclipse{}
	netbeans := implementation_heirarchy.Netbeans{}

	java_file := &abstraction_heirarchy.Java{}
	cpp_file := &abstraction_heirarchy.Cpp{}

	t.Run("run java file on eclipse", func(t *testing.T) {

		java_file.SetIde(eclipse)
		got := java_file.Compile()
		want := "Java Compiling using Eclipse"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("run java file on netbeans", func(t *testing.T) {

		java_file.SetIde(netbeans)
		got := java_file.Compile()
		want := "Java Compiling using Netbeans"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("run cpp file on eclipse", func(t *testing.T) {

		cpp_file.SetIde(eclipse)
		got := cpp_file.Compile()
		want := "Cpp Compiling using Eclipse"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("run cpp file on netbeans", func(t *testing.T) {

		cpp_file.SetIde(netbeans)
		got := cpp_file.Compile()
		want := "Cpp Compiling using Netbeans"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
