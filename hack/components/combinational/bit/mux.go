package bit

import "github.com/pqkallio/hack-emulator/hack/components"

// Mux, or a mutliplexer is a combinational component that selects between two inputs.
type Mux struct {
	not  *Not
	and1 *And
	and2 *And
	or   *Or
}

func NewMux() *Mux {
	return &Mux{
		NewNot(),
		NewAnd(), NewAnd(),
		NewOr(),
	}
}

// Update updates the component.
// In:
// 	a - input a
// 	b - input b
// 	sel - selection
// Out:
// 	val - a if !sel, b otherwise
func (mux *Mux) Update(a, b, sel bool, c chan components.OrderedVal, idx int) bool {
	notSel := mux.not.Update(sel, nil, 0)
	aSel := mux.and1.Update(a, notSel, nil, 0)
	bSel := mux.and2.Update(sel, b, nil, 0)

	val := mux.or.Update(aSel, bSel, nil, 0)

	if c != nil {
		c <- components.OrderedVal{Val: val, Idx: idx}
	}

	return val
}
