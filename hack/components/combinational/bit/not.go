package bit

import "github.com/pqkallio/hack-emulator/hack/components"

// Not is a component that performs a logical NOT operation on a single input.
//
// Truth table:
// 		IN	NOT
// 		0	1
// 		1	0
type Not struct {
	nand *Nand
}

func NewNot() *Not {
	return &Not{NewNand()}
}

func (not *Not) Update(in bool, c chan components.OrderedVal, idx int) bool {
	val := not.nand.Update(in, in)

	if c != nil {
		c <- components.OrderedVal{Idx: idx, Val: val}
	}

	return val
}
