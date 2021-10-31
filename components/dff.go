package components

type DFF struct {
	data Val
}

func NewDFF() *DFF {
	return &DFF{
		&InvalidVal{},
	}
}

func (dff *DFF) Update(opts ...UpdateOpts) {
	var data, load Val

	for _, opt := range opts {
		switch opt.target {
		case TargetData:
			data = opt.val
		case TargetLoad:
			load = opt.val
		}
	}

	if data == nil || load == nil {
		return
	}

	if load.GetBool() {
		dff.data = data
	}
}

func (dff *DFF) Get() Val {
	return dff.data
}

const (
	TargetLoad Target = iota + 200
	TargetData
)
