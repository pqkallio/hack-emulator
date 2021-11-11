package word

import (
	"reflect"
	"testing"
)

func TestMux16(t *testing.T) {
	t.Parallel()

	type args struct {
		a   uint16
		b   uint16
		sel bool
	}

	tests := []struct {
		name     string
		args     args
		expected uint16
	}{
		{
			"a: 0x0001, b: 0x0010, sel: 0 = 0x0001",
			args{
				0x0001,
				0x0010,
				false,
			},
			0x0001,
		},
		{
			"a: 0x0001, b: 0x0010, sel: 1 = 0x0010",
			args{
				0x0001,
				0x0010,
				true,
			},
			0x0010,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux16 := NewMux16()

			result := mux16.Update(
				tt.args.a,
				tt.args.b,
				tt.args.sel,
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
