package word

import (
	"reflect"
	"testing"
)

func TestMux4Way16(t *testing.T) {
	t.Parallel()

	type args struct {
		sel0 bool
		sel1 bool
	}

	tests := []struct {
		name     string
		args     args
		expected uint16
	}{
		{
			"a: sel 0x00 => 0x0001",
			args{
				false,
				false,
			},
			0x0001,
		},
		{
			"b: sel 0x01 => 0x0002",
			args{
				true,
				false,
			},
			0x0002,
		},
		{
			"c: sel 0x02 => 0x0003",
			args{
				false,
				true,
			},
			0x0003,
		},
		{
			"d: sel 0x03 => 0x0004",
			args{
				true,
				true,
			},
			0x0004,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			or := NewMux4Way16()

			result := or.Update(
				0x0001,
				0x0002,
				0x0003,
				0x0004,
				tt.args.sel0,
				tt.args.sel1,
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
