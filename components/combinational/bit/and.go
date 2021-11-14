package bit

type And struct {
	nand1, nand2, nand3 *Nand
}

func NewAnd() *And {
	return &And{
		NewNand(), NewNand(), NewNand(),
	}
}

func (and *And) Update(a, b bool) bool {
	nandAB1 := and.nand1.Update(a, b)

	nandAB2 := and.nand2.Update(a, b)

	return and.nand3.Update(nandAB1, nandAB2)
}
