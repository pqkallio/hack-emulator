package word

import (
	"testing"

	"github.com/pqkallio/hack-emulator/util"
)

func TestRAM512(t *testing.T) {
	t.Parallel()

	val := uint16(0x1234)

	for i := 0; i < 8; i++ {
		ram512 := NewRAM512()

		origAddr0 := util.RandomBool()
		origAddr1 := util.RandomBool()
		origAddr2 := util.RandomBool()
		origAddr3 := util.RandomBool()
		origAddr4 := util.RandomBool()
		origAddr5 := util.RandomBool()
		origAddr6 := util.RandomBool()
		origAddr7 := util.RandomBool()
		origAddr8 := util.RandomBool()

		actual := ram512.Update(
			val, true,
			origAddr0, origAddr1, origAddr2,
			origAddr3, origAddr4, origAddr5,
			origAddr6, origAddr7, origAddr8,
			nil, 0,
		)

		if actual != 0 {
			t.Errorf("expected:\n%+v\ngot:\n%+v", 0, actual)
		}

		ram512.Tick(nil)

		actual = ram512.Update(val, false,
			origAddr0, origAddr1, origAddr2,
			origAddr3, origAddr4, origAddr5,
			origAddr6, origAddr7, origAddr8,
			nil, 0,
		)

		if actual != val {
			t.Errorf("expected:\n%+v\ngot:\n%+v", val, actual)
		}

		ram512.Tick(nil)

		for j := 0; j < 8; j++ {
			addr0 := util.RandomBool()
			addr1 := util.RandomBool()
			addr2 := util.RandomBool()
			addr3 := util.RandomBool()
			addr4 := util.RandomBool()
			addr5 := util.RandomBool()
			addr6 := util.RandomBool()
			addr7 := util.RandomBool()
			addr8 := util.RandomBool()

			actual = ram512.Update(
				val, false,
				addr0, addr1, addr2,
				addr3, addr4, addr5,
				addr6, addr7, addr8,
				nil, 0,
			)
			expected := uint16(0)

			if addr0 == origAddr0 &&
				addr1 == origAddr1 &&
				addr2 == origAddr2 &&
				addr3 == origAddr3 &&
				addr4 == origAddr4 &&
				addr5 == origAddr5 &&
				addr6 == origAddr6 &&
				addr7 == origAddr7 &&
				addr8 == origAddr8 {

				expected = val
			}

			if actual != expected {
				t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
			}

			ram512.Tick(nil)
		}
	}
}
