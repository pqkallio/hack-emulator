package components

type Not struct {
	nand *Nand
}

func NewNot() *Not {
	return &Not{NewNand()}
}

func (not *Not) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		return not.nand.Update(
			UpdateOpts{TargetA, opt.val},
			UpdateOpts{TargetB, opt.val},
		)
	}

	return &InvalidVal{}
}
