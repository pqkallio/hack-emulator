package components

type HalfAdder struct {
	a, b Val
	xor  *Xor
	and  *And
}

func NewHalfAdder() *HalfAdder {
	return &HalfAdder{
		&InvalidVal{}, &InvalidVal{},
		NewXor(), NewAnd(),
	}
}

func (a *HalfAdder) Update(opts ...UpdateOpts) (Val, Val) {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			a.a = opt.val
		case TargetB:
			a.b = opt.val
		}
	}

	return a.xor.Update(
			UpdateOpts{
				TargetA,
				a.a,
			},
			UpdateOpts{
				TargetB,
				a.b,
			},
		),
		a.and.Update(
			UpdateOpts{
				TargetA,
				a.a,
			},
			UpdateOpts{
				TargetB,
				a.b,
			},
		)
}
