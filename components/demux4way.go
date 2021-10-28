package components

type Demux4Way struct {
	in      bool
	sel     uint8
	targetA Target
	targetB Target
	targetC Target
	targetD Target
	outA    Out
	outB    Out
	outC    Out
	outD    Out
}

func NewDemux4Way(
	targetA,
	targetB,
	targetC,
	targetD Target,
	outA,
	outB,
	outC,
	outD Out,
) *Demux4Way {
	return &Demux4Way{
		targetA: targetA,
		targetB: targetB,
		targetC: targetC,
		targetD: targetD,
		outA:    outA,
		outB:    outB,
		outC:    outC,
		outD:    outD,
	}
}

func (demux4Way *Demux4Way) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			demux4Way.in = opt.val.GetBool()
		case TargetSel:
			demux4Way.sel = opt.val.GetSel()
		}
	}

	sel := demux4Way.sel

	demux4Way.outA.Update(UpdateOpts{
		demux4Way.targetA,
		&SingleChan{(sel == 0) && demux4Way.in},
	})

	demux4Way.outB.Update(UpdateOpts{
		demux4Way.targetB,
		&SingleChan{(sel == 1) && demux4Way.in},
	})

	demux4Way.outC.Update(UpdateOpts{
		demux4Way.targetC,
		&SingleChan{(sel == 2) && demux4Way.in},
	})

	return demux4Way.outD.Update(UpdateOpts{
		demux4Way.targetD,
		&SingleChan{(sel == 3) && demux4Way.in},
	})
}
