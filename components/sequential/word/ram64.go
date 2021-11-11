package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/components/combinational/word"
)

type RAM64 struct {
	demux8Way *bit.Demux8Way
	ramA      *RAM8
	ramB      *RAM8
	ramC      *RAM8
	ramD      *RAM8
	ramE      *RAM8
	ramF      *RAM8
	ramG      *RAM8
	ramH      *RAM8
	mux8Way16 *word.Mux8Way16
}

func NewRAM64() *RAM64 {
	return &RAM64{
		bit.NewDemux8Way(),
		NewRAM8(), NewRAM8(), NewRAM8(), NewRAM8(),
		NewRAM8(), NewRAM8(), NewRAM8(), NewRAM8(),
		word.NewMux8Way16(),
	}
}

func (r *RAM64) Update(
	in uint16,
	load bool,
	addr0, addr1, addr2, addr3, addr4, addr5 bool,
) uint16 {
	aLoad, bLoad, cLoad, dLoad,
		eLoad, fLoad, gLoad, hLoad := r.demux8Way.Update(
		load,
		addr0,
		addr1,
		addr2,
	)

	a := r.ramA.Update(in, aLoad, addr3, addr4, addr5)
	b := r.ramB.Update(in, bLoad, addr3, addr4, addr5)
	c := r.ramC.Update(in, cLoad, addr3, addr4, addr5)
	d := r.ramD.Update(in, dLoad, addr3, addr4, addr5)
	e := r.ramE.Update(in, eLoad, addr3, addr4, addr5)
	f := r.ramF.Update(in, fLoad, addr3, addr4, addr5)
	g := r.ramG.Update(in, gLoad, addr3, addr4, addr5)
	h := r.ramH.Update(in, hLoad, addr3, addr4, addr5)

	return r.mux8Way16.Update(
		a, b, c, d, e, f, g, h,
		addr0, addr1, addr2,
	)
}

func (r *RAM64) Tick() {
	r.ramA.Tick()
	r.ramB.Tick()
	r.ramC.Tick()
	r.ramD.Tick()
	r.ramE.Tick()
	r.ramF.Tick()
	r.ramG.Tick()
	r.ramH.Tick()
}
