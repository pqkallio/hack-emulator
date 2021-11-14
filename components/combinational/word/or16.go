package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

type Or16 struct {
	ors [16]*bit.Or
}

func NewOr16() *Or16 {
	ors := [16]*bit.Or{}

	for i := 0; i < 16; i++ {
		ors[i] = bit.NewOr()
	}

	return &Or16{ors}
}

func (or16 *Or16) Update(a, b uint16) uint16 {
	var out uint16

	for i, or := range or16.ors {
		val := or.Update(
			util.GetBoolFromUint16(a, uint16(i)),
			util.GetBoolFromUint16(b, uint16(i)),
		)

		if val {
			out |= 1 << i
		}
	}

	return out
}
