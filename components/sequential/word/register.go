package word

import (
	"github.com/pqkallio/hack-emulator/components/sequential/bit"
	"github.com/pqkallio/hack-emulator/util"
)

type Register struct {
	bits [16]*bit.Bit
}

func NewRegister() *Register {
	bits := [16]*bit.Bit{}

	for i := 0; i < 16; i++ {
		bits[i] = bit.NewBit()
	}

	return &Register{bits}
}

func (reg *Register) Update(in uint16, load bool) uint16 {
	outVal := uint16(0)

	for i, bit := range reg.bits {
		val := bit.Update(util.GetBoolFromUint16(in, uint16(i)), load)

		if val {
			outVal |= 1 << uint16(i)
		}
	}

	return outVal
}

func (reg *Register) Tick() {
	for _, bit := range reg.bits {
		bit.Tick()
	}
}
