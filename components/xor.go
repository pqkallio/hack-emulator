package components

type Xor struct {
	a bool
	b bool
}

func NewXor() *Xor {
	return &Xor{}
}

func (xor *Xor) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			xor.a = opt.val.GetBool()
		case TargetB:
			xor.b = opt.val.GetBool()
		}
	}

	return &SingleChan{val: xor.a != xor.b}
}
