package bit

type DFF struct {
	curr, next bool
}

func NewDFF() *DFF {
	return &DFF{false, false}
}

func (dff *DFF) Update(in bool) bool {
	dff.next = in

	return dff.curr
}

func (dff *DFF) Tick() {
	dff.curr = dff.next
}
