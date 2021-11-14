package bit

type HalfAdder struct {
	xor *Xor
	and *And
}

func NewHalfAdder() *HalfAdder {
	return &HalfAdder{
		NewXor(), NewAnd(),
	}
}

func (ha *HalfAdder) Update(a, b bool) (bool, bool) {
	return ha.xor.Update(a, b),
		ha.and.Update(a, b)
}
