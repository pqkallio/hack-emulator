package word

import (
	"github.com/pqkallio/hack-emulator/hack/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/hack/components/combinational/word"
)

// CPU is a sequential component that implements the CPU of the Hack computer.
type CPU struct {
	prevAluOut         uint16
	instructionMux     *word.Mux16
	aRegister          *Register
	dRegister          *Register
	memMux             *word.Mux16
	alu                *word.ALU
	pc                 *PC
	jmpNegAnd          *bit.And
	jmpZeroAnd         *bit.And
	jmpPosAnd          *bit.And
	notNg              *bit.Not
	notZr              *bit.Not
	posAnd             *bit.And
	posOrZero          *bit.Or
	negOrPos           *bit.Or
	negOrZero          *bit.Or
	zeroOrPosOrNeg     *bit.Or
	jumpOr             *bit.Or8Way
	notAInstr          *bit.Not
	cInstrAndD1        *bit.And
	cInstrAndD2        *bit.And
	cInstrAndD3        *bit.And
	cInstrAndMemAccess *bit.And
	aOrCInstrAndD1     *bit.Or
	cInstrAndOpCodeZx  *bit.And
	cInstrAndOpCodeNx  *bit.And
	cInstrAndOpCodeZy  *bit.And
	cInstrAndOpCodeNy  *bit.And
	cInstrAndOpCodeF   *bit.And
	cInstrAndOpCodeNO  *bit.And
	cInstrAndZr        *bit.And
	cInstrAndNg        *bit.And
	cInstrAndJ1        *bit.And
	cInstrAndJ2        *bit.And
	cInstrAndJ3        *bit.And
	memOut             uint16
	loadMem            bool
}

func NewCPU() *CPU {
	return &CPU{
		0,
		word.NewMux16(),
		NewRegister(), NewRegister(),
		word.NewMux16(),
		word.NewALU(),
		NewPC(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
		bit.NewNot(), bit.NewNot(),
		bit.NewAnd(),
		bit.NewOr(), bit.NewOr(), bit.NewOr(), bit.NewOr(),
		bit.NewOr8Way(),
		bit.NewNot(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
		bit.NewOr(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
		0,
		false,
	}
}

// Execute evaluates the CPU's next state and returns the current state.
//
// Inputs:
//  - instruction: The instruction to be executed.
//  - in: The input value from the memory.
//  - reset: The reset signal. If true, PC will be set to 0.
func (cpu *CPU) Execute(instruction, in uint16, reset bool) {
	// 1. Extract lines from instruction.

	l := getLines(instruction)

	// 2. Evaluate A register.

	isAInstr := cpu.notAInstr.Update(l.isCInstr, nil, 0)

	aInput := cpu.instructionMux.Update(cpu.prevAluOut, instruction, isAInstr, nil, 0)

	isCInstrAndD1 := cpu.cInstrAndD1.Update(l.isCInstr, l.d1, nil, 0)
	isAInstrOrCInstrAndD1 := cpu.aOrCInstrAndD1.Update(isAInstr, isCInstrAndD1, nil, 0)
	aRegVal := cpu.aRegister.Update(aInput, isAInstrOrCInstrAndD1, nil, 0)

	isCInstrAndMemAccess := cpu.cInstrAndMemAccess.Update(l.isCInstr, l.memAccess, nil, 0)
	aluAVal := cpu.memMux.Update(aRegVal, in, isCInstrAndMemAccess, nil, 0)

	// 3. Evaluate D register.

	isCInstrAndD2 := cpu.cInstrAndD2.Update(l.isCInstr, l.d2, nil, 0)
	aluDVal := cpu.dRegister.Update(cpu.prevAluOut, isCInstrAndD2, nil, 0)

	// 4. Evaluate ALU.

	opCodeZx := cpu.cInstrAndOpCodeZx.Update(l.isCInstr, l.zx, nil, 0)
	opCodeNx := cpu.cInstrAndOpCodeNx.Update(l.isCInstr, l.nx, nil, 0)
	opCodeZy := cpu.cInstrAndOpCodeZy.Update(l.isCInstr, l.zy, nil, 0)
	opCodeNy := cpu.cInstrAndOpCodeNy.Update(l.isCInstr, l.ny, nil, 0)
	opCodeF := cpu.cInstrAndOpCodeF.Update(l.isCInstr, l.f, nil, 0)
	opCodeNO := cpu.cInstrAndOpCodeNO.Update(l.isCInstr, l.no, nil, 0)

	result, zr, ng := cpu.alu.Update(aluDVal, aluAVal, opCodeZx, opCodeNx, opCodeZy, opCodeNy, opCodeF, opCodeNO)
	cpu.prevAluOut = result

	// 5. Evaluate jump and program counter.

	isCInstrAndZr := cpu.cInstrAndZr.Update(l.isCInstr, zr, nil, 0)
	isCInstrAndNg := cpu.cInstrAndNg.Update(l.isCInstr, ng, nil, 0)
	isNotNg := cpu.notNg.Update(isCInstrAndNg, nil, 0)
	isNotZr := cpu.notZr.Update(isCInstrAndZr, nil, 0)
	isPos := cpu.posAnd.Update(isNotZr, isNotNg, nil, 0)
	isCInstrAndJ1 := cpu.cInstrAndJ1.Update(l.isCInstr, l.j1, nil, 0)
	isCInstrAndJ2 := cpu.cInstrAndJ2.Update(l.isCInstr, l.j2, nil, 0)
	isCInstrAndJ3 := cpu.cInstrAndJ3.Update(l.isCInstr, l.j3, nil, 0)

	jlt := cpu.jmpNegAnd.Update(isCInstrAndJ1, isCInstrAndNg, nil, 0)
	jeq := cpu.jmpZeroAnd.Update(isCInstrAndJ2, isCInstrAndZr, nil, 0)
	jgt := cpu.jmpPosAnd.Update(isCInstrAndJ3, isPos, nil, 0)
	jge := cpu.posOrZero.Update(jgt, jeq, nil, 0)
	jne := cpu.negOrPos.Update(jlt, jgt, nil, 0)
	jle := cpu.negOrZero.Update(jlt, jeq, nil, 0)
	jmp := cpu.zeroOrPosOrNeg.Update(jge, jlt, nil, 0)
	pcJump := cpu.jumpOr.Update(jmp, jne, jle, jeq, jgt, jge, jlt, false)

	// 6. Evaluate A register after ALU operation.

	aInput = cpu.instructionMux.Update(cpu.prevAluOut, instruction, isAInstr, nil, 0)

	isCInstrAndD1 = cpu.cInstrAndD1.Update(l.isCInstr, l.d1, nil, 0)
	isAInstrOrCInstrAndD1 = cpu.aOrCInstrAndD1.Update(isAInstr, isCInstrAndD1, nil, 0)
	_ = cpu.aRegister.Update(aInput, isAInstrOrCInstrAndD1, nil, 0)

	// 7. Evaluate D register after ALU operation.

	isCInstrAndD2 = cpu.cInstrAndD2.Update(l.isCInstr, l.d2, nil, 0)
	_ = cpu.dRegister.Update(cpu.prevAluOut, isCInstrAndD2, nil, 0)

	// 8. Return.

	_ = cpu.pc.Update(aRegVal, pcJump, true, reset)
	cpu.loadMem = cpu.cInstrAndD3.Update(l.isCInstr, l.d3, nil, 0)
	cpu.memOut = result

	return
}

func (cpu *CPU) Fetch() (pc, memOut, memAddr uint16, loadMem bool) {
	return cpu.pc.Update(0, false, false, false),
		cpu.memOut,
		cpu.aRegister.Update(0, false, nil, 0),
		cpu.loadMem
}

func (cpu *CPU) Tick() {
	cpu.aRegister.Tick(nil)
	cpu.dRegister.Tick(nil)
	cpu.pc.Tick()
}

type lines struct {
	isCInstr              bool
	memAccess             bool
	j1, j2, j3            bool
	d1, d2, d3            bool
	zx, nx, zy, ny, f, no bool
}

func getLines(instruction uint16) (l lines) {
	// A or C instruction?
	l.isCInstr = instruction&0x8000 != 0

	// Destination flags.
	l.d1 = instruction&0x20 != 0
	l.d2 = instruction&0x10 != 0
	l.d3 = instruction&0x08 != 0

	// Jump flags.
	l.j1 = instruction&0x4 != 0
	l.j2 = instruction&0x2 != 0
	l.j3 = instruction&0x1 != 0

	// Memory access flag.
	l.memAccess = instruction&0x1000 != 0

	// C instruction opcode.
	l.zx = instruction&0x800 != 0
	l.nx = instruction&0x400 != 0
	l.zy = instruction&0x200 != 0
	l.ny = instruction&0x100 != 0
	l.f = instruction&0x80 != 0
	l.no = instruction&0x40 != 0

	return
}
