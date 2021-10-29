package components

type Not16 struct {
	in   Val
	nots [16]*Not
}

func NewNot16() *Not16 {
	nots := [16]*Not{}

	for i := 0; i < 16; i++ {
		nots[i] = NewNot()
	}

	return &Not16{&InvalidVal{}, nots}
}

func (not16 *Not16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			not16.in = opt.val
		}
	}

	var out uint16

	for i, not := range not16.nots {
		val := not.Update(UpdateOpts{
			TargetIn,
			&SingleChan{not16.in.GetBoolFromUint16(uint16(i))},
		}).GetBool()

		if val {
			out |= 1 << i
		}
	}

	return &SixteenChan{out}
}
