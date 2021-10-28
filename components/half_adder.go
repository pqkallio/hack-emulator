package components

type HalfAdder struct {
	a   bool
	b   bool
	xor Xor
	and And
}

func NewHalfAdder() *HalfAdder {
	return &HalfAdder{false, false, Xor{}, And{}}
}

func (a *HalfAdder) Update(opts ...UpdateOpts) (Val, Val) {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			a.a = opt.val.GetBool()
		case TargetB:
			a.b = opt.val.GetBool()
		}
	}

	return a.xor.Update(
			UpdateOpts{
				TargetA,
				&SingleChan{a.a},
			},
			UpdateOpts{
				TargetB,
				&SingleChan{a.b},
			},
		),
		a.and.Update(
			UpdateOpts{
				TargetA,
				&SingleChan{a.a},
			},
			UpdateOpts{
				TargetB,
				&SingleChan{a.b},
			},
		)
}
