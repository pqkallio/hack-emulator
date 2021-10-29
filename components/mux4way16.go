package components

type Mux4Way16 struct {
	a      Val
	b      Val
	c      Val
	d      Val
	sel    Val
	mux161 *Mux16
	mux162 *Mux16
	mux163 *Mux16
}

func NewMux4Way16() *Mux4Way16 {
	return &Mux4Way16{
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{},
		NewMux16(), NewMux16(), NewMux16(),
	}
}

func (mux4Way16 *Mux4Way16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			mux4Way16.a = opt.val
		case TargetB:
			mux4Way16.b = opt.val
		case TargetC:
			mux4Way16.c = opt.val
		case TargetD:
			mux4Way16.d = opt.val
		case TargetSel:
			mux4Way16.sel = opt.val
		}
	}

	sel := mux4Way16.sel
	sel0 := SelectChan{sel.GetSel() & 1}
	sel1 := SelectChan{(sel.GetSel() & 2) >> 1}

	abMux := mux4Way16.mux161.Update(
		UpdateOpts{TargetA, mux4Way16.a},
		UpdateOpts{TargetB, mux4Way16.b},
		UpdateOpts{TargetSel, &sel0},
	)
	cdMux := mux4Way16.mux161.Update(
		UpdateOpts{TargetA, mux4Way16.c},
		UpdateOpts{TargetB, mux4Way16.d},
		UpdateOpts{TargetSel, &sel0},
	)

	return mux4Way16.mux161.Update(
		UpdateOpts{TargetA, abMux},
		UpdateOpts{TargetB, cdMux},
		UpdateOpts{TargetSel, &sel1},
	)
}
