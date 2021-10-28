package components

type Inc16 struct {
	add16 *Add16
	one   Val
}

func NewInc16() *Inc16 {
	return &Inc16{NewAdd16(), &SixteenChan{1}}
}

func (i *Inc16) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		return i.add16.Update(
			UpdateOpts{TargetA, opt.val},
			UpdateOpts{TargetB, i.one},
		)
	}

	return &InvalidVal{}
}
