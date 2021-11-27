package word

import "testing"

func TestLoadToMemory(t *testing.T) {
	t.Parallel()

	test := args{
		name: "load 12345 to mem addr 23456",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001}, // @12345
			{instruction: 0b11101100_00010000}, // D=A
			{instruction: 0b01011011_10100000}, // @23456
			{instruction: 0b11100011_00001000}, // M=D
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{12345, 12345, 2, false},
			{12345, 12345, 3, false},
			{12345, 23456, 4, true},
		},
	}

	testCPU(t, test)
}

func TestLoadFromMemory(t *testing.T) {
	t.Parallel()

	test := args{
		name: "load M-1 to mem addr 0",
		args: []cpuUpdateArgs{
			{instruction: 0b11111100_10010000, in: 43}, // D=M-1
			{instruction: 0b00000000_00000000},         // @0
			{instruction: 0b11100011_00001000},         // M=D
		},
		expected: []cpuUpdateResult{
			{42, 0, 1, false},
			{0, 0, 2, false},
			{42, 0, 3, true},
		},
	}

	testCPU(t, test)
}

func TestJLT(t *testing.T) {
	t.Parallel()

	test := args{
		name: "JLT",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001}, // @12345
			{instruction: 0b11101111_11010000}, // D=1
			{instruction: 0b11100011_00000100}, // D;JLT
			{instruction: 0b11100011_10010100}, // D=D-1;JLT
			{instruction: 0b11100011_10010100}, // D=D-1;JLT
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{1, 12345, 2, false},
			{1, 12345, 3, false},
			{0, 12345, 4, false},
			{65535, 12345, 12345, false},
		},
	}

	testCPU(t, test)
}

func TestJLE(t *testing.T) {
	t.Parallel()

	test := args{
		name: "JLE",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001}, // @12345
			{instruction: 0b11101111_11010000}, // D=1
			{instruction: 0b11100011_00000110}, // D;JLE
			{instruction: 0b11100011_10010110}, // D=D-1;JLE
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{1, 12345, 2, false},
			{1, 12345, 3, false},
			{0, 12345, 12345, false},
		},
	}

	testCPU(t, test)
}

func TestJEQ(t *testing.T) {
	t.Parallel()

	test := args{
		name: "JEQ",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001}, // @12345
			{instruction: 0b11101111_11010000}, // D=1
			{instruction: 0b11100011_00000010}, // D;JEQ
			{instruction: 0b11100011_10010010}, // D=D-1;JEQ
			{instruction: 0b11100011_10010010}, // D=D-1;JEQ
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{1, 12345, 2, false},
			{1, 12345, 3, false},
			{0, 12345, 12345, false},
			{65535, 12345, 12346, false},
		},
	}

	testCPU(t, test)
}

func TestJGE(t *testing.T) {
	t.Parallel()

	test := args{
		name: "JGE",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001}, // @12345
			{instruction: 0b11101111_11010000}, // D=1
			{instruction: 0b11100011_00000011}, // D;JGE
			{instruction: 0b11100011_10010011}, // D=D-1;JGE
			{instruction: 0b11100011_10010011}, // D=D-1;JGE
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{1, 12345, 2, false},
			{1, 12345, 12345, false},
			{0, 12345, 12345, false},
			{65535, 12345, 12346, false},
		},
	}

	testCPU(t, test)
}

func TestJGT(t *testing.T) {
	t.Parallel()

	test := args{
		name: "JGT",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001}, // @12345
			{instruction: 0b11101111_11010000}, // D=1
			{instruction: 0b11100011_00000001}, // D;JGT
			{instruction: 0b11100011_10010001}, // D=D-1;JGT
			{instruction: 0b11100011_10010001}, // D=D-1;JGT
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{1, 12345, 2, false},
			{1, 12345, 12345, false},
			{0, 12345, 12346, false},
			{65535, 12345, 12347, false},
		},
	}

	testCPU(t, test)
}

func TestJNE(t *testing.T) {
	t.Parallel()

	test := args{
		name: "JNE",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001}, // @12345
			{instruction: 0b11101111_11010000}, // D=1
			{instruction: 0b11100011_00000101}, // D;JNE
			{instruction: 0b11100011_10010101}, // D=D-1;JNE
			{instruction: 0b11100011_10010101}, // D=D-1;JNE
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{1, 12345, 2, false},
			{1, 12345, 12345, false},
			{0, 12345, 12346, false},
			{65535, 12345, 12345, false},
		},
	}

	testCPU(t, test)
}

func TestJMP(t *testing.T) {
	t.Parallel()

	test := args{
		name: "JMP",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001}, // @12345
			{instruction: 0b11101111_11010000}, // D=1
			{instruction: 0b11100011_00000111}, // D;JMP
			{instruction: 0b11100011_10010111}, // D=D-1;JMP
			{instruction: 0b11100011_10010111}, // D=D-1;JMP
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{1, 12345, 2, false},
			{1, 12345, 12345, false},
			{0, 12345, 12345, false},
			{65535, 12345, 12345, false},
		},
	}

	testCPU(t, test)
}

func TestReset(t *testing.T) {
	t.Parallel()

	test := args{
		name: "reset",
		args: []cpuUpdateArgs{
			{instruction: 0b00110000_00111001},              // @12345
			{instruction: 0b11101111_11010000, reset: true}, // D=1
		},
		expected: []cpuUpdateResult{
			{0, 0, 1, false},
			{1, 12345, 0, false},
		},
	}

	testCPU(t, test)
}

func testCPU(t *testing.T, test args) {
	t.Helper()

	cpu := NewCPU()

	for i, arg := range test.args {
		memOut, memAddr, loadMem := cpu.Execute(arg.instruction, arg.in, arg.reset)

		if memOut != test.expected[i].memOut {
			t.Errorf("%d: memOut = %v, want %v", i, memOut, test.expected[i].memOut)
		}

		if memAddr != test.expected[i].memAddr {
			t.Errorf("%d: memAddr = %v, want %v", i, memAddr, test.expected[i].memAddr)
		}

		if loadMem != test.expected[i].loadMem {
			t.Errorf("%d: loadMem = %v, want %v", i, loadMem, test.expected[i].loadMem)
		}

		cpu.Tick()

		pc := cpu.Fetch()

		if pc != test.expected[i].pc {
			t.Errorf("%d: pc = %v, want %v", i, pc, test.expected[i].pc)
		}
	}
}

type cpuUpdateArgs struct {
	instruction, in uint16
	reset           bool
}

type cpuUpdateResult struct {
	memOut, memAddr, pc uint16
	loadMem             bool
}

type args struct {
	name     string
	args     []cpuUpdateArgs
	expected []cpuUpdateResult
}
