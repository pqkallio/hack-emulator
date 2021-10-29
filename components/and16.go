package components

type And16 struct {
	a    Val
	b    Val
	ands [16]*And
}

func NewAnd16() *And16 {
	ands := [16]*And{}

	for i := 0; i < 16; i++ {
		ands[i] = NewAnd()
	}

	return &And16{&InvalidVal{}, &InvalidVal{}, ands}
}

func (and16 *And16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			and16.a = opt.val
		case TargetB:
			and16.b = opt.val
		}
	}

	var out uint16

	for i, and := range and16.ands {
		val := and.Update(
			UpdateOpts{TargetA, &SingleChan{and16.a.GetBoolFromUint16(uint16(i))}},
			UpdateOpts{TargetB, &SingleChan{and16.b.GetBoolFromUint16(uint16(i))}},
		).GetBool()

		if val {
			out |= 1 << i
		}
	}

	return &SixteenChan{out}
}
