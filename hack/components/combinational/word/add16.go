package word

import (
	"github.com/pqkallio/hack-emulator/hack/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

const nFullAdders = 15

// Add16 is a 16-bit adder.
type Add16 struct {
	ha  *bit.HalfAdder
	fas [nFullAdders]*bit.FullAdder
}

func NewAdd16() *Add16 {
	fas := [nFullAdders]*bit.FullAdder{}

	for i := 0; i < nFullAdders; i++ {
		fas[i] = bit.NewFullAdder()
	}

	return &Add16{
		bit.NewHalfAdder(),
		fas,
	}
}

// Updates adds the input values and returns the sum.
// Carry bit after the operation is ignored.
func (add16 *Add16) Update(a, b uint16) uint16 {
	var (
		out   uint16
		carry bool
	)

	sum, carry := add16.ha.Update(
		util.GetBoolFromUint16(a, 0),
		util.GetBoolFromUint16(b, 0),
	)

	if sum {
		out = 1
	}

	for i := uint16(0); i < nFullAdders; i++ {
		sum, carry = add16.fas[i].Update(
			util.GetBoolFromUint16(a, i+1),
			util.GetBoolFromUint16(b, i+1),
			carry,
		)

		if sum {
			out |= 1 << (i + 1)
		}
	}

	return out
}
