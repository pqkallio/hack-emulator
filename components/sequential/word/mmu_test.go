package word

import (
	"testing"

	"github.com/pqkallio/hack-emulator/components"
	"github.com/pqkallio/hack-emulator/components/combinational/word"
)

func TestMMU(t *testing.T) {
	t.Parallel()

	ramStart := components.Addr16K(0)
	ramEnd := components.Addr16K(0x3FFF)
	scrStart := uint16(0x4000)
	scrEnd := uint16(0x5FFF)
	kbdAddr := uint16(0x6000)

	ramStartAddrLines := ramStart.ToAddressLines()
	ramEndAddrLines := ramEnd.ToAddressLines()

	ram := NewRam16kFlat()
	scr := NewScreenMem()
	kbd := word.NewKeyboardMem()

	scr.Update(84, scrStart, true)
	scr.Update(85, scrEnd, true)
	kbd.Update(42)

	ram.Update(
		420, true,
		ramStartAddrLines[0], ramStartAddrLines[1], ramStartAddrLines[2], ramStartAddrLines[3],
		ramStartAddrLines[4], ramStartAddrLines[5], ramStartAddrLines[6], ramStartAddrLines[7],
		ramStartAddrLines[8], ramStartAddrLines[9], ramStartAddrLines[10], ramStartAddrLines[11],
		ramStartAddrLines[12], ramStartAddrLines[13], nil, 0,
	)
	ram.Update(
		421, true,
		ramEndAddrLines[0], ramEndAddrLines[1], ramEndAddrLines[2], ramEndAddrLines[3],
		ramEndAddrLines[4], ramEndAddrLines[5], ramEndAddrLines[6], ramEndAddrLines[7],
		ramEndAddrLines[8], ramEndAddrLines[9], ramEndAddrLines[10], ramEndAddrLines[11],
		ramEndAddrLines[12], ramEndAddrLines[13], nil, 0,
	)

	mmu := NewMMU(ram, scr, kbd)

	expected := uint16(42)
	actual := mmu.Update(77, kbdAddr, true)

	// Check that kbd is not updated and always returns its current value.
	if actual != expected {
		t.Errorf("KBD: expected %d, got %d", expected, actual)
	}

	expected = 0
	actual = mmu.Update(84, scrStart, false)

	// Check that if the screen memory is not ticked, it will return 0 at first.
	if actual != expected {
		t.Errorf("SCR start: expected %d, got %d", expected, actual)
	}

	actual = mmu.Update(85, scrEnd, false)

	// Check that if the screen memory is not ticked, it will return 0 at first.
	if actual != expected {
		t.Errorf("SCR end: expected %d, got %d", expected, actual)
	}

	actual = mmu.Update(420, uint16(ramStart), false)

	// Check that if the MMU is not ticked, RAM returns 0 at first.
	if actual != expected {
		t.Errorf("RAM start: expected %d, got %d", expected, actual)
	}

	actual = mmu.Update(421, uint16(ramEnd), false)

	// Check that if the MMU is not ticked, RAM returns 0 at first.
	if actual != expected {
		t.Errorf("RAM end: expected %d, got %d", expected, actual)
	}

	mmu.Tick()

	expected = 42
	actual = mmu.Update(77, kbdAddr, true)

	// Check that keyboard value has not been updated after the MMU tick and attempt
	// to load a value to it.
	if actual != expected {
		t.Errorf("KBD: expected %d, got %d", expected, actual)
	}

	expected = 0
	actual = mmu.Update(21, scrStart, false)

	// Check that if the screen is not ticked, it will still return 0.
	if actual != expected {
		t.Errorf("SCR start: expected %d, got %d", expected, actual)
	}

	actual = mmu.Update(22, scrEnd, false)

	// Check that if the screen is not ticked, it will still return 0.
	if actual != expected {
		t.Errorf("SCR end: expected %d, got %d", expected, actual)
	}

	expected = 420
	actual = mmu.Update(9, uint16(ramStart), true)

	// Check that the RAM has updated after an MMU tick.
	if actual != expected {
		t.Errorf("RAM start: expected %d, got %d", expected, actual)
	}

	expected = 421
	actual = mmu.Update(9, uint16(ramEnd), true)

	// Check that the RAM has updated after an MMU tick.
	if actual != expected {
		t.Errorf("RAM end: expected %d, got %d", expected, actual)
	}

	scr.Tick()

	expected = 84
	actual = mmu.Update(21, scrStart, true)

	// Check that the screen memory has updated after screen memory's been ticked.
	if actual != expected {
		t.Errorf("SCR start: expected %d, got %d", expected, actual)
	}

	expected = 85
	actual = mmu.Update(21, scrEnd, true)

	// Check that the screen memory has updated after screen memory's been ticked.
	if actual != expected {
		t.Errorf("SCR end: expected %d, got %d", expected, actual)
	}
}
