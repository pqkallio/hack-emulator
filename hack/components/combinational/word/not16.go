package word

import (
	"github.com/pqkallio/hack-emulator/hack/components"
	"github.com/pqkallio/hack-emulator/hack/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

// Not16 is a 16-bit NOT gate.
type Not16 struct {
	nots [16]*bit.Not
	c    chan components.OrderedVal
}

func NewNot16() *Not16 {
	nots := [16]*bit.Not{}

	for i := 0; i < 16; i++ {
		nots[i] = bit.NewNot()
	}

	return &Not16{nots, make(chan components.OrderedVal)}
}

func (not16 *Not16) Update(in uint16) uint16 {
	var out uint16

	for i, not := range not16.nots {
		go not.Update(util.GetBoolFromUint16(in, uint16(i)), not16.c, i)
	}

	for i := 0; i < 16; i++ {
		ov := <-not16.c

		if ov.Val {
			out |= 1 << ov.Idx
		}
	}

	return out
}
