package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/components/combinational/word"
)

type RAM8 struct {
	demux8Way *bit.Demux8Way
	regA      *Register
	regB      *Register
	regC      *Register
	regD      *Register
	regE      *Register
	regF      *Register
	regG      *Register
	regH      *Register
	mux8Way16 *word.Mux8Way16
}

func NewRAM8() *RAM8 {
	return &RAM8{
		bit.NewDemux8Way(),
		NewRegister(), NewRegister(), NewRegister(), NewRegister(),
		NewRegister(), NewRegister(), NewRegister(), NewRegister(),
		word.NewMux8Way16(),
	}
}

func (r *RAM8) Update(in uint16, load, addr0, addr1, addr2 bool) uint16 {
	aLoad, bLoad, cLoad, dLoad,
		eLoad, fLoad, gLoad, hLoad := r.demux8Way.Update(
		load,
		addr0,
		addr1,
		addr2,
	)

	a := r.regA.Update(in, aLoad)
	b := r.regB.Update(in, bLoad)
	c := r.regC.Update(in, cLoad)
	d := r.regD.Update(in, dLoad)
	e := r.regE.Update(in, eLoad)
	f := r.regF.Update(in, fLoad)
	g := r.regG.Update(in, gLoad)
	h := r.regH.Update(in, hLoad)

	return r.mux8Way16.Update(
		a,
		b,
		c,
		d,
		e,
		f,
		g,
		h,
		addr0,
		addr1,
		addr2,
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
