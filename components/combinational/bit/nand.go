package bit

type Nand struct{}

func NewNand() *Nand {
	return &Nand{}
}

func (n *Nand) Update(a, b bool) bool {
	return !(a && b)
}
