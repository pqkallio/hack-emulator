package bit

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

func (mux *Mux) Update(a, b, sel bool) bool {
	notSel := mux.not.Update(sel)
	aSel := mux.and1.Update(a, notSel)
	bSel := mux.and2.Update(sel, b)

	return mux.or.Update(aSel, bSel)
}
