package bit

import "github.com/pqkallio/hack-emulator/components"

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
