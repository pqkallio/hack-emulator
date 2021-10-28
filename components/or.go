package components

type Or struct {
	a bool
	b bool
}

func NewOr() *Or {
	return &Or{}
}

func (or *Or) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			or.a = opt.val.GetBool()
		case TargetB:
			or.b = opt.val.GetBool()
		}
	}

	return &SingleChan{val: or.a || or.b}
}
