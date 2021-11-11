package word

import (
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	type opts struct {
		in   uint16
		load bool
	}

	type args struct {
		opts []opts
	}

	tests := []struct {
		name     string
		args     args
		expected []uint16
	}{
		{
			"load 10101010_10101010",
			args{
				[]opts{
					{
						0b10101010_10101010,
						true,
					},
					{
						0,
						true,
					},
				},
			},
			[]uint16{
				0,
				0xaaaa,
			},
		},
		{
			"load 01010101_01010101",
			args{
				[]opts{
					{
						0b0101010101010101,
						true,
					},
					{
						0,
						true,
					},
				},
			},
			[]uint16{
				0,
				0x5555,
			},
		},
		{
			"don't load 10101010_10101010",
			args{
				[]opts{
					{
						0b10101010_10101010,
						false,
					},
					{
						0,
						true,
					},
				},
			},
			[]uint16{
				0,
				0,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			reg := NewRegister()

			for i, opt := range tt.args.opts {
				actual := reg.Update(opt.in, opt.load)

				if !reflect.DeepEqual(tt.expected[i], actual) {
					t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected[i], actual)
				}

				reg.Tick()
			}
		})
	}
}
