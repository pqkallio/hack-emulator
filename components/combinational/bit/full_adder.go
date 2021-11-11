package bit

type FullAdder struct {
	ha1, ha2, ha3 *HalfAdder
}

func NewFullAdder() *FullAdder {
	return &FullAdder{
		NewHalfAdder(), NewHalfAdder(), NewHalfAdder(),
	}
}

func (fa *FullAdder) Update(a, b, c bool) (bool, bool) {
	abSum, abCarry := fa.ha1.Update(a, b)

	sum, abcCarry := fa.ha2.Update(abSum, c)

	carry, _ := fa.ha3.Update(abcCarry, abCarry)

	return sum, carry
}
