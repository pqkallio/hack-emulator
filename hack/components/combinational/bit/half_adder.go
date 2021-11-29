package bit

// HalfAdder is a combinational component that implements a half adder.
type HalfAdder struct {
	xor *Xor
	and *And
}

func NewHalfAdder() *HalfAdder {
	return &HalfAdder{
		NewXor(), NewAnd(),
	}
}

// Update evaluates the half adder.
//
// Input:
// 	a, b: input bits
//
// Output:
//  sum: sum of inputs
//  carry: carry bit after addition
func (ha *HalfAdder) Update(a, b bool) (bool, bool) {
	return ha.xor.Update(a, b),
		ha.and.Update(a, b, nil, 0)
}
