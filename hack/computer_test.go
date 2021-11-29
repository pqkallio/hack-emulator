package hack

import (
	"testing"

	"github.com/pqkallio/hack-emulator/hack/components/combinational/word"
	seq "github.com/pqkallio/hack-emulator/hack/components/sequential/word"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	// RAM[0] = 2 + 3
	instructions := []uint16{
		0b0000000000000010, // @2
		0b1110110000010000, // D=A
		0b0000000000000011, // @3
		0b1110000010010000, // D=D+A
		0b0000000000000000, // @0
		0b1110001100001000, // M=D
	}

	rom := word.NewROM32KFlat()
	rom.Flash(instructions)

	ram := seq.NewRam16kFlat()
	scr := seq.NewScreenMem()
	kbd := word.NewKeyboardMem()

	computer := NewComputer(ram, scr, kbd, rom)

	for i := 0; i < 10; i++ {
		computer.Next(false)
	}

	actual := ram.Read(0)
	expected := uint16(5)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	// RAM[2] = max(RAM[0], RAM[1])
	instructions := []uint16{
		0b0000000000000000, // @0
		0b1111110000010000, // D=M
		0b0000000000000001, // @1
		0b1111010011010000, // D=D-M
		0b0000000000001010, // @10
		0b1110001100000001, // D;JGT
		0b0000000000000001, // @1
		0b1111110000010000, // D=M
		0b0000000000001100, // @12
		0b1110101010000111, // 0;JMP
		0b0000000000000000, // @0
		0b1111110000010000, // D=M
		0b0000000000000010, // @2
		0b1110001100001000, // M=D
		0b0000000000001110, // @14
		0b1110101010000111, // 0;JMP
	}

	rom := word.NewROM32KFlat()
	rom.Flash(instructions)

	ram := seq.NewRam16kFlat()
	scr := seq.NewScreenMem()
	kbd := word.NewKeyboardMem()

	computer := NewComputer(ram, scr, kbd, rom)

	ram.Write(0, uint16(3))
	ram.Write(1, uint16(5))

	for i := 0; i < 30; i++ {
		computer.Next(false)
	}

	actual := ram.Read(2)
	expected := uint16(5)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}

	computer.Next(true) // reset the computer

	ram.Write(0, uint16(23456))
	ram.Write(1, uint16(12345))

	for i := 0; i < 30; i++ {
		computer.Next(false)
	}

	actual = ram.Read(2)
	expected = uint16(23456)

	if actual != expected {
		t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
	}
}

func TestRect(t *testing.T) {
	t.Parallel()

	// Draw a 16xRAM[0] rectangle
	instructions := []uint16{
		0b0000000000000000, // 0  @0
		0b1111110000010000, // 1  D=M
		0b0000000000010111, // 2  @23
		0b1110001100000110, // 3  D;JLE
		0b0000000000010000, // 4  @16
		0b1110001100001000, // 5  M=D
		0b0100000000000000, // 6  @16384
		0b1110110000010000, // 7  D=A
		0b0000000000010001, // 8  @17
		0b1110001100001000, // 9  M=D
		0b0000000000010001, // 10 @17
		0b1111110000100000, // 11 A=M
		0b1110111010001000, // 12 M=-1
		0b0000000000010001, // 13 @17
		0b1111110000010000, // 14 D=M
		0b0000000000100000, // 15 @32
		0b1110000010010000, // 16 D=D+A
		0b0000000000010001, // 17 @17
		0b1110001100001000, // 18 M=D
		0b0000000000010000, // 19 @16
		0b1111110010011000, // 20 DM=M-1
		0b0000000000001010, // 21 @10
		0b1110001100000001, // 22 D;JGT
		0b0000000000010111, // 23 @23
		0b1110101010000111, // 24 @0;JMP
	}

	rom := word.NewROM32KFlat()
	rom.Flash(instructions)

	ram := seq.NewRam16kFlat()
	scr := seq.NewScreenMem()
	kbd := word.NewKeyboardMem()

	computer := NewComputer(ram, scr, kbd, rom)

	ram.Write(0, 4)

	for i := 0; i < 90; i++ {
		computer.Next(false)
	}

	scr.Tick()

	for i := uint16(0); i < 32*4; i += 32 {
		actual := scr.Read(i)
		expected := uint16(0xffff)

		if actual != expected {
			t.Errorf("%d: expected:\n%+v\ngot:\n%+v", i, expected, actual)
		}
	}
}
