package components

type Mux8Way16 struct {
	vals map[uint8]uint16
	sel  uint8
}

func NewMux8Way16() *Mux8Way16 {
	return &Mux8Way16{
		vals: make(map[uint8]uint16),
	}
}

func (mux8Way16 *Mux8Way16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			fallthrough
		case TargetB:
			fallthrough
		case TargetC:
			fallthrough
		case TargetD:
			fallthrough
		case TargetE:
			fallthrough
		case TargetF:
			fallthrough
		case TargetG:
			fallthrough
		case TargetH:
			mux8Way16.vals[uint8(opt.target)] = opt.val.GetUint16()
		case TargetSel:
			mux8Way16.sel = opt.val.GetSel()
		}
	}

	return &SixteenChan{mux8Way16.vals[mux8Way16.sel]}
}
