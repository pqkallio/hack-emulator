package components

const nFullAdders = 15

type Add16 struct {
	a, b Val
	ha   *HalfAdder
	fas  [nFullAdders]*FullAdder
}

func NewAdd16() *Add16 {
	fas := [nFullAdders]*FullAdder{}

	for i := 0; i < nFullAdders; i++ {
		fas[i] = NewFullAdder()
	}

	return &Add16{
		&SingleChan{}, &SingleChan{},
		NewHalfAdder(),
		fas,
	}
}

func (a *Add16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			a.a = opt.val
		case TargetB:
			a.b = opt.val
		}
	}

	var (
		out   uint16
		carry Val
	)

	sum, carry := a.ha.Update(
		UpdateOpts{TargetA, &SingleChan{a.a.GetBoolFromUint16(0)}},
		UpdateOpts{TargetB, &SingleChan{a.b.GetBoolFromUint16(0)}},
	)

	if sum.GetBool() {
		out = 1
	}

	for i := uint16(0); i < nFullAdders; i++ {
		sum, carry = a.fas[i].Update(
			UpdateOpts{TargetA, &SingleChan{a.a.GetBoolFromUint16(i + 1)}},
			UpdateOpts{TargetB, &SingleChan{a.b.GetBoolFromUint16(i + 1)}},
			UpdateOpts{TargetC, carry},
		)

		if sum.GetBool() {
			out |= 1 << (i + 1)
		}
	}

	return &SixteenChan{out}
}
