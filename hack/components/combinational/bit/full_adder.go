package bit

// FullAdder is a combinational circuit that implements a full adder.
type FullAdder struct {
	ha1, ha2, ha3 *HalfAdder
}

func NewFullAdder() *FullAdder {
	return &FullAdder{
		NewHalfAdder(), NewHalfAdder(), NewHalfAdder(),
	}
}

// Update evaluates the circuit.
//
// Inputs:
// 	a, b: bits to add.
// 	c: carry bit from the previous adder.
//
// Outputs:
// 	sum: sum of a and b.
// 	carry: carry bit of the sum.
func (fa *FullAdder) Update(a, b, c bool) (bool, bool) {
	abSum, abCarry := fa.ha1.Update(a, b)

	sum, abcCarry := fa.ha2.Update(abSum, c)

	carry, _ := fa.ha3.Update(abcCarry, abCarry)

	return sum, carry
}
