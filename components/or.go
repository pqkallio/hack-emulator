package components

type Or struct {
	a    Val
	b    Val
	not1 *Not
	not2 *Not
	not3 *Not
	and  *And
}

func NewOr() *Or {
	return &Or{
		&InvalidVal{}, &InvalidVal{},
		NewNot(), NewNot(), NewNot(),
		NewAnd(),
	}
}

func (or *Or) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			or.a = opt.val
		case TargetB:
			or.b = opt.val
		}
	}

	notA := or.not1.Update(UpdateOpts{TargetIn, or.a})
	notB := or.not2.Update(UpdateOpts{TargetIn, or.b})
	notaAndNotb := or.and.Update(
		UpdateOpts{TargetA, notA},
		UpdateOpts{TargetB, notB},
	)
	return or.not3.Update(UpdateOpts{TargetIn, notaAndNotb})
}
