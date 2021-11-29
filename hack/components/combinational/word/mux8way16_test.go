package word

import (
	"reflect"
	"testing"
)

func TestMux8Way16(t *testing.T) {
	t.Parallel()

	type args struct {
		sel0 bool
		sel1 bool
		sel2 bool
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
				false,
			},
			0x0001,
		},
		{
			"b: sel 0x01 => 0x0002",
			args{
				true,
				false,
				false,
			},
			0x0002,
		},
		{
			"c: sel 0x02 => 0x0003",
			args{
				false,
				true,
				false,
			},
			0x0003,
		},
		{
			"d: sel 0x03 => 0x0004",
			args{
				true,
				true,
				false,
			},
			0x0004,
		},
		{
			"e: sel 0x04 => 0x0005",
			args{
				false,
				false,
				true,
			},
			0x0005,
		},
		{
			"f: sel 0x05 => 0x0006",
			args{
				true,
				false,
				true,
			},
			0x0006,
		},
		{
			"g: sel 0x06 => 0x0007",
			args{
				false,
				true,
				true,
			},
			0x0007,
		},
		{
			"h: sel 0x07 => 0x0008",
			args{
				true,
				true,
				true,
			},
			0x0008,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux8Way := NewMux8Way16()

			result := mux8Way.Update(
				0x0001,
				0x0002,
				0x0003,
				0x0004,
				0x0005,
				0x0006,
				0x0007,
				0x0008,
				tt.args.sel0,
				tt.args.sel1,
				tt.args.sel2,
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
