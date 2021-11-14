package word

import "github.com/pqkallio/hack-emulator/components/combinational/word"

type PC struct {
	inc16  *word.Inc16
	mux161 *word.Mux16
	mux162 *word.Mux16
	mux163 *word.Mux16
	reg    *Register
}

func NewPC() *PC {
	return &PC{
		word.NewInc16(),
		word.NewMux16(), word.NewMux16(), word.NewMux16(),
		NewRegister(),
	}
}

func (pc *PC) Update(in uint16, load, inc, reset bool) uint16 {
	prev := pc.reg.Update(0, false, nil, 0)
	incd := pc.inc16.Update(prev)
	postInc := pc.mux161.Update(prev, incd, inc)
	postLoad := pc.mux162.Update(postInc, in, load)
	postReset := pc.mux163.Update(postLoad, 0, reset)
	out := pc.reg.Update(postReset, true, nil, 0)

	return out
}

func (pc *PC) Tick() {
	pc.reg.Tick()
}
