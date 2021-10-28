package components

type FullAdder struct {
	a, b, c       Val
	ha1, ha2, ha3 *HalfAdder
}

func NewFullAdder() *FullAdder {
	return &FullAdder{
		&SingleChan{}, &SingleChan{}, &SingleChan{},
		NewHalfAdder(), NewHalfAdder(), NewHalfAdder(),
	}
}

func (a *FullAdder) Update(opts ...UpdateOpts) (Val, Val) {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			a.a = opt.val
		case TargetB:
			a.b = opt.val
		case TargetC:
			a.c = opt.val
		}
	}

	abSum, abCarry := a.ha1.Update(
		UpdateOpts{TargetA, a.a},
		UpdateOpts{TargetB, a.b},
	)

	sum, abcCarry := a.ha2.Update(
		UpdateOpts{TargetA, abSum},
		UpdateOpts{TargetB, a.c},
	)

	carry, _ := a.ha3.Update(
		UpdateOpts{TargetA, abcCarry},
		UpdateOpts{TargetB, abCarry},
	)

	return sum, carry
}
