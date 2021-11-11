package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/components/combinational/word"
)

type RAM512 struct {
	demux8Way *bit.Demux8Way
	ramA      *RAM64
	ramB      *RAM64
	ramC      *RAM64
	ramD      *RAM64
	ramE      *RAM64
	ramF      *RAM64
	ramG      *RAM64
	ramH      *RAM64
	mux8Way16 *word.Mux8Way16
}

func NewRAM512() *RAM512 {
	return &RAM512{
		bit.NewDemux8Way(),
		NewRAM64(), NewRAM64(), NewRAM64(), NewRAM64(),
		NewRAM64(), NewRAM64(), NewRAM64(), NewRAM64(),
		word.NewMux8Way16(),
	}
}

func (r *RAM512) Update(
	in uint16,
	load bool,
	addr0, addr1, addr2 bool,
	addr3, addr4, addr5 bool,
	addr6, addr7, addr8 bool,
) uint16 {
	aLoad, bLoad, cLoad, dLoad,
		eLoad, fLoad, gLoad, hLoad := r.demux8Way.Update(
		load,
		addr0,
		addr1,
		addr2,
	)

	a := r.ramA.Update(in, aLoad, addr3, addr4, addr5, addr6, addr7, addr8)
	b := r.ramB.Update(in, bLoad, addr3, addr4, addr5, addr6, addr7, addr8)
	c := r.ramC.Update(in, cLoad, addr3, addr4, addr5, addr6, addr7, addr8)
	d := r.ramD.Update(in, dLoad, addr3, addr4, addr5, addr6, addr7, addr8)
	e := r.ramE.Update(in, eLoad, addr3, addr4, addr5, addr6, addr7, addr8)
	f := r.ramF.Update(in, fLoad, addr3, addr4, addr5, addr6, addr7, addr8)
	g := r.ramG.Update(in, gLoad, addr3, addr4, addr5, addr6, addr7, addr8)
	h := r.ramH.Update(in, hLoad, addr3, addr4, addr5, addr6, addr7, addr8)

	return r.mux8Way16.Update(
		a, b, c, d, e, f, g, h,
		addr0, addr1, addr2,
	)
}

func (r *RAM512) Tick() {
	r.ramA.Tick()
	r.ramB.Tick()
	r.ramC.Tick()
	r.ramD.Tick()
	r.ramE.Tick()
	r.ramF.Tick()
	r.ramG.Tick()
	r.ramH.Tick()
}
