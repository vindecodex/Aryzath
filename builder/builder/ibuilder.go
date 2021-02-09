package builder

import (
	"Aryzath/builder/builder/product"
	"errors"
)

type IBuilder interface {
	SetRightArm()
	SetLeftArm()
	SetLegs()
	SetProcessor()
	BuildCyborg() product.Cyborg
}

func GetBuilder(builder string) (IBuilder, error) {
	switch builder {
	case "crystal":
		return &CrystalBuilder{}, nil
	case "ruby":
		return &RubyBuilder{}, nil
	default:
		return nil, errors.New("Builder not supported")
	}
}
