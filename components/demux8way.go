package components

type Demux8Way struct {
	in     Val
	sel0   Val
	sel1   Val
	sel2   Val
	demux1 *Demux
	demux2 *Demux
	demux3 *Demux
	demux4 *Demux
	demux5 *Demux
	demux6 *Demux
	demux7 *Demux
}

func NewDemux8Way() *Demux8Way {
	return &Demux8Way{
		&InvalidVal{},
		&InvalidVal{}, &InvalidVal{}, &InvalidVal{},
		NewDemux(), NewDemux(), NewDemux(), NewDemux(),
		NewDemux(), NewDemux(), NewDemux(),
	}
}

func (demux8Way *Demux8Way) Update(
	opts ...UpdateOpts,
) (Val, Val, Val, Val, Val, Val, Val, Val) {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			demux8Way.in = opt.val
		case TargetSel0:
			demux8Way.sel0 = opt.val
		case TargetSel1:
			demux8Way.sel1 = opt.val
		case TargetSel2:
			demux8Way.sel2 = opt.val
		}
	}

	aceg, bdfh := demux8Way.demux1.Update(
		UpdateOpts{TargetIn, demux8Way.in},
		UpdateOpts{TargetSel0, demux8Way.sel0},
	)

	ae, cg := demux8Way.demux2.Update(
		UpdateOpts{TargetIn, aceg},
		UpdateOpts{TargetSel0, demux8Way.sel1},
	)
	bf, dh := demux8Way.demux3.Update(
		UpdateOpts{TargetIn, bdfh},
		UpdateOpts{TargetSel0, demux8Way.sel1},
	)

	a, e := demux8Way.demux4.Update(
		UpdateOpts{TargetIn, ae},
		UpdateOpts{TargetSel0, demux8Way.sel2},
	)
	c, g := demux8Way.demux5.Update(
		UpdateOpts{TargetIn, cg},
		UpdateOpts{TargetSel0, demux8Way.sel2},
	)
	b, f := demux8Way.demux6.Update(
		UpdateOpts{TargetIn, bf},
		UpdateOpts{TargetSel0, demux8Way.sel2},
	)
	d, h := demux8Way.demux7.Update(
		UpdateOpts{TargetIn, dh},
		UpdateOpts{TargetSel0, demux8Way.sel2},
	)

	return a, b, c, d, e, f, g, h
}
