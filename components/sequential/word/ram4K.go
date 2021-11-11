package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/components/combinational/word"
)

type RAM4K struct {
	demux8Way *bit.Demux8Way
	ramA      *RAM512
	ramB      *RAM512
	ramC      *RAM512
	ramD      *RAM512
	ramE      *RAM512
	ramF      *RAM512
	ramG      *RAM512
	ramH      *RAM512
	mux8Way16 *word.Mux8Way16
}

func NewRAM4K() *RAM4K {
	return &RAM4K{
		bit.NewDemux8Way(),
		NewRAM512(), NewRAM512(), NewRAM512(), NewRAM512(),
		NewRAM512(), NewRAM512(), NewRAM512(), NewRAM512(),
		word.NewMux8Way16(),
	}
}

func (r *RAM4K) Update(
	in uint16,
	load bool,
	addr0, addr1, addr2 bool,
	addr3, addr4, addr5 bool,
	addr6, addr7, addr8 bool,
	addr9, addr10, addr11 bool,
) uint16 {
	aLoad, bLoad, cLoad, dLoad,
		eLoad, fLoad, gLoad, hLoad := r.demux8Way.Update(
		load,
		addr0,
		addr1,
		addr2,
	)

	a := r.ramA.Update(in, aLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11)
	b := r.ramB.Update(in, bLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11)
	c := r.ramC.Update(in, cLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11)
	d := r.ramD.Update(in, dLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11)
	e := r.ramE.Update(in, eLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11)
	f := r.ramF.Update(in, fLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11)
	g := r.ramG.Update(in, gLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11)
	h := r.ramH.Update(in, hLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11)

	return r.mux8Way16.Update(
		a, b, c, d, e, f, g, h,
		addr0, addr1, addr2,
	)
}

func (r *RAM4K) Tick() {
	r.ramA.Tick()
	r.ramB.Tick()
	r.ramC.Tick()
	r.ramD.Tick()
	r.ramE.Tick()
	r.ramF.Tick()
	r.ramG.Tick()
	r.ramH.Tick()
}
