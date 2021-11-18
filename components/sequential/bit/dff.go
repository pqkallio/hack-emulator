package bit

// DFF is a single-bit D-type flip-flop.
type DFF struct {
	curr, next bool
}

func NewDFF() *DFF {
	return &DFF{false, false}
}

// Update evaluates the next state of the DFF and returns its current state.
func (dff *DFF) Update(in bool) bool {
	dff.next = in

	return dff.curr
}

// Tick advances the DFF to the next state.
func (dff *DFF) Tick() {
	dff.curr = dff.next
}
