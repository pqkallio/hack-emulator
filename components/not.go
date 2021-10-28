package components

type Not struct {
}

func (not *Not) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		return &SingleChan{val: !opt.val.GetBool()}
	}

	return &InvalidVal{}
}
