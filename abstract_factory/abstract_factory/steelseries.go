package abstract_factory

import (
	"Aryzath/abstract_factory/device"
)

type Steelseries struct{}

func (Steelseries) CreateMouse() IMouse {
	return &device.SteelseriesMouse{
		Mouse: device.Mouse{
			Color: "Yellow",
		},
	}
}

func (Steelseries) CreateKeyboard() IKeyboard {
	return &device.SteelseriesKeyboard{
		Keyboard: device.Keyboard{
			Color: "Yellow",
		},
	}
}
