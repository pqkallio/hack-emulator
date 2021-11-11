package bit

type Demux struct {
	not  *Not
	and1 *And
	and2 *And
}

func NewDemux() *Demux {
	return &Demux{
		NewNot(),
		NewAnd(), NewAnd(),
	}
}

func (demux *Demux) Update(in, sel bool) (bool, bool) {
	notSel := demux.not.Update(sel)

	return demux.and1.Update(in, notSel),
		demux.and2.Update(in, sel)
}
