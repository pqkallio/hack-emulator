package bit

// Demux, or demultiplexer, is a combinational component that routes the input
// into two outputs depending on the selection input.
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

// Update updates the state of the demux.
//
// in: the input to the demux.
// sel: the selection input to the demux.
//
// out1: in if !sel, false otherwise
// out2: in if sel, false otherwise
func (demux *Demux) Update(in, sel bool) (bool, bool) {
	notSel := demux.not.Update(sel, nil, 0)

	return demux.and1.Update(in, notSel, nil, 0),
		demux.and2.Update(in, sel, nil, 0)
}
