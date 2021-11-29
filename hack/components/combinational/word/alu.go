package word

import (
	"github.com/pqkallio/hack-emulator/hack/components/combinational/bit"
	"github.com/pqkallio/hack-emulator/util"
)

// ALU, or the arithmetic-logical unit, performs calculations
// based on the two 16-bit inputs x and y, and the opcode composed
// of 6 separate channels:
//   zx => x = 0
//   nx => x = !x
//   zy => y = 0
//   ny => y = !y
//   f  => out = two's compliment x + y, else out = x & y
//   no => out = !out
type ALU struct {
	// x preprocessing gates
	zxMux *Mux16
	nxNot *Not16
	nxMux *Mux16

	// y preprocessing gates
	zyMux *Mux16
	nyNot *Not16
	nyMux *Mux16

	// function gates
	fAdd *Add16
	fAnd *And16
	fMux *Mux16

	// postprocess gates
	noNot *Not16
	noMux *Mux16

	// zero flag gates
	zrOr8Way1 *bit.Or8Way
	zrOr8Way2 *bit.Or8Way
	zrOr      *bit.Or
	zrNot     *bit.Not
}

func NewALU() *ALU {
	return &ALU{
		NewMux16(), NewNot16(), NewMux16(),
		NewMux16(), NewNot16(), NewMux16(),
		NewAdd16(), NewAnd16(), NewMux16(),
		NewNot16(), NewMux16(),
		bit.NewOr8Way(), bit.NewOr8Way(), bit.NewOr(), bit.NewNot(),
	}
}

// Update updates the ALU's channels and returns the result of
// the computation as three different components in the following
// order:
//   1. result of the computation, a SixteenChannel,
//   2. zero flag, a SingleChannel, true if the result of the computation
//      equals to 0,
//   3. negative flag, a SingleChannel, true if the result of the computation
//      is less than 0.
//
// Inputs x and y can be updated with an UpdateOpts to TargetX and TargetY,
// respectively. The UpdateOpts value must be a SixteenChan.
//
// The following table represents the valid opcodes and their output:
//
// | zx  | nx  | zy  | ny  |  f  | no  | out |
// |-----|-----|-----|-----|-----|-----|-----|
// |  1  |  0  |  1  |  0  |  1  |  0  |  0  |
// |  1  |  1  |  1  |  1  |  1  |  1  |  1  |
// |  1  |  1  |  1  |  0  |  1  |  0  | -1  |
// |  0  |  0  |  1  |  1  |  0  |  0  |  x  |
// |  1  |  1  |  0  |  0  |  0  |  0  |  y  |
// |  0  |  0  |  1  |  1  |  0  |  1  | !x  |
// |  1  |  1  |  0  |  0  |  0  |  1  | !y  |
// |  0  |  0  |  1  |  1  |  1  |  1  | -x  |
// |  1  |  1  |  0  |  0  |  1  |  1  | -y  |
// |  0  |  1  |  1  |  1  |  1  |  1  | x+1 |
// |  1  |  1  |  0  |  1  |  1  |  1  | y+1 |
// |  0  |  0  |  1  |  1  |  1  |  0  | x-1 |
// |  1  |  1  |  0  |  0  |  1  |  0  | y-1 |
// |  0  |  0  |  0  |  0  |  1  |  0  | x+y |
// |  0  |  1  |  0  |  0  |  1  |  1  | x-y |
// |  0  |  0  |  0  |  1  |  1  |  1  | y-x |
// |  0  |  0  |  0  |  0  |  0  |  0  | x&y |
// |  0  |  1  |  0  |  1  |  0  |  1  | x|y |
//
// The column name to UpdateOpts target is the following:
//   zx = TargetZeroX
//   nx = TargetNegX
//   zy = TargetZeroY
//   ny = TargetNegY
//   f  = TargetFunc
//   no = TargetNegOut
func (alu *ALU) Update(
	x, y uint16,
	zx, nx, zy, ny, f, no bool,
) (result uint16, zr bool, ng bool) {
	// preprocess x
	xZero := alu.zxMux.Update(x, 0, zx, nil, 0)
	xNeg := alu.nxNot.Update(xZero)
	xPreprocessed := alu.nxMux.Update(xZero, xNeg, nx, nil, 0)

	// preprocess y
	yZero := alu.zyMux.Update(y, 0, zy, nil, 0)
	yNeg := alu.nyNot.Update(yZero)
	yPreprocessed := alu.nyMux.Update(yZero, yNeg, ny, nil, 0)

	// function(x, y)
	xyAdd := alu.fAdd.Update(xPreprocessed, yPreprocessed)
	xyAnd := alu.fAnd.Update(xPreprocessed, yPreprocessed)
	xyF := alu.fMux.Update(xyAnd, xyAdd, f, nil, 0)

	// postprocess xyF
	negXy := alu.noNot.Update(xyF)
	result = alu.noMux.Update(xyF, negXy, no, nil, 0)

	// set status flags
	ng = util.GetBoolFromUint16(result, 15)

	loByteOr := alu.zrOr8Way1.Update(
		util.GetBoolFromUint16(result, 0),
		util.GetBoolFromUint16(result, 1),
		util.GetBoolFromUint16(result, 2),
		util.GetBoolFromUint16(result, 3),
		util.GetBoolFromUint16(result, 4),
		util.GetBoolFromUint16(result, 5),
		util.GetBoolFromUint16(result, 6),
		util.GetBoolFromUint16(result, 7),
	)
	hiByteOr := alu.zrOr8Way2.Update(
		util.GetBoolFromUint16(result, 8),
		util.GetBoolFromUint16(result, 9),
		util.GetBoolFromUint16(result, 10),
		util.GetBoolFromUint16(result, 11),
		util.GetBoolFromUint16(result, 12),
		util.GetBoolFromUint16(result, 13),
		util.GetBoolFromUint16(result, 14),
		util.GetBoolFromUint16(result, 15),
	)
	zrFlag := alu.zrOr.Update(loByteOr, hiByteOr, nil, 0)
	zr = alu.zrNot.Update(zrFlag, nil, 0)

	return result, zr, ng
}
