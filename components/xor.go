package components

type Xor struct {
	a    Val
	b    Val
	not1 *Not
	not2 *Not
	and1 *And
	and2 *And
	or   *Or
}

func NewXor() *Xor {
	return &Xor{
		&InvalidVal{}, &InvalidVal{},
		NewNot(), NewNot(),
		NewAnd(), NewAnd(),
		NewOr(),
	}
}

func (xor *Xor) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			xor.a = opt.val
		case TargetB:
			xor.b = opt.val
		}
	}

	nota := xor.not1.Update(UpdateOpts{TargetIn, xor.a})
	notb := xor.not1.Update(UpdateOpts{TargetIn, xor.b})
	aAndNotb := xor.and1.Update(
		UpdateOpts{TargetA, xor.a},
		UpdateOpts{TargetB, notb},
	)
	notaAndb := xor.and2.Update(
		UpdateOpts{TargetA, nota},
		UpdateOpts{TargetB, xor.b},
	)

	return xor.or.Update(
		UpdateOpts{TargetA, aAndNotb},
		UpdateOpts{TargetB, notaAndb},
	)
}
