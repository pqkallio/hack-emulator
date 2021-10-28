package components

type Not16 struct {
}

func NewNot16() *Not16 {
	return &Not16{}
}

func (not16 *Not16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			return &SixteenChan{^opt.val.GetUint16()}
		}
	}

	return &InvalidVal{}
}
