package word

import (
	"reflect"
	"testing"
)

type aluTestArgs struct {
	x, y                  uint16 // 16-bit input
	zx, nx, zy, ny, f, no bool   // 1-bit flags
}

type aluTest struct {
	name                 string
	args                 aluTestArgs
	expectedResult       uint16
	expectedZeroFlag     bool
	expectedNegativeFlag bool
}

type aluTests []aluTest

func opCodeWithXY(x, y uint16, opCode uint8) aluTestArgs {
	return aluTestArgs{
		x,
		y,
		opCode&32 != 0,
		opCode&16 != 0,
		opCode&8 != 0,
		opCode&4 != 0,
		opCode&2 != 0,
		opCode&1 != 0,
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
			0x0,
			true,
			false,
		},
		{
			"non-zero y",
			opCodeWithXY(0x0, 0xffff, opCode),
			0x0,
			true,
			false,
		},
		{
			"non-zero x & y",
			opCodeWithXY(0xf0f0, 0x0f0f, opCode),
			0x0,
			true,
			false,
		},
		{
			"zero x & y",
			opCodeWithXY(0x0000, 0x0000, opCode),
			0x0,
			true,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0x1,
			false,
			false,
		},
		{
			"non-one y",
			opCodeWithXY(0x0001, 0xffff, opCode),
			0x1,
			false,
			false,
		},
		{
			"non-one x & y",
			opCodeWithXY(0xf0f0, 0x0f0f, opCode),
			0x1,
			false,
			false,
		},
		{
			"one x & y",
			opCodeWithXY(0x0001, 0x0001, opCode),
			0x1,
			false,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0xffff,
			false,
			true,
		},
		{
			"non-one y",
			opCodeWithXY(0x0001, 0xffff, opCode),
			0xffff,
			false,
			true,
		},
		{
			"non-one x & y",
			opCodeWithXY(0xf0f0, 0x0f0f, opCode),
			0xffff,
			false,
			true,
		},
		{
			"one x & y",
			opCodeWithXY(0x0001, 0x0001, opCode),
			0xffff,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0x0000,
			true,
			false,
		},
		{
			"x = 1",
			opCodeWithXY(0x0001, 0x0000, opCode),
			0x0001,
			false,
			false,
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			0xffff,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0x0000,
			true,
			false,
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			0x0001,
			false,
			false,
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			0xffff,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0xffff,
			false,
			true,
		},
		{
			"x = 1",
			opCodeWithXY(0b00000000_00000001, 0x0000, opCode),
			0b11111111_11111110,
			false,
			true,
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			0x0000,
			true,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0xffff,
			false,
			true,
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			0xfffe,
			false,
			true,
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			0x0000,
			true,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0x0000,
			true,
			false,
		},
		{
			"x = 1",
			opCodeWithXY(0x0001, 0x0000, opCode),
			0xffff,
			false,
			true,
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			0x0001,
			false,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0x0000,
			true,
			false,
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			0xffff,
			false,
			true,
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			0x0001,
			false,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0x0001,
			false,
			false,
		},
		{
			"x = 1",
			opCodeWithXY(0x0001, 0x0000, opCode),
			0x0002,
			false,
			false,
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			0x0000,
			true,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0x0001,
			false,
			false,
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			0x0002,
			false,
			false,
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			0x0000,
			true,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0xffff,
			false,
			true,
		},
		{
			"x = 1",
			opCodeWithXY(0x0001, 0x0000, opCode),
			0x0000,
			true,
			false,
		},
		{
			"x = -1",
			opCodeWithXY(0xffff, 0x0000, opCode),
			0xfffe,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			0xffff,
			false,
			true,
		},
		{
			"y = 1",
			opCodeWithXY(0x0000, 0x0001, opCode),
			0x0000,
			true,
			false,
		},
		{
			"y = -1",
			opCodeWithXY(0x0000, 0xffff, opCode),
			0xfffe,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
			7,
			false,
			false,
		},
		{
			"2 + 5 = 7",
			opCodeWithXY(2, 5, opCode),
			7,
			false,
			false,
		},
		{
			"-2 + 4 = 2",
			opCodeWithXY(0xfffe, 0x0004, opCode),
			0x0002,
			false,
			false,
		},
		{
			"2 + -4 = -2",
			opCodeWithXY(0x0002, 0xfffc, opCode),
			0xfffe,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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

// TestOpCode010011 tests that the ALU always returns
// x-y as a result.
func TestOpCode010011(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b010011)

	tests := aluTests{
		{
			"5 - 2 = 3",
			opCodeWithXY(5, 2, opCode),
			3,
			false,
			false,
		},
		{
			"2 - 5 = -3",
			opCodeWithXY(2, 5, opCode),
			0xfffd,
			false,
			true,
		},
		{
			"-2 - 4 = -6",
			opCodeWithXY(0xfffe, 0x0004, opCode),
			0xfffa,
			false,
			true,
		},
		{
			"2 - -4 = 6",
			opCodeWithXY(0x0002, 0xfffc, opCode),
			0x0006,
			false,
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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

// TestOpCode000111 tests that the ALU always returns
// y-x as a result.
func TestOpCode000111(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b000111)

	tests := aluTests{
		{
			"2 - 5 = -3",
			opCodeWithXY(5, 2, opCode),
			0xfffd,
			false,
			true,
		},
		{
			"5 - 2 = 3",
			opCodeWithXY(2, 5, opCode),
			3,
			false,
			false,
		},
		{
			"4 - -2 = 6",
			opCodeWithXY(0xfffe, 0x0004, opCode),
			6,
			false,
			false,
		},
		{
			"-4 - 2 = -6",
			opCodeWithXY(0x0002, 0xfffc, opCode),
			0xfffa,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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

// TestOpCode000000 tests that the ALU always returns
// x&y as a result.
func TestOpCode000000(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b000000)

	tests := aluTests{
		{
			"0x0000 & 0xffff = 0x0000",
			opCodeWithXY(0x0000, 0xffff, opCode),
			0x0000,
			true,
			false,
		},
		{
			"0xffff & 0x0000 = 0x0000",
			opCodeWithXY(0xffff, 0x0000, opCode),
			0x0000,
			true,
			false,
		},
		{
			"0xffaf & 0xe0e0 = 0xe0a0",
			opCodeWithXY(0xffaf, 0xe0e0, opCode),
			0xe0a0,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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

// TestOpCode010101 tests that the ALU always returns
// x&y as a result.
func TestOpCode010101(t *testing.T) {
	t.Parallel()

	opCode := uint8(0b010101)

	tests := aluTests{
		{
			"0x0000 | 0xffff = 0xffff",
			opCodeWithXY(0x0000, 0xffff, opCode),
			0xffff,
			false,
			true,
		},
		{
			"0xffff | 0x0000 = 0xffff",
			opCodeWithXY(0xffff, 0x0000, opCode),
			0xffff,
			false,
			true,
		},
		{
			"0xffaf & 0xe0e0 = 0xffef",
			opCodeWithXY(0xffaf, 0xe0e0, opCode),
			0xffef,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			alu := NewALU()

			result, zr, ng := alu.Update(
				tt.args.x,
				tt.args.y,
				tt.args.zx,
				tt.args.nx,
				tt.args.zy,
				tt.args.ny,
				tt.args.f,
				tt.args.no,
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
