package components

import "math"

type Addr16K uint16

func (addr Addr16K) ToAddressLines() (lines [14]bool) {
	for i := 0; i < 14; i++ {
		lines[i] = addr&Addr16K(math.Pow(2, float64(i))) != 0
	}

	return
}

type Addr32K uint16

func (addr Addr32K) ToAddressLines() (lines [15]bool) {
	for i := 0; i < 15; i++ {
		lines[i] = addr&Addr32K(math.Pow(2, float64(i))) != 0
	}

	return
}
