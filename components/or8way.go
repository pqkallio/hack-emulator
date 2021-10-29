package components

type Or8Way struct {
	vals [8]Val
	or1  *Or
	or2  *Or
	or3  *Or
	or4  *Or
	or5  *Or
	or6  *Or
	or7  *Or
}

func NewOr8Way() *Or8Way {
	vals := [8]Val{
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{},
	}

	return &Or8Way{
		vals,
		NewOr(), NewOr(),
		NewOr(), NewOr(),
		NewOr(), NewOr(),
		NewOr(),
	}
}

func (or8Way *Or8Way) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetA:
			fallthrough
		case TargetB:
			fallthrough
		case TargetC:
			fallthrough
		case TargetD:
			fallthrough
		case TargetE:
			fallthrough
		case TargetF:
			fallthrough
		case TargetG:
			fallthrough
		case TargetH:
			or8Way.vals[opt.target] = opt.val
		}
	}

	temp1 := or8Way.or1.Update(
		UpdateOpts{TargetA, or8Way.vals[0]},
		UpdateOpts{TargetB, or8Way.vals[1]},
	)
	temp2 := or8Way.or2.Update(
		UpdateOpts{TargetA, temp1},
		UpdateOpts{TargetB, or8Way.vals[2]},
	)
	temp3 := or8Way.or3.Update(
		UpdateOpts{TargetA, temp2},
		UpdateOpts{TargetB, or8Way.vals[3]},
	)
	temp4 := or8Way.or4.Update(
		UpdateOpts{TargetA, temp3},
		UpdateOpts{TargetB, or8Way.vals[4]},
	)
	temp5 := or8Way.or5.Update(
		UpdateOpts{TargetA, temp4},
		UpdateOpts{TargetB, or8Way.vals[5]},
	)
	temp6 := or8Way.or6.Update(
		UpdateOpts{TargetA, temp5},
		UpdateOpts{TargetB, or8Way.vals[6]},
	)
	return or8Way.or7.Update(
		UpdateOpts{TargetA, temp6},
		UpdateOpts{TargetB, or8Way.vals[7]},
	)
}
