package word

import (
	"github.com/pqkallio/hack-emulator/components"
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

// Or16 is a 16-bit OR gate.
type Or16 struct {
	ors [16]*bit.Or
	c   chan components.OrderedVal
}

func NewOr16() *Or16 {
	ors := [16]*bit.Or{}

	for i := 0; i < 16; i++ {
		ors[i] = bit.NewOr()
	}

	return &Or16{ors, make(chan components.OrderedVal)}
}

func (or16 *Or16) Update(a, b uint16) uint16 {
	var out uint16

	for i, or := range or16.ors {
		go or.Update(
			util.GetBoolFromUint16(a, uint16(i)),
			util.GetBoolFromUint16(b, uint16(i)),
			or16.c, i,
		)
	}

	for i := 0; i < 16; i++ {
		ov := <-or16.c

		if ov.Val {
			out |= 1 << ov.Idx
		}
	}

	return out
}
