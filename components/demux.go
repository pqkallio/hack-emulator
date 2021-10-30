package components

type Demux struct {
	in   Val
	sel  Val
	not  *Not
	and1 *And
	and2 *And
}

func NewDemux() *Demux {
	return &Demux{
		&InvalidVal{},
		&InvalidVal{},
		NewNot(),
		NewAnd(), NewAnd(),
	}
}

func (demux *Demux) Update(opts ...UpdateOpts) (Val, Val) {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			demux.in = opt.val
		case TargetSel0:
			demux.sel = opt.val
		}
	}

	sel := demux.sel

	notSel := demux.not.Update(UpdateOpts{TargetIn, sel})

	return demux.and1.Update(
			UpdateOpts{TargetA, demux.in},
			UpdateOpts{TargetB, notSel},
		),
		demux.and2.Update(
			UpdateOpts{TargetA, demux.in},
			UpdateOpts{TargetB, sel},
		)
}
