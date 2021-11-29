package bit

import (
	"reflect"
	"testing"
)

func TestDemux(t *testing.T) {
	t.Parallel()

	type args struct {
		in  bool
		sel bool
	}

	tests := []struct {
		name      string
		args      args
		expectedA bool
		expectedB bool
	}{
		{
			"in: 0, sel: 0 => a: 0, b: 0",
			args{
				false,
				false,
			},
			false,
			false,
		},
		{
			"in: 0, sel: 1 => a: 0, b: 0",
			args{
				false,
				true,
			},
			false,
			false,
		},
		{
			"in: 1, sel: 0 => a: 1, b: 0",
			args{
				true,
				false,
			},
			true,
			false,
		},
		{
			"in: 1, sel: 1 => a: 0, b: 1",
			args{
				true,
				true,
			},
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			demux := NewDemux()

			a, b := demux.Update(tt.args.in, tt.args.sel)

			if !reflect.DeepEqual(tt.expectedA, a) {
				t.Errorf("A: expected:\n%+v\ngot:\n%+v", tt.expectedA, a)
			}

			if !reflect.DeepEqual(tt.expectedB, b) {
				t.Errorf("B: expected:\n%+v\ngot:\n%+v", tt.expectedB, b)
			}
		})
	}
}
