package bit

import (
	"reflect"
	"testing"
)

func TestDemux4Way(t *testing.T) {
	t.Parallel()

	type args struct {
		in   bool
		sel0 bool
		sel1 bool
	}

	tests := []struct {
		name      string
		args      args
		expectedA bool
		expectedB bool
		expectedC bool
		expectedD bool
	}{
		{
			"in: 1, sel: 0 => a: 1, b: 0, c: 0, d: 0",
			args{
				true,
				false,
				false,
			},
			true,
			false,
			false,
			false,
		},
		{
			"in: 1, sel: 1 => a: 0, b: 1, c: 0, d: 0",
			args{
				true,
				true,
				false,
			},
			false,
			true,
			false,
			false,
		},
		{
			"in: 1, sel: 2 => a: 0, b: 0, c: 1, d: 0",
			args{
				true,
				false,
				true,
			},
			false,
			false,
			true,
			false,
		},
		{
			"in: 1, sel: 3 => a: 0, b: 0, c: 0, d: 1",
			args{
				true,
				true,
				true,
			},
			false,
			false,
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			demux4Way := NewDemux4Way()

			a, b, c, d := demux4Way.Update(tt.args.in, tt.args.sel0, tt.args.sel1)

			if !reflect.DeepEqual(tt.expectedA, a) {
				t.Errorf("A: expected:\n%+v\ngot:\n%+v", tt.expectedA, a)
			}

			if !reflect.DeepEqual(tt.expectedB, b) {
				t.Errorf("B: expected:\n%+v\ngot:\n%+v", tt.expectedB, b)
			}

			if !reflect.DeepEqual(tt.expectedC, c) {
				t.Errorf("C: expected:\n%+v\ngot:\n%+v", tt.expectedC, c)
			}

			if !reflect.DeepEqual(tt.expectedD, d) {
				t.Errorf("D: expected:\n%+v\ngot:\n%+v", tt.expectedD, d)
			}
		})
	}
}
