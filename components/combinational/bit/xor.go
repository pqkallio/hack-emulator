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
	nota := xor.not1.Update(a, nil, 0)
	notb := xor.not1.Update(b, nil, 0)
	aAndNotb := xor.and1.Update(a, notb, nil, 0)
	notaAndb := xor.and2.Update(nota, b, nil, 0)

	return xor.or.Update(aAndNotb, notaAndb, nil, 0)
}
