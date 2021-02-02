package branding

import (
	"Aryzath/abstract_factory/abstract_factory"
	"Aryzath/abstract_factory/device"
)

type SteelSeries struct{}

func (c *SteelSeries) CreateMouse() abstract_factory.IMouse {
	return &device.SteelSeriesMouse{
		SMouse: &device.Mouse{
			Color: "Yellow",
		},
	}
}

func (c *SteelSeries) CreateKeyboard() abstract_factory.IKeyboard {
	return &device.SteelSeriesKeyboard{
		SKeyboard: &device.Keyboard{
			Color: "Yellow",
		},
	}
}
