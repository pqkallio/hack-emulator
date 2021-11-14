package bit

import "github.com/pqkallio/hack-emulator/components"

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

func (or *Or) Update(a, b bool, c chan components.OrderedVal, idx int) bool {
	notA := or.not1.Update(a, nil, 0)
	notB := or.not2.Update(b, nil, 0)
	notaAndNotb := or.and.Update(notA, notB, nil, 0)

	val := or.not3.Update(notaAndNotb, nil, 0)

	if c != nil {
		c <- components.OrderedVal{Idx: idx, Val: val}
	}

	return val
}
