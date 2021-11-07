package components

import (
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	type args struct {
		opts [][]UpdateOpts
	}

	tests := []struct {
		name     string
		args     args
		expected []Val
	}{
		{
			"load 10101010_10101010",
			args{
				[][]UpdateOpts{
					{
						{TargetIn, &SixteenChan{0b10101010_10101010}},
						{TargetLoad, &SingleChan{true}},
					},
					{
						{TargetIn, &SixteenChan{0}},
						{TargetLoad, &SingleChan{true}},
					},
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0xaaaa},
			},
		},
		{
			"load 01010101_01010101",
			args{
				[][]UpdateOpts{
					{
						{TargetIn, &SixteenChan{0b0101010101010101}},
						{TargetLoad, &SingleChan{true}},
					},
					{
						{TargetIn, &SixteenChan{0}},
						{TargetLoad, &SingleChan{true}},
					},
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0x5555},
			},
		},
		{
			"don't load 10101010_10101010",
			args{
				[][]UpdateOpts{
					{
						{TargetIn, &SixteenChan{0b10101010_10101010}},
						{TargetLoad, &SingleChan{false}},
					},
					{
						{TargetIn, &SixteenChan{0}},
						{TargetLoad, &SingleChan{true}},
					},
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
			},
		},
		{
			"don't load 01010101_01010101",
			args{
				[][]UpdateOpts{
					{
						{TargetIn, &SixteenChan{0b0101010101010101}},
					},
					{
						{TargetIn, &SixteenChan{0}},
						{TargetLoad, &SingleChan{true}},
					},
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			reg := NewRegister()

			for i, opt := range tt.args.opts {
				actual := reg.Update(opt...)

				if !reflect.DeepEqual(tt.expected[i], actual) {
					t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected[i], actual)
				}

				reg.Tick()
			}
		})
	}
}
