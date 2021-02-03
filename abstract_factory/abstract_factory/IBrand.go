package abstract_factory

import (
	"errors"
)

type IBrand interface {
	CreateMouse() IMouse
	CreateKeyboard() IKeyboard
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
