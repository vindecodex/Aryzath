package abstract_factory

import "github.com/vindecodex/Aryzath/abstract_factory/device"

type Razer struct{}

func (Razer) CreateMouse() IMouse {
	return &device.RazerMouse{
		Mouse: device.Mouse{
			Color: "Green",
		},
	}
}

func (Razer) CreateKeyboard() IKeyboard {
	return &device.RazerKeyboard{
		Keyboard: device.Keyboard{
			Color: "Green",
		},
	}
}
