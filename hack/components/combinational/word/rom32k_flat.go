package word

type ROM32KFlat struct {
	mem [32768]uint16
}

func NewROM32KFlat() *ROM32KFlat {
	return &ROM32KFlat{}
}

func (rom *ROM32KFlat) Get(addr uint16) uint16 {
	return rom.mem[addr]
}

func (rom *ROM32KFlat) Flash(data []uint16) {
	for i, d := range data {
		if i > 32767 {
			panic("ROM32KFlat: data out of range")
		}
		rom.mem[i] = d
	}
}
