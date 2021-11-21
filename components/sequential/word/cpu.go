package word

import (
	"github.com/pqkallio/hack-emulator/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/components/combinational/word"
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
	posXor             *bit.Xor
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
		bit.NewNot(),
		bit.NewXor(),
		bit.NewOr(), bit.NewOr(), bit.NewOr(), bit.NewOr(),
		bit.NewOr8Way(),
		bit.NewNot(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
		bit.NewOr(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
		bit.NewAnd(), bit.NewAnd(), bit.NewAnd(), bit.NewAnd(),
	}
}

// Update evaluates the CPU's next state and returns the current state.
//
// Inputs:
//  - instruction: The instruction to be executed.
//  - in: The input value from the memory.
//  - reset: The reset signal. If true, PC will be set to 0.
//
// Outputs:
//  - memOut: The output value to be written to memory.
//  - memAddr: The memory address to write to.
//  - pc: The current value of the PC.
func (cpu *CPU) Update(
	instruction, in uint16,
	reset bool,
) (memOut, memAddr, pc uint16, loadMem bool) {
	isCInstr := instruction&0x8000 != 0
	d1 := instruction&0x20 != 0
	d2 := instruction&0x10 != 0
	d3 := instruction&0x08 != 0

	j1 := instruction&0x4 != 0
	j2 := instruction&0x2 != 0
	j3 := instruction&0x1 != 0

	a := instruction&0x1000 != 0

	zx := instruction&0x800 != 0
	nx := instruction&0x400 != 0
	zy := instruction&0x200 != 0
	ny := instruction&0x100 != 0
	f := instruction&0x80 != 0
	no := instruction&0x40 != 0

	isAInstr := cpu.notAInstr.Update(isCInstr, nil, 0)

	aInput := cpu.instructionMux.Update(cpu.prevAluOut, instruction, isAInstr, nil, 0)

	isCInstrAndD1 := cpu.cInstrAndD1.Update(isCInstr, d1, nil, 0)
	isAInstrOrCInstrAndD1 := cpu.aOrCInstrAndD1.Update(isAInstr, isCInstrAndD1, nil, 0)
	aRegVal := cpu.aRegister.Update(aInput, isAInstrOrCInstrAndD1, nil, 0)

	isCInstrAndMemAccess := cpu.cInstrAndMemAccess.Update(isCInstr, a, nil, 0)
	aluAVal := cpu.memMux.Update(aRegVal, in, isCInstrAndMemAccess, nil, 0)

	isCInstrAndD2 := cpu.cInstrAndD2.Update(isCInstr, d2, nil, 0)
	aluDVal := cpu.dRegister.Update(cpu.prevAluOut, isCInstrAndD2, nil, 0)

	opCodeZx := cpu.cInstrAndOpCodeZx.Update(isCInstr, zx, nil, 0)
	opCodeNx := cpu.cInstrAndOpCodeNx.Update(isCInstr, nx, nil, 0)
	opCodeZy := cpu.cInstrAndOpCodeZy.Update(isCInstr, zy, nil, 0)
	opCodeNy := cpu.cInstrAndOpCodeNy.Update(isCInstr, ny, nil, 0)
	opCodeF := cpu.cInstrAndOpCodeF.Update(isCInstr, f, nil, 0)
	opCodeNO := cpu.cInstrAndOpCodeNO.Update(isCInstr, no, nil, 0)

	result, zr, ng := cpu.alu.Update(aluDVal, aluAVal, opCodeZx, opCodeNx, opCodeZy, opCodeNy, opCodeF, opCodeNO)
	cpu.prevAluOut = result

	isCInstrAndZr := cpu.cInstrAndZr.Update(isCInstr, zr, nil, 0)
	isCInstrAndNg := cpu.cInstrAndNg.Update(isCInstr, ng, nil, 0)
	isNotNg := cpu.notNg.Update(isCInstrAndNg, nil, 0)
	isPos := cpu.posXor.Update(isCInstrAndZr, isNotNg)

	jlt := cpu.jmpNegAnd.Update(j1, isCInstrAndNg, nil, 0)
	jeq := cpu.jmpZeroAnd.Update(j2, isCInstrAndZr, nil, 0)
	jgt := cpu.jmpPosAnd.Update(j3, isPos, nil, 0)
	jge := cpu.posOrZero.Update(jgt, jeq, nil, 0)
	jne := cpu.negOrPos.Update(jlt, jgt, nil, 0)
	jle := cpu.negOrZero.Update(jlt, jeq, nil, 0)
	jmp := cpu.zeroOrPosOrNeg.Update(jge, jlt, nil, 0)
	pcJump := cpu.jumpOr.Update(jmp, jne, jle, jeq, jgt, jge, jlt, false)

	pc = cpu.pc.Update(aRegVal, pcJump, true, reset)
	loadMem = cpu.cInstrAndD3.Update(isCInstr, d3, nil, 0)
	memAddr = aRegVal
	memOut = result

	return
}

func (cpu *CPU) Tick() {
	cpu.aRegister.Tick(nil)
	cpu.dRegister.Tick(nil)
	cpu.pc.Tick()
}
