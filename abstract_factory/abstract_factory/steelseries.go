package abstract_factory

import (
	"Aryzath/abstract_factory/device"
)

type Steelseries struct{}

func (Steelseries) CreateMouse() IMouse {
	return &device.SteelseriesMouse{
		SSMouse: &device.Mouse{
			Color: "Yellow",
		},
	}
}

func (Steelseries) CreateKeyboard() IKeyboard {
	return &device.SteelseriesKeyboard{
		SSKeyboard: &device.Keyboard{
			Color: "Yellow",
		},
	}
}
