package word

import (
	"github.com/pqkallio/hack-emulator/components"
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
	c         chan components.OrderedVal16
}

func NewRAM8() *RAM8 {
	return &RAM8{
		bit.NewDemux8Way(),
		NewRegister(), NewRegister(), NewRegister(), NewRegister(),
		NewRegister(), NewRegister(), NewRegister(), NewRegister(),
		word.NewMux8Way16(),
		make(chan components.OrderedVal16),
	}
}

func (r *RAM8) Update(
	in uint16,
	load, addr0, addr1, addr2 bool,
	c chan components.OrderedVal16,
	idx int,
) uint16 {
	aLoad, bLoad, cLoad, dLoad,
		eLoad, fLoad, gLoad, hLoad := r.demux8Way.Update(
		load,
		addr0,
		addr1,
		addr2,
	)

	vals := [8]uint16{}

	go r.regA.Update(in, aLoad, r.c, 0)
	go r.regB.Update(in, bLoad, r.c, 1)
	go r.regC.Update(in, cLoad, r.c, 2)
	go r.regD.Update(in, dLoad, r.c, 3)
	go r.regE.Update(in, eLoad, r.c, 4)
	go r.regF.Update(in, fLoad, r.c, 5)
	go r.regG.Update(in, gLoad, r.c, 6)
	go r.regH.Update(in, hLoad, r.c, 7)

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
