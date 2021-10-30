package components

type Demux4Way struct {
	in     Val
	sel0   Val
	sel1   Val
	demux1 *Demux
	demux2 *Demux
	demux3 *Demux
}

func NewDemux4Way() *Demux4Way {
	return &Demux4Way{
		&InvalidVal{},
		&InvalidVal{}, &InvalidVal{},
		NewDemux(), NewDemux(), NewDemux(),
	}
}

func (demux4Way *Demux4Way) Update(opts ...UpdateOpts) (Val, Val, Val, Val) {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			demux4Way.in = opt.val
		case TargetSel0:
			demux4Way.sel0 = opt.val
		case TargetSel1:
			demux4Way.sel1 = opt.val
		}
	}

	demux1a, demux1b := demux4Way.demux1.Update(
		UpdateOpts{TargetIn, demux4Way.in},
		UpdateOpts{TargetSel0, demux4Way.sel0},
	)

	a, c := demux4Way.demux2.Update(
		UpdateOpts{TargetIn, demux1a},
		UpdateOpts{TargetSel0, demux4Way.sel1},
	)
	b, d := demux4Way.demux3.Update(
		UpdateOpts{TargetIn, demux1b},
		UpdateOpts{TargetSel0, demux4Way.sel1},
	)

	return a, b, c, d
}
