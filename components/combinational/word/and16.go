package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

type And16 struct {
	ands [16]*bit.And
}

func NewAnd16() *And16 {
	ands := [16]*bit.And{}

	for i := 0; i < 16; i++ {
		ands[i] = bit.NewAnd()
	}

	return &And16{ands}
}

func (and16 *And16) Update(a, b uint16) uint16 {
	var out uint16

	for i, and := range and16.ands {
		val := and.Update(
			util.GetBoolFromUint16(a, uint16(i)),
			util.GetBoolFromUint16(b, uint16(i)),
		)

		if val {
			out |= 1 << i
		}
	}

	return out
}
