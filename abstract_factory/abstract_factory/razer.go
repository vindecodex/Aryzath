package branding

import (
	"Aryzath/abstract_factory/device"
)

type Razer struct{}

func (Razer) CreateMouse() IMouse {
	return &device.RazerMouse{
		RMouse: &device.Mouse{
			Color: "Green",
		},
	}
}

func (Razer) CreateKeyboard() IKeyboard {
	return &device.RazerKeyboard{
		RKeyboard: &device.Keyboard{
			Color: "Green",
		},
	}
}
