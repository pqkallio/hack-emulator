package components

type Or8Way struct {
	vals uint8
}

func NewOr8Way() *Or8Way {
	return &Or8Way{
		vals: 0,
	}
}

func (or8Way *Or8Way) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		shift := uint8(opt.target)
		val := uint8(0)

		if opt.val.GetBool() {
			val = 1
		}

		if shift < 8 {
			or8Way.vals |= val << shift
		}
	}

	return &SingleChan{or8Way.vals > 0}
}
