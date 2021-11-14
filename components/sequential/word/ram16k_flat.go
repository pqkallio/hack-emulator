package word

import "github.com/pqkallio/hack-emulator/components"

type Ram16kFlat struct {
	mem [16384]*Register
}

func NewRam16kFlat() *Ram16kFlat {
	mem := [16384]*Register{}

	for i := 0; i < 16384; i++ {
		mem[i] = NewRegister()
	}

	return &Ram16kFlat{mem}
}

func (r *Ram16kFlat) Update(
	in uint16,
	load bool,
	addr0, addr1, addr2 bool,
	addr3, addr4, addr5 bool,
	addr6, addr7, addr8 bool,
	addr9, addr10, addr11 bool,
	addr12, addr13 bool,
	ch chan components.OrderedVal16, idx int,
) uint16 {
	addr := uint16(0)

	if addr0 {
		addr |= 0x2000
	}
	if addr1 {
		addr |= 0x1000
	}
	if addr2 {
		addr |= 0x0800
	}
	if addr3 {
		addr |= 0x0400
	}
	if addr4 {
		addr |= 0x0200
	}
	if addr5 {
		addr |= 0x0100
	}
	if addr6 {
		addr |= 0x0080
	}
	if addr7 {
		addr |= 0x0040
	}
	if addr8 {
		addr |= 0x0020
	}
	if addr9 {
		addr |= 0x0010
	}
	if addr10 {
		addr |= 0x0008
	}
	if addr11 {
		addr |= 0x0004
	}
	if addr12 {
		addr |= 0x0002
	}
	if addr13 {
		addr |= 0x0001
	}

	return r.mem[addr].Update(in, load, nil, 0)
}
