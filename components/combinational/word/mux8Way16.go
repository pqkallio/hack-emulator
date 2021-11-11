package word

type Mux8Way16 struct {
	mux4Way161 *Mux4Way16
	mux4Way162 *Mux4Way16
	mux16      *Mux16
}

func NewMux8Way16() *Mux8Way16 {
	return &Mux8Way16{
		NewMux4Way16(), NewMux4Way16(),
		NewMux16(),
	}
}

func (mux8Way16 *Mux8Way16) Update(
	a, b, c, d, e, f, g, h uint16,
	sel0, sel1, sel2 bool,
) uint16 {
	abcdMux := mux8Way16.mux4Way161.Update(
		a,
		b,
		c,
		d,
		sel0,
		sel1,
	)
	efghMux := mux8Way16.mux4Way161.Update(
		e,
		f,
		g,
		h,
		sel0,
		sel1,
	)

	return mux8Way16.mux16.Update(
		abcdMux,
		efghMux,
		sel2,
	)
}
