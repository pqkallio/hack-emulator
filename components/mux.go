package components

type Mux struct {
	a    Val
	b    Val
	sel  bool
	not  *Not
	and1 *And
	and2 *And
	or   *Or
}

func NewMux() *Mux {
	return &Mux{
		&InvalidVal{}, &InvalidVal{},
		false,
		NewNot(),
		NewAnd(), NewAnd(),
		NewOr(),
	}
}

func (mux *Mux) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			mux.a = opt.val
		case TargetB:
			mux.b = opt.val
		case TargetSel:
			mux.sel = opt.val.GetSel()&1 != 0
		}
	}

	sel := mux.sel

	notSel := mux.not.Update(UpdateOpts{TargetIn, &SingleChan{sel}})
	aSel := mux.and1.Update(
		UpdateOpts{TargetA, mux.a},
		UpdateOpts{TargetB, notSel},
	)
	bSel := mux.and2.Update(
		UpdateOpts{TargetA, &SingleChan{sel}},
		UpdateOpts{TargetB, mux.b},
	)

	return mux.or.Update(
		UpdateOpts{TargetA, aSel},
		UpdateOpts{TargetB, bSel},
	)
}
