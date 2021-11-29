package word

import (
	"reflect"
	"testing"
)

func TestOr16(t *testing.T) {
	t.Parallel()

	type args struct {
		a uint16
		b uint16
	}

	tests := []struct {
		name     string
		args     args
		expected uint16
	}{
		{
			"0x0000 | 0x0000 = 0x0000",
			args{0x0000, 0x0000},
			0x0000,
		},
		{
			"0x0000 | 0xffff = 0xffff",
			args{0x0000, 0xffff},
			0xffff,
		},
		{
			"0xffff | 0x0000 = 0x0000",
			args{0xffff, 0x0000},
			0xffff,
		},
		{
			"0xffff | 0xffff = 0xffff",
			args{0xffff, 0xffff},
			0xffff,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			and := NewOr16()

			result := and.Update(tt.args.a, tt.args.b)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
