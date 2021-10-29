package components

type Nand struct {
	a Val
	b Val
}

func NewNand() *Nand {
	return &Nand{
		&InvalidVal{},
		&InvalidVal{},
	}
}

func (n *Nand) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			n.a = opt.val
		case TargetB:
			n.b = opt.val
		}
	}

	return &SingleChan{!(n.a.GetBool() && n.b.GetBool())}
}
