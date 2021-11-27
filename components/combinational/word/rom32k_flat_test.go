package word

import (
	"math/rand"
	"testing"
)

func TestROM32KFlat(t *testing.T) {
	t.Parallel()

	data := make([]uint16, 32768)

	for i := 0; i < 32768; i++ {
		data[i] = uint16(rand.Intn(65536))
	}

	rom := NewROM32KFlat()
	rom.Flash(data)

	for i := 0; i < 32768; i++ {
		actual := rom.Get(uint16(i))
		expected := data[i]

		if actual != expected {
			t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
		}
	}
}
