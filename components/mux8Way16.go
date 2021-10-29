package components

type Mux8Way16 struct {
	a, b, c, d, e, f, g, h Val
	sel                    Val
	mux4Way161             *Mux4Way16
	mux4Way162             *Mux4Way16
	mux16                  *Mux16
}

func NewMux8Way16() *Mux8Way16 {
	return &Mux8Way16{
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{},
		NewMux4Way16(), NewMux4Way16(),
		NewMux16(),
	}
}

func (mux8Way16 *Mux8Way16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			mux8Way16.a = opt.val
		case TargetB:
			mux8Way16.b = opt.val
		case TargetC:
			mux8Way16.c = opt.val
		case TargetD:
			mux8Way16.d = opt.val
		case TargetE:
			mux8Way16.e = opt.val
		case TargetF:
			mux8Way16.f = opt.val
		case TargetG:
			mux8Way16.g = opt.val
		case TargetH:
			mux8Way16.h = opt.val
		case TargetSel:
			mux8Way16.sel = opt.val
		}
	}

	sel := mux8Way16.sel
	sel01 := SelectChan{sel.GetSel() & 3}
	sel2 := SelectChan{(sel.GetSel() >> 2) & 1}

	abcdMux := mux8Way16.mux4Way161.Update(
		UpdateOpts{TargetA, mux8Way16.a},
		UpdateOpts{TargetB, mux8Way16.b},
		UpdateOpts{TargetC, mux8Way16.c},
		UpdateOpts{TargetD, mux8Way16.d},
		UpdateOpts{TargetSel, &sel01},
	)
	efghMux := mux8Way16.mux4Way161.Update(
		UpdateOpts{TargetA, mux8Way16.e},
		UpdateOpts{TargetB, mux8Way16.f},
		UpdateOpts{TargetC, mux8Way16.g},
		UpdateOpts{TargetD, mux8Way16.h},
		UpdateOpts{TargetSel, &sel01},
	)

	return mux8Way16.mux16.Update(
		UpdateOpts{TargetA, abcdMux},
		UpdateOpts{TargetB, efghMux},
		UpdateOpts{TargetSel, &sel2},
	)
}
