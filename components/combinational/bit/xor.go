package bit

type Xor struct {
	not1 *Not
	not2 *Not
	and1 *And
	and2 *And
	or   *Or
}

func NewXor() *Xor {
	return &Xor{
		NewNot(), NewNot(),
		NewAnd(), NewAnd(),
		NewOr(),
	}
}

func (xor *Xor) Update(a, b bool) bool {
	nota := xor.not1.Update(a)
	notb := xor.not1.Update(b)
	aAndNotb := xor.and1.Update(a, notb)
	notaAndb := xor.and2.Update(nota, b)

	return xor.or.Update(aAndNotb, notaAndb)
}
