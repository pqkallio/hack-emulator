package components

type Or16 struct {
	a   Val
	b   Val
	ors [16]*Or
}

func NewOr16() *Or16 {
	ors := [16]*Or{}

	for i := 0; i < 16; i++ {
		ors[i] = NewOr()
	}

	return &Or16{&InvalidVal{}, &InvalidVal{}, ors}
}

func (or16 *Or16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			or16.a = opt.val
		case TargetB:
			or16.b = opt.val
		}
	}

	var out uint16

	for i, or := range or16.ors {
		val := or.Update(
			UpdateOpts{TargetA, &SingleChan{or16.a.GetBoolFromUint16(uint16(i))}},
			UpdateOpts{TargetB, &SingleChan{or16.b.GetBoolFromUint16(uint16(i))}},
		).GetBool()

		if val {
			out |= 1 << i
		}
	}

	return &SixteenChan{out}
}
