package word

type ScreenMem struct {
	mem   [8192]uint16
	dirty [][2]uint16
}

func NewScreenMem() *ScreenMem {
	return &ScreenMem{dirty: make([][2]uint16, 0, 8192)}
}

func (sm *ScreenMem) Update(in, addr uint16, load bool) uint16 {
	// Use only the lower 13 bits of the address.
	addr = addr & 8191

	if load {
		sm.dirty = append(sm.dirty, [2]uint16{addr, in})
	}

	return sm.mem[addr]
}

func (sm *ScreenMem) Read(addr uint16) uint16 {
	return sm.mem[addr]
}

func (sm *ScreenMem) Tick() {
	for _, entry := range sm.dirty {
		sm.mem[entry[0]] = entry[1]
	}

	sm.dirty = sm.dirty[:0]
}

func (sm *ScreenMem) GetMem() [8192]uint16 {
	return sm.mem
}

func (sm *ScreenMem) GetDirty() [][2]uint16 {
	return sm.dirty
}
