package abstract_factory

import (
	"Aryzath/abstract_factory/branding"
	"errors"
)

type IBrand interface {
	CreateMouse()
	CreateKeyboard()
}

func GetBrand(brand string) (IBrand, error) {
	switch brand {
	case "razer":
		return branding.Razer{}, nil
	case "steelseries":
		return branding.SteelSeries{}, nil
	default:
		return nil, errors.New("Not supported branding")
	}
}
