package abstract_factory

import (
	"Aryzath/abstract_factory/branding"
	"Aryzath/abstract_factory/device"
)

type Razer struct{}

func (Razer) CreateMouse() IMouse {
	return &branding.RazerMouse{
		RMouse: &device.Mouse{
			Color: "Green",
		},
	}
}

func (Razer) CreateKeyboard() IKeyboard {
	return &branding.RazerKeyboard{
		RKeyboard: &device.Keyboard{
			Color: "Green",
		},
	}
}
