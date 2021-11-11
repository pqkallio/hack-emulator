package bit

type Bit struct {
	dff *DFF
}

func NewBit() *Bit {
	return &Bit{
		NewDFF(),
	}
}

func (bit *Bit) Update(data, load bool) bool {
	if load {
		return bit.dff.Update(data)
	}

	return bit.dff.Update(bit.dff.curr)
}

func (bit *Bit) Tick() {
	bit.dff.Tick()
}
