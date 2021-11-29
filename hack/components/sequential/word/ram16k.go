package word

import (
	"github.com/pqkallio/hack-emulator/hack/components"
	"github.com/pqkallio/hack-emulator/hack/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/hack/components/combinational/word"
)

// RAM16K is a 16384-register 16-bit RAM.
type RAM16K struct {
	demux4Way *bit.Demux4Way
	ramA      *RAM4K
	ramB      *RAM4K
	ramC      *RAM4K
	ramD      *RAM4K
	mux4Way16 *word.Mux4Way16
	c         chan components.OrderedVal16
	tickChan  chan bool
}

func NewRAM16K() *RAM16K {
	return &RAM16K{
		bit.NewDemux4Way(),
		NewRAM4K(), NewRAM4K(), NewRAM4K(), NewRAM4K(),
		word.NewMux4Way16(),
		make(chan components.OrderedVal16),
		make(chan bool),
	}
}

func (r *RAM16K) Update(
	in uint16,
	load bool,
	addr0, addr1, addr2 bool,
	addr3, addr4, addr5 bool,
	addr6, addr7, addr8 bool,
	addr9, addr10, addr11 bool,
	addr12, addr13 bool,
	ch chan components.OrderedVal16, idx int,
) uint16 {
	vals := [4]uint16{}

	aLoad, bLoad, cLoad, dLoad := r.demux4Way.Update(
		load,
		addr0,
		addr1,
	)

	go r.ramA.Update(in, aLoad, addr2, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, addr12, addr13, r.c, 0)
	go r.ramB.Update(in, bLoad, addr2, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, addr12, addr13, r.c, 1)
	go r.ramC.Update(in, cLoad, addr2, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, addr12, addr13, r.c, 2)
	go r.ramD.Update(in, dLoad, addr2, addr3, addr4, addr5, addr6, addr7, addr8, addr9, addr10, addr11, addr12, addr13, r.c, 3)

	for i := 0; i < 4; i++ {
		ud := <-r.c

		vals[ud.Idx] = ud.Val
	}

	val := r.mux4Way16.Update(
		vals[0], vals[1], vals[2], vals[3],
		addr0, addr1,
		nil, 0,
	)

	if ch != nil {
		ch <- components.OrderedVal16{val, idx}
	}

	return val
}

func (r *RAM16K) Tick() {
	go r.ramA.Tick(r.tickChan)
	go r.ramB.Tick(r.tickChan)
	go r.ramC.Tick(r.tickChan)
	go r.ramD.Tick(r.tickChan)

	for i := 0; i < 4; i++ {
		<-r.tickChan
	}
}
