package abstract_factory

import (
	"errors"
)

type IBrand interface {
	CreateMouse()
	CreateKeyboard()
}

func GetBrand(brand string) (IBrand, error) {
	switch brand {
	case "razer":
		return &Razer{}, nil
	case "steelseries":
		return &Steelseries{}, nil
	default:
		return nil, errors.New("Brand not Exist")
	}
}
