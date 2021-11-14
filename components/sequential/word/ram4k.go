package word

import (
	"github.com/pqkallio/hack-emulator/components"
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
	c         chan components.OrderedVal16
}

func NewRAM4K() *RAM4K {
	return &RAM4K{
		bit.NewDemux8Way(),
		NewRAM512(), NewRAM512(), NewRAM512(), NewRAM512(),
		NewRAM512(), NewRAM512(), NewRAM512(), NewRAM512(),
		word.NewMux8Way16(),
		make(chan components.OrderedVal16),
	}
}

func (r *RAM4K) Update(
	in uint16,
	load bool,
	addr0, addr1, addr2 bool,
	addr3, addr4, addr5 bool,
	addr6, addr7, addr8 bool,
	addr9, addr10, addr11 bool,
	ch chan components.OrderedVal16, idx int,
) uint16 {
	vals := [8]uint16{}

	aLoad, bLoad, cLoad, dLoad,
		eLoad, fLoad, gLoad, hLoad := r.demux8Way.Update(
		load,
		addr0,
		addr1,
		addr2,
	)

	go r.ramA.Update(in, aLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, r.c, 0)
	go r.ramB.Update(in, bLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, r.c, 1)
	go r.ramC.Update(in, cLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, r.c, 2)
	go r.ramD.Update(in, dLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, r.c, 3)
	go r.ramE.Update(in, eLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, r.c, 4)
	go r.ramF.Update(in, fLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, r.c, 5)
	go r.ramG.Update(in, gLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, r.c, 6)
	go r.ramH.Update(in, hLoad, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, r.c, 7)

	for i := 0; i < 8; i++ {
		ud := <-r.c

		vals[ud.Idx] = ud.Val
	}

	val := r.mux8Way16.Update(
		vals[0], vals[1], vals[2], vals[3],
		vals[4], vals[5], vals[6], vals[7],
		addr0, addr1, addr2,
	)

	if ch != nil {
		ch <- components.OrderedVal16{val, idx}
	}

	return val
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
