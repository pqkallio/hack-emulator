package components

import (
	"reflect"
	"testing"
)

func TestDemux(t *testing.T) {
	t.Parallel()

	type args struct {
		in  Val
		sel Val
	}

	tests := []struct {
		name      string
		args      args
		expectedA Val
		expectedB Val
	}{
		{
			"in: 0, sel: 0 => a: 0, b: 0",
			args{
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 0, sel: 1 => a: 0, b: 0",
			args{
				&SingleChan{false},
				&SingleChan{true},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 0 => a: 1, b: 0",
			args{
				&SingleChan{true},
				&SingleChan{false},
			},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 1 => a: 0, b: 1",
			args{
				&SingleChan{true},
				&SingleChan{true},
			},
			&SingleChan{val: false},
			&SingleChan{val: true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			demux := NewDemux()

			a, b := demux.Update(
				UpdateOpts{TargetIn, tt.args.in},
				UpdateOpts{TargetSel0, tt.args.sel},
			)

			if !reflect.DeepEqual(tt.expectedA, a) {
				t.Errorf("A: expected:\n%+v\ngot:\n%+v", tt.expectedA, a)
			}

			if !reflect.DeepEqual(tt.expectedB, b) {
				t.Errorf("B: expected:\n%+v\ngot:\n%+v", tt.expectedB, b)
			}
		})
	}
}
