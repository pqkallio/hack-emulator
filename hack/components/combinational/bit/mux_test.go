package bit

import (
	"reflect"
	"testing"
)

func TestMux(t *testing.T) {
	t.Parallel()

	type args struct {
		a   bool
		b   bool
		sel bool
	}

	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			"a: 0, b: 0, sel: 0 = 0",
			args{
				false,
				false,
				false,
			},
			false,
		},
		{
			"a: 0, b: 0, sel: 1 = 0",
			args{
				false,
				false,
				true,
			},
			false,
		},
		{
			"a: 0, b: 1, sel: 0 = 0",
			args{
				false,
				true,
				false,
			},
			false,
		},
		{
			"a: 0, b: 1, sel: 1 = 1",
			args{
				false,
				true,
				true,
			},
			true,
		},
		{
			"a: 1, b: 0, sel: 0 = 1",
			args{
				true,
				false,
				false,
			},
			true,
		},
		{
			"a: 1, b: 0, sel: 1 = 0",
			args{
				true,
				false,
				true,
			},
			false,
		},
		{
			"a: 1, b: 1, sel: 0 = 1",
			args{
				true,
				true,
				false,
			},
			true,
		},
		{
			"a: 1, b: 1, sel: 1 = 1",
			args{
				true,
				true,
				true,
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux := NewMux()

			result := mux.Update(tt.args.a, tt.args.b, tt.args.sel, nil, 0)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
