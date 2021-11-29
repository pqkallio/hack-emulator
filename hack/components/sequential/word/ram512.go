package word

import (
	"github.com/pqkallio/hack-emulator/hack/components"
	"github.com/pqkallio/hack-emulator/hack/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/hack/components/combinational/word"
)

// RAM512 is a 512-register 16-bit wide RAM.
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
	c         chan components.OrderedVal16
	tickChan  chan bool
}

func NewRAM512() *RAM512 {
	return &RAM512{
		bit.NewDemux8Way(),
		NewRAM64(), NewRAM64(), NewRAM64(), NewRAM64(),
		NewRAM64(), NewRAM64(), NewRAM64(), NewRAM64(),
		word.NewMux8Way16(),
		make(chan components.OrderedVal16),
		make(chan bool),
	}
}

func (r *RAM512) Update(
	in uint16,
	load bool,
	addr0, addr1, addr2 bool,
	addr3, addr4, addr5 bool,
	addr6, addr7, addr8 bool,
	c chan components.OrderedVal16, idx int,
) uint16 {
	vals := [8]uint16{}

	aLoad, bLoad, cLoad, dLoad,
		eLoad, fLoad, gLoad, hLoad := r.demux8Way.Update(
		load,
		addr0,
		addr1,
		addr2,
	)

	go r.ramA.Update(in, aLoad, addr3, addr4, addr5, addr6, addr7, addr8, r.c, 0)
	go r.ramB.Update(in, bLoad, addr3, addr4, addr5, addr6, addr7, addr8, r.c, 1)
	go r.ramC.Update(in, cLoad, addr3, addr4, addr5, addr6, addr7, addr8, r.c, 2)
	go r.ramD.Update(in, dLoad, addr3, addr4, addr5, addr6, addr7, addr8, r.c, 3)
	go r.ramE.Update(in, eLoad, addr3, addr4, addr5, addr6, addr7, addr8, r.c, 4)
	go r.ramF.Update(in, fLoad, addr3, addr4, addr5, addr6, addr7, addr8, r.c, 5)
	go r.ramG.Update(in, gLoad, addr3, addr4, addr5, addr6, addr7, addr8, r.c, 6)
	go r.ramH.Update(in, hLoad, addr3, addr4, addr5, addr6, addr7, addr8, r.c, 7)

	for i := 0; i < 8; i++ {
		ud := <-r.c

		vals[ud.Idx] = ud.Val
	}

	val := r.mux8Way16.Update(
		vals[0], vals[1], vals[2], vals[3],
		vals[4], vals[5], vals[6], vals[7],
		addr0, addr1, addr2,
	)

	if c != nil {
		c <- components.OrderedVal16{val, idx}
	}

	return val
}

func (r *RAM512) Tick(c chan bool) {
	go r.ramA.Tick(r.tickChan)
	go r.ramB.Tick(r.tickChan)
	go r.ramC.Tick(r.tickChan)
	go r.ramD.Tick(r.tickChan)
	go r.ramE.Tick(r.tickChan)
	go r.ramF.Tick(r.tickChan)
	go r.ramG.Tick(r.tickChan)
	go r.ramH.Tick(r.tickChan)

	for i := 0; i < 8; i++ {
		<-r.tickChan
	}

	if c != nil {
		c <- true
	}
}
