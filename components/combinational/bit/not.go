package bit

type Not struct {
	nand *Nand
}

func NewNot() *Not {
	return &Not{NewNand()}
}

func (not *Not) Update(in bool) bool {
	return not.nand.Update(in, in)
}
