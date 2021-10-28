package components

type Demux8Way struct {
	in      bool
	sel     uint8
	targetA Target
	targetB Target
	targetC Target
	targetD Target
	targetE Target
	targetF Target
	targetG Target
	targetH Target
	outA    Out
	outB    Out
	outC    Out
	outD    Out
	outE    Out
	outF    Out
	outG    Out
	outH    Out
}

func NewDemux8Way(
	targetA,
	targetB,
	targetC,
	targetD,
	targetE,
	targetF,
	targetG,
	targetH Target,
	outA,
	outB,
	outC,
	outD,
	outE,
	outF,
	outG,
	outH Out,
) *Demux8Way {
	return &Demux8Way{
		targetA: targetA,
		targetB: targetB,
		targetC: targetC,
		targetD: targetD,
		targetE: targetE,
		targetF: targetF,
		targetG: targetG,
		targetH: targetH,
		outA:    outA,
		outB:    outB,
		outC:    outC,
		outD:    outD,
		outE:    outE,
		outF:    outF,
		outG:    outG,
		outH:    outH,
	}
}

func (demux8Way *Demux8Way) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			demux8Way.in = opt.val.GetBool()
		case TargetSel:
			demux8Way.sel = opt.val.GetSel()
		}
	}

	sel := demux8Way.sel

	demux8Way.outA.Update(UpdateOpts{
		demux8Way.targetA,
		&SingleChan{(sel == 0) && demux8Way.in},
	})

	demux8Way.outB.Update(UpdateOpts{
		demux8Way.targetB,
		&SingleChan{(sel == 1) && demux8Way.in},
	})

	demux8Way.outC.Update(UpdateOpts{
		demux8Way.targetC,
		&SingleChan{(sel == 2) && demux8Way.in},
	})

	demux8Way.outD.Update(UpdateOpts{
		demux8Way.targetD,
		&SingleChan{(sel == 3) && demux8Way.in},
	})

	demux8Way.outE.Update(UpdateOpts{
		demux8Way.targetE,
		&SingleChan{(sel == 4) && demux8Way.in},
	})

	demux8Way.outF.Update(UpdateOpts{
		demux8Way.targetF,
		&SingleChan{(sel == 5) && demux8Way.in},
	})

	demux8Way.outG.Update(UpdateOpts{
		demux8Way.targetG,
		&SingleChan{(sel == 6) && demux8Way.in},
	})

	return demux8Way.outH.Update(UpdateOpts{
		demux8Way.targetH,
		&SingleChan{(sel == 7) && demux8Way.in},
	})
}
