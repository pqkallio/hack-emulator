package components

import (
	"reflect"
	"testing"
)

type aluTestArgs struct {
	x, y                  Val // 16-bit input
	zx, nx, zy, ny, f, no Val // 1-bit flags
}

type aluTest struct {
	name                 string
	args                 aluTestArgs
	expectedResult       Val
	expectedZeroFlag     Val
	expectedNegativeFlag Val
}

type aluTests []aluTest

func opCodeWithXY(x, y uint16, opCode uint8) aluTestArgs {
	return aluTestArgs{
		&SixteenChan{x},
		&SixteenChan{y},
		&SingleChan{opCode&32 != 0},
		&SingleChan{opCode&16 != 0},
		&SingleChan{opCode&8 != 0},
		&SingleChan{opCode&4 != 0},
		&SingleChan{opCode&2 != 0},
		&SingleChan{opCode&1 != 0},
	}
}

// TestOpCode101010 tests that the ALU always returns
// 0 as a result.
func TestOpCode101010(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b101010)

	tests := aluTests{
		{
			"non-zero x",
			opCodeWithXY(0xffff, 0x0, opCode),
			&SixteenChan{0x0},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"non-zero y",
			opCodeWithXY(0x0, 0xffff, opCode),
			&SixteenChan{0x0},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"non-zero x & y",
			opCodeWithXY(0xf0f0, 0x0f0f, opCode),
			&SixteenChan{0x0},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"zero x & y",
			opCodeWithXY(0x0000, 0x0000, opCode),
			&SixteenChan{0x0},
			&SingleChan{true},
			&SingleChan{false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode111111 tests that the ALU always returns
// 1 as a result.
func TestOpCode111111(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b111111)

	tests := aluTests{
		{
			"non-one x",
			opCodeWithXY(0xffff, 0x0001, opCode),
			&SixteenChan{0x1},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"non-one y",
			opCodeWithXY(0x0001, 0xffff, opCode),
			&SixteenChan{0x1},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"non-one x & y",
			opCodeWithXY(0xf0f0, 0x0f0f, opCode),
			&SixteenChan{0x1},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"one x & y",
			opCodeWithXY(0x0001, 0x0001, opCode),
			&SixteenChan{0x1},
			&SingleChan{false},
			&SingleChan{false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode111010 tests that the ALU always returns
// -1 as a result.
func TestOpCode111010(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b111010)

	tests := aluTests{
		{
			"non-one x",
			opCodeWithXY(0xffff, 0x0001, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"non-one y",
			opCodeWithXY(0x0001, 0xffff, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"non-one x & y",
			opCodeWithXY(0xf0f0, 0x0f0f, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"one x & y",
			opCodeWithXY(0x0001, 0x0001, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode001100 tests that the ALU always returns
// x as a result.
func TestOpCode001100(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b001100)

	tests := aluTests{
		{
			"x = 0",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"x = 1",
			opCodeWithXY(0x0001, 0x0000, opCode),
			&SixteenChan{0x0001},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode110000 tests that the ALU always returns
// y as a result.
func TestOpCode110000(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b110000)

	tests := aluTests{
		{
			"y = 0",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			&SixteenChan{0x0001},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode001101 tests that the ALU always returns
// !x as a result.
func TestOpCode001101(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b001101)

	tests := aluTests{
		{
			"x = 0",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"x = 1",
			opCodeWithXY(0b00000000_00000001, 0x0000, opCode),
			&SixteenChan{0b11111111_11111110},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode110001 tests that the ALU always returns
// !y as a result.
func TestOpCode110001(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b110001)

	tests := aluTests{
		{
			"y = 0",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			&SixteenChan{0xfffe},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode001111 tests that the ALU always returns
// -x as a result.
func TestOpCode001111(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b001111)

	tests := aluTests{
		{
			"x = 0",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"x = 1",
			opCodeWithXY(0x0001, 0x0000, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0x0001},
			&SingleChan{false},
			&SingleChan{false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode110011 tests that the ALU always returns
// -y as a result.
func TestOpCode110011(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b110011)

	tests := aluTests{
		{
			"y = 0",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0x0001},
			&SingleChan{false},
			&SingleChan{false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode011111 tests that the ALU always returns
// x+1 as a result.
func TestOpCode011111(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b011111)

	tests := aluTests{
		{
			"x = 0",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0x0001},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"x = 1",
			opCodeWithXY(0x0001, 0x0000, opCode),
			&SixteenChan{0x0002},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode110111 tests that the ALU always returns
// y+1 as a result.
func TestOpCode110111(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b110111)

	tests := aluTests{
		{
			"y = 0",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0x0001},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			&SixteenChan{0x0002},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode001110 tests that the ALU always returns
// x-1 as a result.
func TestOpCode001110(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b001110)

	tests := aluTests{
		{
			"x = 0",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"x = 1",
			opCodeWithXY(0x0001, 0x0000, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0xfffe},
			&SingleChan{false},
			&SingleChan{true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode110010 tests that the ALU always returns
// y-1 as a result.
func TestOpCode110010(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b110010)

	tests := aluTests{
		{
			"y = 0",
			opCodeWithXY(0xffff, 0x0000, opCode),
			&SixteenChan{0xffff},
			&SingleChan{false},
			&SingleChan{true},
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			&SixteenChan{0x0000},
			&SingleChan{true},
			&SingleChan{false},
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			&SixteenChan{0xfffe},
			&SingleChan{false},
			&SingleChan{true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}

// TestOpCode000010 tests that the ALU always returns
// x+y as a result.
func TestOpCode000010(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b000010)

	tests := aluTests{
		{
			"5 + 2 = 7",
			opCodeWithXY(5, 2, opCode),
			&SixteenChan{7},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"2 + 5 = 7",
			opCodeWithXY(2, 5, opCode),
			&SixteenChan{7},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"-2 + 4 = 2",
			opCodeWithXY(0xfffe, 0x0004, opCode),
			&SixteenChan{0x0002},
			&SingleChan{false},
			&SingleChan{false},
		},
		{
			"2 + -4 = -2",
			opCodeWithXY(0x0002, 0xfffc, opCode),
			&SixteenChan{0xfffe},
			&SingleChan{false},
			&SingleChan{true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				UpdateOpts{TargetX, tt.args.x},
				UpdateOpts{TargetY, tt.args.y},
				UpdateOpts{TargetZeroX, tt.args.zx},
				UpdateOpts{TargetNegX, tt.args.nx},
				UpdateOpts{TargetZeroY, tt.args.zy},
				UpdateOpts{TargetNegY, tt.args.ny},
				UpdateOpts{TargetFunc, tt.args.f},
				UpdateOpts{TargetNegOut, tt.args.no},
			)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("RESULT: expected:\n%+v\ngot:\n%+v", tt.expectedResult, result)
			}

			if !reflect.DeepEqual(zr, tt.expectedZeroFlag) {
				t.Errorf("ZERO FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedZeroFlag, zr)
			}

			if !reflect.DeepEqual(ng, tt.expectedNegativeFlag) {
				t.Errorf("NEGATIVE FLAG: expected:\n%+v\ngot:\n%+v", tt.expectedNegativeFlag, ng)
			}
		})
	}
}
