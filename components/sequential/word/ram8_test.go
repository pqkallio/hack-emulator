package word

import (
	"reflect"
	"testing"
)

func TestRAM8(t *testing.T) {
	t.Parallel()

	val := uint16(0x1234)

	type opts struct {
		in                        uint16
		load, addr0, addr1, addr2 bool
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
			"sel 0, load 0x1234",
			args{
				[]opts{
					{val, true, false, false, false},
					{val, false, false, false, false},
					{val, false, false, false, true},
					{val, false, false, true, false},
					{val, false, false, true, true},
					{val, false, true, false, false},
					{val, false, true, false, true},
					{val, false, true, true, false},
					{val, false, true, true, true},
				},
			},
			[]uint16{0, val, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			"sel 1, load 0x1234",
			args{
				[]opts{
					{val, true, false, false, true},
					{val, false, false, false, false},
					{val, false, false, false, true},
					{val, false, false, true, false},
					{val, false, false, true, true},
					{val, false, true, false, false},
					{val, false, true, false, true},
					{val, false, true, true, false},
					{val, false, true, true, true},
				},
			},
			[]uint16{0, 0, val, 0, 0, 0, 0, 0, 0},
		},
		{
			"sel 2, load 0x1234",
			args{
				[]opts{
					{val, true, false, true, false},
					{val, false, false, false, false},
					{val, false, false, false, true},
					{val, false, false, true, false},
					{val, false, false, true, true},
					{val, false, true, false, false},
					{val, false, true, false, true},
					{val, false, true, true, false},
					{val, false, true, true, true},
				},
			},
			[]uint16{0, 0, 0, val, 0, 0, 0, 0, 0},
		},
		{
			"sel 3, load 0x1234",
			args{
				[]opts{
					{val, true, false, true, true},
					{val, false, false, false, false},
					{val, false, false, false, true},
					{val, false, false, true, false},
					{val, false, false, true, true},
					{val, false, true, false, false},
					{val, false, true, false, true},
					{val, false, true, true, false},
					{val, false, true, true, true},
				},
			},
			[]uint16{0, 0, 0, 0, val, 0, 0, 0, 0},
		},
		{
			"sel 4, load 0x1234",
			args{
				[]opts{
					{val, true, true, false, false},
					{val, false, false, false, false},
					{val, false, false, false, true},
					{val, false, false, true, false},
					{val, false, false, true, true},
					{val, false, true, false, false},
					{val, false, true, false, true},
					{val, false, true, true, false},
					{val, false, true, true, true},
				},
			},
			[]uint16{0, 0, 0, 0, 0, val, 0, 0, 0},
		},
		{
			"sel 5, load 0x1234",
			args{
				[]opts{
					{val, true, true, false, true},
					{val, false, false, false, false},
					{val, false, false, false, true},
					{val, false, false, true, false},
					{val, false, false, true, true},
					{val, false, true, false, false},
					{val, false, true, false, true},
					{val, false, true, true, false},
					{val, false, true, true, true},
				},
			},
			[]uint16{0, 0, 0, 0, 0, 0, val, 0, 0},
		},
		{
			"sel 6, load 0x1234",
			args{
				[]opts{
					{val, true, true, true, false},
					{val, false, false, false, false},
					{val, false, false, false, true},
					{val, false, false, true, false},
					{val, false, false, true, true},
					{val, false, true, false, false},
					{val, false, true, false, true},
					{val, false, true, true, false},
					{val, false, true, true, true},
				},
			},
			[]uint16{0, 0, 0, 0, 0, 0, 0, val, 0},
		},
		{
			"sel 7, load 0x1234",
			args{
				[]opts{
					{val, true, true, true, true},
					{val, false, false, false, false},
					{val, false, false, false, true},
					{val, false, false, true, false},
					{val, false, false, true, true},
					{val, false, true, false, false},
					{val, false, true, false, true},
					{val, false, true, true, false},
					{val, false, true, true, true},
				},
			},
			[]uint16{0, 0, 0, 0, 0, 0, 0, 0, val},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ram8 := NewRAM8()

			for i, opt := range tt.args.opts {
				actual := ram8.Update(opt.in, opt.load, opt.addr0, opt.addr1, opt.addr2, nil, 0)

				if !reflect.DeepEqual(tt.expected[i], actual) {
					t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected[i], actual)
				}

				ram8.Tick()
			}
		})
	}
}
