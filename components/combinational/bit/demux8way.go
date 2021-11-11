package bit

type Demux8Way struct {
	demux1 *Demux
	demux2 *Demux
	demux3 *Demux
	demux4 *Demux
	demux5 *Demux
	demux6 *Demux
	demux7 *Demux
}

func NewDemux8Way() *Demux8Way {
	return &Demux8Way{
		NewDemux(), NewDemux(), NewDemux(), NewDemux(),
		NewDemux(), NewDemux(), NewDemux(),
	}
}

func (demux8Way *Demux8Way) Update(
	in, sel0, sel1, sel2 bool,
) (bool, bool, bool, bool, bool, bool, bool, bool) {
	aceg, bdfh := demux8Way.demux1.Update(in, sel0)

	ae, cg := demux8Way.demux2.Update(aceg, sel1)
	bf, dh := demux8Way.demux3.Update(bdfh, sel1)

	a, e := demux8Way.demux4.Update(ae, sel2)
	c, g := demux8Way.demux5.Update(cg, sel2)
	b, f := demux8Way.demux6.Update(bf, sel2)
	d, h := demux8Way.demux7.Update(dh, sel2)

	return a, b, c, d, e, f, g, h
}
