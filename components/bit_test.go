package components

import (
	"reflect"
	"testing"
)

func TestBit(t *testing.T) {
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
			"load 1",
			args{
				[][]UpdateOpts{
					{
						{TargetIn, &SingleChan{true}},
						{TargetLoad, &SingleChan{true}},
					},
					{
						{TargetIn, &SingleChan{false}},
						{TargetLoad, &SingleChan{true}},
					},
				},
			},
			[]Val{
				&SingleChan{false},
				&SingleChan{true},
			},
		},
		{
			"don't load 1, explicitly",
			args{
				[][]UpdateOpts{
					{
						{TargetIn, &SingleChan{true}},
						{TargetLoad, &SingleChan{false}},
					},
					{
						{TargetIn, &SingleChan{false}},
						{TargetLoad, &SingleChan{true}},
					},
				},
			},
			[]Val{
				&SingleChan{false},
				&SingleChan{false},
			},
		},
		{
			"don't load 1, implicitly",
			args{
				[][]UpdateOpts{
					{
						{TargetIn, &SingleChan{true}},
					},
					{
						{TargetIn, &SingleChan{false}},
						{TargetLoad, &SingleChan{true}},
					},
				},
			},
			[]Val{
				&SingleChan{false},
				&SingleChan{false},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bit := NewBit()

			for i, opt := range tt.args.opts {
				actual := bit.Update(opt...)

				if !reflect.DeepEqual(tt.expected[i], actual) {
					t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected[i], actual)
				}

				bit.Tick()
			}
		})
	}
}
