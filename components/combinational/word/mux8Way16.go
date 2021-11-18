package word

import "github.com/pqkallio/hack-emulator/components"

// Mux8Way16 is a 8-way 16-bit multiplexer with 1-bit select lines.
type Mux8Way16 struct {
	mux4Way161 *Mux4Way16
	mux4Way162 *Mux4Way16
	mux16      *Mux16
	c          chan components.OrderedVal16
}

func NewMux8Way16() *Mux8Way16 {
	return &Mux8Way16{
		NewMux4Way16(), NewMux4Way16(),
		NewMux16(),
		make(chan components.OrderedVal16),
	}
}

func (mux8Way16 *Mux8Way16) Update(
	a, b, c, d, e, f, g, h uint16,
	sel0, sel1, sel2 bool,
) uint16 {
	abcdMux, efghMux := uint16(0), uint16(0)

	go mux8Way16.mux4Way161.Update(
		a,
		b,
		c,
		d,
		sel0,
		sel1,
		mux8Way16.c, 0,
	)
	go mux8Way16.mux4Way162.Update(
		e,
		f,
		g,
		h,
		sel0,
		sel1,
		mux8Way16.c, 1,
	)

	for i := 0; i < 2; i++ {
		val := <-mux8Way16.c
		if val.Idx == 0 {
			abcdMux = val.Val
		} else {
			efghMux = val.Val
		}
	}

	return mux8Way16.mux16.Update(
		abcdMux,
		efghMux,
		sel2,
		nil, 0,
	)
}
