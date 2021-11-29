package word

import (
	"github.com/pqkallio/hack-emulator/hack/components"
	"github.com/pqkallio/hack-emulator/hack/components/sequential/bit"
	"github.com/pqkallio/hack-emulator/util"
)

// Register is a 16-bit register.
type Register struct {
	curr, next uint16 // for debugging purposes
	bits       [16]*bit.Bit
	c          chan components.OrderedVal
	tickChan   chan bool
}

func NewRegister() *Register {
	bits := [16]*bit.Bit{}

	for i := 0; i < 16; i++ {
		bits[i] = bit.NewBit()
	}

	return &Register{0, 0, bits, make(chan components.OrderedVal), make(chan bool)}
}

func (reg *Register) Update(in uint16, load bool, c chan components.OrderedVal16, idx int) uint16 {
	if load {
		reg.next = in
	}

	outVal := uint16(0)

	for i, bit := range reg.bits {
		go bit.Update(util.GetBoolFromUint16(in, uint16(i)), load, reg.c, i)
	}

	for i := 0; i < 16; i++ {
		d1 := <-reg.c

		if d1.Val {
			outVal |= 1 << uint16(d1.Idx)
		}
	}

	if c != nil {
		c <- components.OrderedVal16{outVal, idx}
	}

	return outVal
}

func (reg *Register) Tick(c chan bool) {
	reg.curr = reg.next

	for _, bit := range reg.bits {
		go bit.Tick(reg.tickChan)
	}

	for i := 0; i < 16; i++ {
		<-reg.tickChan
	}

	if c != nil {
		c <- true
	}
}
