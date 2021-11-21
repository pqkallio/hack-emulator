package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/components/combinational/word"
)

// MMU is the memory management unit that maps address space 0x0000-0x6000 in the following manner:
// 0x0000-0x3FFF:   RAM
// 0x4000-0x5FFF:   Screen memory
// 0x6000:          Keyboard memory
// Trying to access addresses outside of this range will cause undefined behaviour.
type MMU struct {
	ram        *Ram16kFlat
	scr        *ScreenMem
	kbd        *word.KeyboardMem
	notAddr14  *bit.Not
	notAddr13  *bit.Not
	loadRamAnd *bit.And
	loadIoAnd1 *bit.And
	loadIoAnd2 *bit.And
	mux        *word.Mux4Way16
}

func NewMMU(ram *Ram16kFlat, scr *ScreenMem, kbd *word.KeyboardMem) *MMU {
	return &MMU{
		ram, scr, kbd,
		bit.NewNot(), bit.NewNot(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
		word.NewMux4Way16(),
	}
}

func (m *MMU) Update(in, addr uint16, load bool) uint16 {
	addr14 := (addr & (1 << 14)) != 0
	addr13 := (addr & (1 << 13)) != 0

	notAddr14 := m.notAddr14.Update(addr14, nil, 0)
	notAddr13 := m.notAddr13.Update(addr13, nil, 0)

	loadRam := m.loadRamAnd.Update(load, notAddr14, nil, 0)
	loadIo1 := m.loadIoAnd1.Update(notAddr13, addr14, nil, 0)
	loadIo2 := m.loadIoAnd2.Update(load, loadIo1, nil, 0)

	addr0 := addr&1 != 0
	addr1 := addr&2 != 0
	addr2 := addr&4 != 0
	addr3 := addr&8 != 0
	addr4 := addr&16 != 0
	addr5 := addr&32 != 0
	addr6 := addr&64 != 0
	addr7 := addr&128 != 0
	addr8 := addr&256 != 0
	addr9 := addr&512 != 0
	addr10 := addr&1024 != 0
	addr11 := addr&2048 != 0
	addr12 := addr&4096 != 0

	ramVal := m.ram.Update(
		in, loadRam,
		addr0, addr1, addr2, addr3,
		addr4, addr5, addr6, addr7,
		addr8, addr9, addr10, addr11,
		addr12, addr13, nil, 0,
	)

	scrVal := m.scr.Update(in, addr, loadIo2)

	return m.mux.Update(
		ramVal, ramVal, scrVal, m.kbd.Get(),
		addr13, addr14,
		nil, 0,
	)
}

func (m *MMU) Tick() {
	m.ram.Tick(nil)
}
