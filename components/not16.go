package components

type Not16 struct {
	in   uint16
	nots [16]*Not
}

func NewNot16() *Not16 {
	nots := [16]*Not{}

	for i := 0; i < 16; i++ {
		nots[i] = NewNot()
	}

	return &Not16{0, nots}
}

func (not16 *Not16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			not16.in = opt.val.GetUint16()
		}
	}

	var out uint16

	for i, not := range not16.nots {
		val := not.Update(UpdateOpts{
			TargetIn,
			&SingleChan{(not16.in & 1 << i) != 0},
		}).GetBool()

		if val {
			out |= 1 << i
		}
	}

	return &SixteenChan{out}
}
