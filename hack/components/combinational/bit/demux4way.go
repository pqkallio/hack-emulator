package bit

// Demux4Way is a 4-way demultiplexer.
type Demux4Way struct {
	demux1 *Demux
	demux2 *Demux
	demux3 *Demux
}

func NewDemux4Way() *Demux4Way {
	return &Demux4Way{
		NewDemux(), NewDemux(), NewDemux(),
	}
}

// Update evaluates the circuit.
//
// Inputs:
// 	in: input bit
// 	sel0: select bit 0
// 	sel1: select bit 1
//
// Output:
//  a: in if !sel0 and !sel1, false otherwise
//  b: in if sel0 and !sel1, false otherwise
//  c: in if !sel0 and sel1, false otherwise
//  d: in if sel0 and sel1, false otherwise
func (demux4Way *Demux4Way) Update(in, sel0, sel1 bool) (bool, bool, bool, bool) {
	demux1a, demux1b := demux4Way.demux1.Update(in, sel0)

	a, c := demux4Way.demux2.Update(demux1a, sel1)
	b, d := demux4Way.demux3.Update(demux1b, sel1)

	return a, b, c, d
}
