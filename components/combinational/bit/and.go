package bit

import "github.com/pqkallio/hack-emulator/components"

type And struct {
	nand1, nand2, nand3 *Nand
}

func NewAnd() *And {
	return &And{
		NewNand(), NewNand(), NewNand(),
	}
}

func (and *And) Update(a, b bool, c chan components.OrderedVal, idx int) bool {
	nandAB1 := and.nand1.Update(a, b)

	nandAB2 := and.nand2.Update(a, b)

	val := and.nand3.Update(nandAB1, nandAB2)

	if c != nil {
		c <- components.OrderedVal{Idx: idx, Val: val}
	}

	return val
}
