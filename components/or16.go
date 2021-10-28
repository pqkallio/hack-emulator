package components

type Or16 struct {
	a uint16
	b uint16
}

func NewOr16() *Or16 {
	return &Or16{}
}

func (or16 *Or16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			or16.a = opt.val.GetUint16()
		case TargetB:
			or16.b = opt.val.GetUint16()
		}
	}

	return &SixteenChan{or16.a | or16.b}
}
