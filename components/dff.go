package components

type DFF struct {
	curr, next SingleChan
}

func NewDFF() *DFF {
	return &DFF{
		SingleChan{}, SingleChan{},
	}
}

func (dff *DFF) Update(opts ...UpdateOpts) Val {
	for _, opt := range opts {
		switch opt.target {
		case TargetIn:
			if singleChan, ok := opt.val.(*SingleChan); ok {
				dff.next = *singleChan
			}
		}
	}

	return &dff.curr
}

func (dff *DFF) Tick() {
	dff.curr = dff.next
}
