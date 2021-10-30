package components

type Mux16 struct {
	a    Val
	b    Val
	sel  Val
	muxs [16]*Mux
}

func NewMux16() *Mux16 {
	muxs := [16]*Mux{}

	for i := 0; i < 16; i++ {
		muxs[i] = NewMux()
	}

	return &Mux16{&InvalidVal{}, &InvalidVal{}, &InvalidVal{}, muxs}
}

func (mux16 *Mux16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			mux16.a = opt.val
		case TargetB:
			mux16.b = opt.val
		case TargetSel0:
			mux16.sel = opt.val
		}
	}

	var out uint16

	for i, mux := range mux16.muxs {
		val := mux.Update(
			UpdateOpts{TargetA, &SingleChan{mux16.a.GetBoolFromUint16(uint16(i))}},
			UpdateOpts{TargetB, &SingleChan{mux16.b.GetBoolFromUint16(uint16(i))}},
			UpdateOpts{TargetSel0, mux16.sel},
		).GetBool()

		if val {
			out |= 1 << i
		}
	}

	return &SixteenChan{out}
}
