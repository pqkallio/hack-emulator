package components

type Mux4Way16 struct {
	vals map[uint8]uint16
	sel  uint8
}

func NewMux4Way16() *Mux4Way16 {
	return &Mux4Way16{
		vals: make(map[uint8]uint16),
	}
}

func (mux4Way16 *Mux4Way16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			fallthrough
		case TargetB:
			fallthrough
		case TargetC:
			fallthrough
		case TargetD:
			mux4Way16.vals[uint8(opt.target)] = opt.val.GetUint16()
		case TargetSel:
			mux4Way16.sel = opt.val.GetSel()
		}
	}

	return &SixteenChan{mux4Way16.vals[mux4Way16.sel]}
}
