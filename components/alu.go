package components

type ALU struct {
	x, y                  Val // inputs
	zx, nx, zy, ny, f, no Val // flags

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

	// zero flag gatest
	zrOr8Way1 *Or8Way
	zrOr8Way2 *Or8Way
	zrOr      *Or
	zrNot     *Not
}

func NewALU() *ALU {
	return &ALU{
		&InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{}, &InvalidVal{},
		&InvalidVal{}, &InvalidVal{}, &InvalidVal{},
		NewMux16(), NewNot16(), NewMux16(),
		NewMux16(), NewNot16(), NewMux16(),
		NewAdd16(), NewAnd16(), NewMux16(),
		NewNot16(), NewMux16(),
		NewOr8Way(), NewOr8Way(), NewOr(), NewNot(),
	}
}

func (alu *ALU) Update(opts ...UpdateOpts) (Val, Val, Val) {
	for _, opt := range opts {
		switch opt.target {
		case TargetX:
			alu.x = opt.val
		case TargetY:
			alu.y = opt.val
		case TargetZeroX:
			alu.zx = opt.val
		case TargetNegX:
			alu.nx = opt.val
		case TargetZeroY:
			alu.zy = opt.val
		case TargetNegY:
			alu.ny = opt.val
		case TargetFunc:
			alu.f = opt.val
		case TargetNegOut:
			alu.no = opt.val
		}
	}

	// preprocess x
	xZero := alu.zxMux.Update(
		UpdateOpts{TargetA, alu.x},
		UpdateOpts{TargetB, &SixteenChan{0}},
		UpdateOpts{TargetSel0, alu.zx},
	)

	xNeg := alu.nxNot.Update(
		UpdateOpts{TargetIn, xZero},
	)
	xPreprocessed := alu.nxMux.Update(
		UpdateOpts{TargetA, xZero},
		UpdateOpts{TargetB, xNeg},
		UpdateOpts{TargetSel0, alu.nx},
	)

	// preprocess y
	yZero := alu.zyMux.Update(
		UpdateOpts{TargetA, alu.y},
		UpdateOpts{TargetB, &SixteenChan{0}},
		UpdateOpts{TargetSel0, alu.zy},
	)

	yNeg := alu.nyNot.Update(
		UpdateOpts{TargetIn, yZero},
	)
	yPreprocessed := alu.nyMux.Update(
		UpdateOpts{TargetA, yZero},
		UpdateOpts{TargetB, yNeg},
		UpdateOpts{TargetSel0, alu.ny},
	)

	// function(x, y)
	xyAdd := alu.fAdd.Update(
		UpdateOpts{TargetA, xPreprocessed},
		UpdateOpts{TargetB, yPreprocessed},
	)
	xyAnd := alu.fAnd.Update(
		UpdateOpts{TargetA, xPreprocessed},
		UpdateOpts{TargetB, yPreprocessed},
	)
	xyF := alu.fMux.Update(
		UpdateOpts{TargetA, xyAnd},
		UpdateOpts{TargetB, xyAdd},
		UpdateOpts{TargetSel0, alu.f},
	)

	// postprocess xyF
	negXy := alu.noNot.Update(
		UpdateOpts{TargetIn, xyF},
	)
	result := alu.noMux.Update(
		UpdateOpts{TargetA, xyF},
		UpdateOpts{TargetB, negXy},
		UpdateOpts{TargetSel0, alu.no},
	)

	// set status flags
	ng := SingleChan{result.GetBoolFromUint16(15)}

	loByteOr := alu.zrOr8Way1.Update(
		UpdateOpts{TargetA, &SingleChan{result.GetBoolFromUint16(0)}},
		UpdateOpts{TargetB, &SingleChan{result.GetBoolFromUint16(1)}},
		UpdateOpts{TargetC, &SingleChan{result.GetBoolFromUint16(2)}},
		UpdateOpts{TargetD, &SingleChan{result.GetBoolFromUint16(3)}},
		UpdateOpts{TargetE, &SingleChan{result.GetBoolFromUint16(4)}},
		UpdateOpts{TargetF, &SingleChan{result.GetBoolFromUint16(5)}},
		UpdateOpts{TargetG, &SingleChan{result.GetBoolFromUint16(6)}},
		UpdateOpts{TargetH, &SingleChan{result.GetBoolFromUint16(7)}},
	)
	hiByteOr := alu.zrOr8Way2.Update(
		UpdateOpts{TargetA, &SingleChan{result.GetBoolFromUint16(8)}},
		UpdateOpts{TargetB, &SingleChan{result.GetBoolFromUint16(9)}},
		UpdateOpts{TargetC, &SingleChan{result.GetBoolFromUint16(10)}},
		UpdateOpts{TargetD, &SingleChan{result.GetBoolFromUint16(11)}},
		UpdateOpts{TargetE, &SingleChan{result.GetBoolFromUint16(12)}},
		UpdateOpts{TargetF, &SingleChan{result.GetBoolFromUint16(13)}},
		UpdateOpts{TargetG, &SingleChan{result.GetBoolFromUint16(14)}},
		UpdateOpts{TargetH, &SingleChan{result.GetBoolFromUint16(15)}},
	)
	zrFlag := alu.zrOr.Update(
		UpdateOpts{TargetA, loByteOr},
		UpdateOpts{TargetB, hiByteOr},
	)
	zr := alu.zrNot.Update(
		UpdateOpts{TargetIn, zrFlag},
	)

	return result, zr, &ng
}

const (
	TargetX Target = iota + 100
	TargetY
	TargetZeroX
	TargetNegX
	TargetZeroY
	TargetNegY
	TargetFunc
	TargetNegOut
)
