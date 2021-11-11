package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

type Not16 struct {
	nots [16]*bit.Not
}

func NewNot16() *Not16 {
	nots := [16]*bit.Not{}

	for i := 0; i < 16; i++ {
		nots[i] = bit.NewNot()
	}

	return &Not16{nots}
}

func (not16 *Not16) Update(in uint16) uint16 {
	var out uint16

	for i, not := range not16.nots {
		val := not.Update(util.GetBoolFromUint16(in, uint16(i)))

		if val {
			out |= 1 << i
		}
	}

	return out
}
