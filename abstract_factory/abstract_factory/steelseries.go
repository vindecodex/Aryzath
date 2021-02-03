package abstract_factory

import (
	"Aryzath/abstract_factory/branding"
	"Aryzath/abstract_factory/device"
)

type Steelseries struct{}

func (Steelseries) CreateMouse() IMouse {
	return &branding.SteelseriesMouse{
		SSMouse: &device.Mouse{
			Color: "Yellow",
		},
	}
}

func (Steelseries) CreateKeyboard() IKeyboard {
	return &branding.SteelseriesKeyboard{
		SSKeyboard: &device.Keyboard{
			Color: "Yellow",
		},
	}
}
