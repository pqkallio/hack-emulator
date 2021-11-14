package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

type Mux16 struct {
	muxs [16]*bit.Mux
}

func NewMux16() *Mux16 {
	muxs := [16]*bit.Mux{}

	for i := 0; i < 16; i++ {
		muxs[i] = bit.NewMux()
	}

	return &Mux16{muxs}
}

func (mux16 *Mux16) Update(a, b uint16, sel bool) uint16 {
	var out uint16

	for i, mux := range mux16.muxs {
		val := mux.Update(
			util.GetBoolFromUint16(a, uint16(i)),
			util.GetBoolFromUint16(b, uint16(i)),
			sel,
		)

		if val {
			out |= 1 << i
		}
	}

	return out
}
