package word

import "github.com/pqkallio/hack-emulator/hack/components"

// Mux4Way16 is a 4-way 16-bit multiplexer with 1-bit select lines.
type Mux4Way16 struct {
	mux161 *Mux16
	mux162 *Mux16
	mux163 *Mux16
	c      chan components.OrderedVal16
}

func NewMux4Way16() *Mux4Way16 {
	return &Mux4Way16{
		NewMux16(), NewMux16(), NewMux16(),
		make(chan components.OrderedVal16),
	}
}

func (mux4Way16 *Mux4Way16) Update(a, b, c, d uint16, sel0, sel1 bool, ch chan components.OrderedVal16, idx int) uint16 {
	abMux, cdMux := uint16(0), uint16(0)

	go mux4Way16.mux161.Update(a, b, sel0, mux4Way16.c, 0)
	go mux4Way16.mux162.Update(c, d, sel0, mux4Way16.c, 1)

	for i := 0; i < 2; i++ {
		ov := <-mux4Way16.c
		if ov.Idx == 0 {
			abMux = ov.Val
		} else {
			cdMux = ov.Val
		}
	}

	val := mux4Way16.mux163.Update(abMux, cdMux, sel1, nil, 0)

	if ch != nil {
		ch <- components.OrderedVal16{Idx: idx, Val: val}
	}

	return val
}
