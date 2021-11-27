package word

import "github.com/pqkallio/hack-emulator/components/combinational/word"

// Computer is the HACK computer.
type Computer struct {
	memVal uint16
	rom    *word.ROM32KFlat
	cpu    *CPU
	mmu    *MMU
}

// NewComputer creates a new HACK computer.
// input:
//  ram: the RAM to use for the computer
// 	scr: screen memory (read by screen, read/write by CPU)
// 	kbd: keyboard memory (write by keyboard, read by CPU)
// 	rom: ROM memory (read by CPU, contains the Hack program)
func NewComputer(ram *Ram16kFlat, scr *ScreenMem, kbd *word.KeyboardMem, rom *word.ROM32KFlat) *Computer {
	return &Computer{
		memVal: 0,
		rom:    rom,
		cpu:    NewCPU(),
		mmu:    NewMMU(ram, scr, kbd),
	}
}

// Next evaluates the next state of the computer.
// input:
// 	reset: reset the computer, setting the program counter to 0
func (c *Computer) Next(reset bool) {
	pc := c.cpu.Fetch()                                                        // get the instruction address
	instruction := c.rom.Get(pc)                                               // get the instruction
	memOut, memAddr, writeToMem := c.cpu.Execute(instruction, c.memVal, reset) // execute the instruction
	c.memVal = c.mmu.Update(memOut, memAddr, writeToMem)                       // update memory
	c.cpu.Tick()                                                               // clock pulse for CPU
	c.mmu.Tick()                                                               // clock pulse for MMU
}

// LoadROM loads a new ROM into the computer.
func (c *Computer) LoadROM(rom *word.ROM32KFlat) {
	c.rom = rom
}
