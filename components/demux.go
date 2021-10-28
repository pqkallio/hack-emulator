package components

type Demux struct {
	in      bool
	sel     uint8
	targetA Target
	targetB Target
	outA    Out
	outB    Out
}

func NewDemux(targetA, targetB Target, outA, outB Out) *Demux {
	return &Demux{
		targetA: targetA,
		targetB: targetB,
		outA:    outA,
		outB:    outB,
	}
}

func (demux *Demux) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			demux.in = opt.val.GetBool()
		case TargetSel:
			demux.sel = opt.val.GetSel()
		}
	}

	sel := demux.sel

	demux.outA.Update(UpdateOpts{
		demux.targetA,
		&SingleChan{(sel == 0) && demux.in},
	})

	return demux.outB.Update(UpdateOpts{
		demux.targetB,
		&SingleChan{(sel == 1) && demux.in},
	})
}
