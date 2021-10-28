package components

type And16 struct {
	a uint16
	b uint16
}

func NewAnd16() *And16 {
	return &And16{}
}

func (and16 *And16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			and16.a = opt.val.GetUint16()
		case TargetB:
			and16.b = opt.val.GetUint16()
		}
	}

	return &SixteenChan{and16.a & and16.b}
}
