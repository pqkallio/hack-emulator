package bit

import (
	"reflect"
	"testing"
)

func TestDFF(t *testing.T) {
	t.Parallel()

	type args struct {
		data bool
	}

	tests := []struct {
		name     string
		args     []args
		expected []bool
	}{
		{
			"load 0, 1, 0",
			[]args{
				{
					false,
				},
				{
					true,
				},
				{
					false,
				},
			},
			[]bool{
				false,
				false,
				true,
			},
		},
		{
			"load 1, 0, 1",
			[]args{
				{
					true,
				},
				{
					false,
				},
				{
					true,
				},
			},
			[]bool{
				false,
				true,
				false,
			},
		},
		{
			"load 1, 0, 0, 1",
			[]args{
				{
					true,
				},
				{
					false,
				},
				{
					false,
				},
				{
					true,
				},
			},
			[]bool{
				false,
				true,
				false,
				false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dff := NewDFF()

			for i, arg := range tt.args {
				actual := dff.Update(arg.data)

				expected := tt.expected[i]

				if !reflect.DeepEqual(expected, actual) {
					t.Errorf("expected:\n%+v\ngot:\n%+v", expected, actual)
				}

				dff.Tick()
			}
		})
	}
}
