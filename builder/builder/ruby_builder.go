package builder

import "github.com/vindecodex/Aryzath/builder/builder/product"

type RubyBuilder struct {
	RightArm  string
	LeftArm   string
	Legs      string
	Processor string
}

func (r *RubyBuilder) SetRightArm() {
	r.RightArm = "Right Ruby Arm"
}
func (r *RubyBuilder) SetLeftArm() {
	r.LeftArm = "Left Ruby Arm"
}
func (r *RubyBuilder) SetLegs() {
	r.Legs = "Ruby Legs"
}
func (r *RubyBuilder) SetProcessor() {
	r.Processor = "AMD Ryzeen"
}
func (r *RubyBuilder) BuildCyborg() product.Cyborg {
	return product.Cyborg{
		RightArm:  r.RightArm,
		LeftArm:   r.LeftArm,
		Legs:      r.Legs,
		Processor: r.Processor,
	}
}
