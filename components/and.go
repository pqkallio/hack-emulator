package components

type And struct {
	a, b                Val
	nand1, nand2, nand3 *Nand
}

func NewAnd() *And {
	return &And{
		&InvalidVal{}, &InvalidVal{},
		NewNand(), NewNand(), NewNand(),
	}
}

func (and *And) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			and.a = opt.val
		case TargetB:
			and.b = opt.val
		}
	}

	nandAB1 := and.nand1.Update(
		UpdateOpts{TargetA, and.a},
		UpdateOpts{TargetB, and.b},
	)

	nandAB2 := and.nand2.Update(
		UpdateOpts{TargetA, and.a},
		UpdateOpts{TargetB, and.b},
	)

	return and.nand3.Update(
		UpdateOpts{TargetA, nandAB1},
		UpdateOpts{TargetB, nandAB2},
	)
}
