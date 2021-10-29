package components

type Demux struct {
	in      Val
	sel     bool
	targetA Target
	targetB Target
	outA    Out
	outB    Out
	not     *Not
	and1    *And
	and2    *And
}

func NewDemux(targetA, targetB Target, outA, outB Out) *Demux {
	return &Demux{
		&InvalidVal{},
		false,
		targetA, targetB,
		outA, outB,
		NewNot(),
		NewAnd(), NewAnd(),
	}
}

func (demux *Demux) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			demux.in = opt.val
		case TargetSel:
			demux.sel = (opt.val.GetSel() & 1) != 0
		}
	}

	sel := demux.sel

	notSel := demux.not.Update(UpdateOpts{TargetIn, &SingleChan{sel}})

	demux.outA.Update(UpdateOpts{
		demux.targetA,
		demux.and1.Update(
			UpdateOpts{TargetA, demux.in},
			UpdateOpts{TargetB, notSel},
		),
	})

	demux.outB.Update(UpdateOpts{
		demux.targetB,
		demux.and2.Update(
			UpdateOpts{TargetA, demux.in},
			UpdateOpts{TargetB, &SingleChan{demux.sel}},
		),
	})

	return &InvalidVal{}
}
