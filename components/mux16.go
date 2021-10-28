package components

type Mux16 struct {
	a   uint16
	b   uint16
	sel uint8
}

func NewMux16() *Mux16 {
	return &Mux16{}
}

func (mux16 *Mux16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			mux16.a = opt.val.GetUint16()
		case TargetB:
			mux16.b = opt.val.GetUint16()
		case TargetSel:
			mux16.sel = opt.val.GetSel()
		}
	}

	var out uint16

	switch mux16.sel {
	case 0:
		out = mux16.a
	case 1:
		out = mux16.b
	}

	return &SixteenChan{out}
}
