package components

type RAM8 struct {
	demux8Way *Demux8Way
	regA      *Register
	regB      *Register
	regC      *Register
	regD      *Register
	regE      *Register
	regF      *Register
	regG      *Register
	regH      *Register
	mux8Way16 *Mux8Way16
}

func NewRAM8() *RAM8 {
	return &RAM8{
		NewDemux8Way(),
		NewRegister(), NewRegister(), NewRegister(), NewRegister(),
		NewRegister(), NewRegister(), NewRegister(), NewRegister(),
		NewMux8Way16(),
	}
}

func (r *RAM8) Update(opts ...UpdateOpts) Val {
	var load, addr0, addr1, addr2, in Val

	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			in = opt.val
		case TargetLoad:
			load = opt.val
		case TargetSel0:
			addr0 = opt.val
		case TargetSel1:
			addr1 = opt.val
		case TargetSel2:
			addr2 = opt.val
		}
	}

	if load == nil ||
		addr0 == nil ||
		addr1 == nil ||
		addr2 == nil ||
		in == nil {
		return &InvalidVal{}
	}

	aLoad, bLoad, cLoad, dLoad,
		eLoad, fLoad, gLoad, hLoad := r.demux8Way.Update(
		UpdateOpts{TargetIn, load},
		UpdateOpts{TargetSel0, addr0},
		UpdateOpts{TargetSel1, addr1},
		UpdateOpts{TargetSel2, addr2},
	)

	a := r.regA.Update(
		UpdateOpts{TargetIn, in},
		UpdateOpts{TargetLoad, aLoad},
	)
	b := r.regB.Update(
		UpdateOpts{TargetIn, in},
		UpdateOpts{TargetLoad, bLoad},
	)
	c := r.regC.Update(
		UpdateOpts{TargetIn, in},
		UpdateOpts{TargetLoad, cLoad},
	)
	d := r.regD.Update(
		UpdateOpts{TargetIn, in},
		UpdateOpts{TargetLoad, dLoad},
	)
	e := r.regE.Update(
		UpdateOpts{TargetIn, in},
		UpdateOpts{TargetLoad, eLoad},
	)
	f := r.regF.Update(
		UpdateOpts{TargetIn, in},
		UpdateOpts{TargetLoad, fLoad},
	)
	g := r.regG.Update(
		UpdateOpts{TargetIn, in},
		UpdateOpts{TargetLoad, gLoad},
	)
	h := r.regH.Update(
		UpdateOpts{TargetIn, in},
		UpdateOpts{TargetLoad, hLoad},
	)

	return r.mux8Way16.Update(
		UpdateOpts{TargetA, a},
		UpdateOpts{TargetB, b},
		UpdateOpts{TargetC, c},
		UpdateOpts{TargetD, d},
		UpdateOpts{TargetE, e},
		UpdateOpts{TargetF, f},
		UpdateOpts{TargetG, g},
		UpdateOpts{TargetH, h},
		UpdateOpts{TargetSel0, addr0},
		UpdateOpts{TargetSel1, addr1},
		UpdateOpts{TargetSel2, addr2},
	)
}

func (r *RAM8) Tick() {
	r.regA.Tick()
	r.regB.Tick()
	r.regC.Tick()
	r.regD.Tick()
	r.regE.Tick()
	r.regF.Tick()
	r.regG.Tick()
	r.regH.Tick()
}
