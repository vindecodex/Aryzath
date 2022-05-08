package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/factory_method/tool"
)

func TestFactoryMethod(t *testing.T) {
	t.Run("This will set tools to pencil", func(t *testing.T) {
		got := "Pencil"
		want := tool.SetTool().GetToolName()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
