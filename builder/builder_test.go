package main

import (
	"testing"

	"github.com/vindecodex/Aryzath/builder/builder"
)

func TestBuilder(t *testing.T) {

	t.Run("test crystal builder", func(t *testing.T) {

		crystalBuilder, _ := builder.GetBuilder("crystal")
		director := builder.NewDirector(crystalBuilder)
		crystalCyborg := director.BuildCyborg()

		gotLeftArm := crystalCyborg.LeftArm
		wantLeftArm := "Left Crystal Arm"

		if gotLeftArm != wantLeftArm {
			t.Errorf("got %q want %q", gotLeftArm, wantLeftArm)
		}

		gotRightArm := crystalCyborg.RightArm
		wantRightArm := "Right Crystal Arm"

		if gotLeftArm != wantLeftArm {
			t.Errorf("got %q want %q", gotRightArm, wantRightArm)
		}

	})

	t.Run("test ruby builder", func(t *testing.T) {

		rubyBuilder, _ := builder.GetBuilder("ruby")
		director := builder.NewDirector(rubyBuilder)
		rubyCyborg := director.BuildCyborg()

		gotLeftArm := rubyCyborg.LeftArm
		wantLeftArm := "Left Ruby Arm"

		if gotLeftArm != wantLeftArm {
			t.Errorf("got %q want %q", gotLeftArm, wantLeftArm)
		}

		gotRightArm := rubyCyborg.RightArm
		wantRightArm := "Right Crystal Arm"

		if gotLeftArm != wantLeftArm {
			t.Errorf("got %q want %q", gotRightArm, wantRightArm)
		}

	})
}
