package word

import (
	"github.com/pqkallio/hack-emulator/hack/components"
	"github.com/pqkallio/hack-emulator/hack/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

type Mux16 struct {
	muxs [16]*bit.Mux
	c    chan components.OrderedVal
}

func NewMux16() *Mux16 {
	muxs := [16]*bit.Mux{}

	for i := 0; i < 16; i++ {
		muxs[i] = bit.NewMux()
	}

	return &Mux16{muxs, make(chan components.OrderedVal)}
}

func (mux16 *Mux16) Update(a, b uint16, sel bool, c chan components.OrderedVal16, idx int) uint16 {
	var out uint16

	for i, mux := range mux16.muxs {
		go mux.Update(
			util.GetBoolFromUint16(a, uint16(i)),
			util.GetBoolFromUint16(b, uint16(i)),
			sel,
			mux16.c, i,
		)
	}

	for i := 0; i < 16; i++ {
		ov := <-mux16.c

		if ov.Val {
			out |= 1 << ov.Idx
		}
	}

	if c != nil {
		c <- components.OrderedVal16{out, idx}
	}

	return out
}
