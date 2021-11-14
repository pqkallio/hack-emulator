package bit

import "github.com/pqkallio/hack-emulator/components"

type Bit struct {
	dff *DFF
}

func NewBit() *Bit {
	return &Bit{
		NewDFF(),
	}
}

func (bit *Bit) Update(data, load bool, c chan components.OrderedVal, idx int) bool {
	var val bool

	if load {
		val = bit.dff.Update(data)
	} else {
		val = bit.dff.Update(bit.dff.curr)
	}

	if c != nil {
		c <- components.OrderedVal{val, idx}
	}

	return val
}

func (bit *Bit) Tick(c chan bool) {
	bit.dff.Tick()

	if c != nil {
		c <- true
	}
}
