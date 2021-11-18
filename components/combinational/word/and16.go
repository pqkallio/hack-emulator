package word

import (
	"github.com/pqkallio/hack-emulator/components"
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

// And16 is a 16-bit AND gate.
type And16 struct {
	ands [16]*bit.And
	c    chan components.OrderedVal
}

func NewAnd16() *And16 {
	ands := [16]*bit.And{}

	for i := 0; i < 16; i++ {
		ands[i] = bit.NewAnd()
	}

	return &And16{ands, make(chan components.OrderedVal)}
}

func (and16 *And16) Update(a, b uint16) uint16 {
	var out uint16

	for i, and := range and16.ands {
		go and.Update(
			util.GetBoolFromUint16(a, uint16(i)),
			util.GetBoolFromUint16(b, uint16(i)),
			and16.c, i,
		)

	}

	for i := 0; i < 16; i++ {
		ov := <-and16.c

		if ov.Val {
			out |= 1 << ov.Idx
		}
	}

	return out
}
