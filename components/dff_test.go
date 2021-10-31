package components

import (
	"reflect"
	"testing"
)

func TestDFF(t *testing.T) {
	t.Parallel()

	type args struct {
		data, load bool
	}

	tests := []struct {
		name     string
		args     []args
		expected []Val
	}{
		{
			"load 0",
			[]args{
				{
					false,
					true,
				},
				{
					true,
					false,
				},
				{
					false,
					true,
				},
			},
			[]Val{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
			},
		},
		{
			"load 1",
			[]args{
				{
					true,
					true,
				},
				{
					false,
					false,
				},
				{
					true,
					true,
				},
			},
			[]Val{
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{true},
			},
		},
		{
			"load 1, 0",
			[]args{
				{
					true,
					true,
				},
				{
					false,
					false,
				},
				{
					false,
					true,
				},
				{
					true,
					false,
				},
			},
			[]Val{
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dff := NewDFF()

			for i, arg := range tt.args {
				dff.Update(
					UpdateOpts{TargetData, &SingleChan{arg.data}},
					UpdateOpts{TargetLoad, &SingleChan{arg.load}},
				)

				actual := dff.Get()
				expected := tt.expected[i]

				if !reflect.DeepEqual(expected, actual) {
					t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
				}
			}
		})
	}
}
