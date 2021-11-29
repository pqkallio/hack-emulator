package bit

import "github.com/pqkallio/hack-emulator/hack/components"

// Bit is a sequential component that represents a single bit register.
type Bit struct {
	dff *DFF
}

func NewBit() *Bit {
	return &Bit{
		NewDFF(),
	}
}

// Update evaluates the next state of the bit register and returns its current value.
//
// Inputs:
// 	- data: input data
// 	- load: load signal, if true, data is set as the next state
//	- c: channel to send the output to, optional
//
// Outputs:
// 	- output: current value of the bit register
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

// Tick sets the next state of the bit register.
//
// Inputs:
//  - c: channel to inform when the update is complete, optional.
func (bit *Bit) Tick(c chan bool) {
	bit.dff.Tick()

	if c != nil {
		c <- true
	}
}
