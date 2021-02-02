package branding

import (
	"Aryzath/abstract_factory/abstract_factory"
	"Aryzath/abstract_factory/device"
)

type Razer struct{}

func (c *Razer) CreateMouse() abstract_factory.IMouse {
	return &device.RazerMouse{
		RMouse: &device.Mouse{
			Color: "Green",
		},
	}
}

func (c *Razer) CreateKeyboard() abstract_factory.IKeyboard {
	return device.RazerKeyboard{
		RKeyboard: &device.Keyboard{
			Color: "Green",
		},
	}
}
