package components

type Register struct {
	bits [16]*Bit
}

func NewRegister() *Register {
	bits := [16]*Bit{}

	for i := 0; i < 16; i++ {
		bits[i] = NewBit()
	}

	return &Register{bits}
}

func (reg *Register) Update(opts ...UpdateOpts) Val {
	outVal := uint16(0)

	var load *UpdateOpts

	for _, opt := range opts {
		if opt.target == TargetLoad {
			load = &opt
			break
		}
	}

	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			val := opt.val.GetUint16()

			for i, bit := range reg.bits {
				updateOpts := make([]UpdateOpts, 0, 2)

				if load != nil {
					updateOpts = append(updateOpts, *load)
				}

				updateOpts = append(updateOpts, UpdateOpts{TargetIn, &SingleChan{val&(1<<uint(i)) != 0}})

				bitSet := bit.Update(updateOpts...).GetBool()

				if bitSet {
					outVal |= 1 << i
				}
			}
		}
	}

	return &SixteenChan{outVal}
}

func (reg *Register) Tick() {
	for _, bit := range reg.bits {
		bit.Tick()
	}
}
