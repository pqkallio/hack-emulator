package bit

// Nand is a combinational circuit that returns the logical NAND of two inputs.
// This is the very base component of the emulator.
//
// Truth table:
// 		A	B	NAND
// 		0	0	1
// 		0	1	1
// 		1	0	1
// 		1	1	0
type Nand struct{}

func NewNand() *Nand {
	return &Nand{}
}

func (n *Nand) Update(a, b bool) bool {
	return !(a && b)
}
