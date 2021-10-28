package components

type And struct {
	a bool
	b bool
}

func NewAnd() *And {
	return &And{}
}

func (and *And) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			and.a = opt.val.GetBool()
		case TargetB:
			and.b = opt.val.GetBool()
		}
	}

	return &SingleChan{val: and.a && and.b}
}
