package components

type Mux struct {
	a   bool
	b   bool
	sel uint8
}

func NewMux() *Mux {
	return &Mux{}
}

func (mux *Mux) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			mux.a = opt.val.GetBool()
		case TargetB:
			mux.b = opt.val.GetBool()
		case TargetSel:
			mux.sel = opt.val.GetSel()
		}
	}

	sel := mux.sel

	return &SingleChan{
		((sel == 0) && mux.a) || ((sel == 1) && mux.b),
	}
}
