package components

type Bit struct {
	dff *DFF
}

func NewBit() *Bit {
	return &Bit{
		NewDFF(),
	}
}

func (bit *Bit) Update(opts ...UpdateOpts) Val {
	var data, load Val

	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			data = opt.val
		case TargetLoad:
			load = opt.val
		}
	}

	if load != nil && data != nil && load.GetBool() {
		return bit.dff.Update(
			UpdateOpts{TargetIn, data},
		)
	}

	return bit.dff.Update()
}

func (bit *Bit) Tick() {
	bit.dff.Tick()
}
