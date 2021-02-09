package builder

import "Aryzath/builder/builder/product"

type Director struct {
	builder IBuilder
}

func NewDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(builder IBuilder) {
	d.builder = builder
}

func (d *Director) BuildCyborg() product.Cyborg {
	d.builder.SetLegs()
	d.builder.SetLeftArm()
	d.builder.SetRightArm()
	d.builder.SetProcessor()
	return d.builder.BuildCyborg()
}
