package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/iterator/iterator"
)

func TestIterator(t *testing.T) {
	t.Run("", func(t *testing.T) {
		data_a := &iterator.Data{}
		data_b := &iterator.Data{}

		data_a.SetData("ABC")
		data_b.SetData("XYZ")

		want := iterator.DataCollection{Data: []*iterator.Data{
			data_a,
			data_b,
		}}

		got := []string{
			"ABC",
			"XYZ",
		}

		for i, _want := range want.Data {
			if got[i] != _want.GetData() {
				t.Errorf("got: %q, want: %q", got[i], _want.GetData())
			}
		}
	})
}
