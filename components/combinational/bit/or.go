package bit

type Or struct {
	not1 *Not
	not2 *Not
	not3 *Not
	and  *And
}

func NewOr() *Or {
	return &Or{
		NewNot(), NewNot(), NewNot(),
		NewAnd(),
	}
}

func (or *Or) Update(a, b bool) bool {
	notA := or.not1.Update(a)
	notB := or.not2.Update(b)
	notaAndNotb := or.and.Update(notA, notB)

	return or.not3.Update(notaAndNotb)
}
