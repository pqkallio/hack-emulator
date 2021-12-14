package hack

import (
	"github.com/pqkallio/hack-emulator/hack/components/combinational/word"
	seq "github.com/pqkallio/hack-emulator/hack/components/sequential/word"
)

// Computer is the HACK computer.
type Computer struct {
	memVal uint16
	rom    *word.ROM32KFlat
	cpu    *seq.CPU
	mmu    *seq.MMU
	reset  *bool
}

// NewComputer creates a new HACK computer.
// input:
//  ram: the RAM to use for the computer
// 	scr: screen memory (read by screen, read/write by CPU)
// 	kbd: keyboard memory (write by keyboard, read by CPU)
// 	rom: ROM memory (read by CPU, contains the Hack program)
func NewComputer(
	ram *seq.Ram16kFlat,
	scr *seq.ScreenMem,
	kbd *word.KeyboardMem,
	rom *word.ROM32KFlat,
	reset *bool,
) *Computer {
	return &Computer{
		memVal: 0,
		rom:    rom,
		cpu:    seq.NewCPU(),
		mmu:    seq.NewMMU(ram, scr, kbd),
		reset:  reset,
	}
}

func (c *Computer) Run() {
	for {
		c.Next(*c.reset)
	}
}

// Next evaluates the next state of the computer.
// input:
// 	reset: reset the computer, setting the program counter to 0
func (c *Computer) Next(reset bool) {
	pc, memOut, memAddr, writeToMem := c.cpu.Fetch() // get the instruction address
	instruction := c.rom.Get(pc)                     // get the instruction
	// decoded := util.DecodeInstruction(instruction)       // decode the instruction
	// log.Printf("%04X: %s", pc, decoded)                  // log the instruction
	c.memVal = c.mmu.Update(memOut, memAddr, writeToMem) // update memory
	c.cpu.Execute(instruction, c.memVal, reset)          // execute the instruction
	c.cpu.Tick()                                         // clock pulse for CPU
	c.mmu.Tick()                                         // clock pulse for MMU
}

// LoadROM loads a new ROM into the computer.
func (c *Computer) LoadROM(rom *word.ROM32KFlat) {
	c.rom = rom
}
