package word

type Mux4Way16 struct {
	mux161 *Mux16
	mux162 *Mux16
	mux163 *Mux16
}

func NewMux4Way16() *Mux4Way16 {
	return &Mux4Way16{
		NewMux16(), NewMux16(), NewMux16(),
	}
}

func (mux4Way16 *Mux4Way16) Update(a, b, c, d uint16, sel0, sel1 bool) uint16 {
	abMux := mux4Way16.mux161.Update(a, b, sel0)
	cdMux := mux4Way16.mux162.Update(c, d, sel0)

	return mux4Way16.mux163.Update(abMux, cdMux, sel1)
}
