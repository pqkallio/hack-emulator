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
				&SelectChan{0},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 0, sel: 1 => a: 0, b: 0",
			args{
				&SingleChan{false},
				&SelectChan{1},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 0 => a: 1, b: 0",
			args{
				&SingleChan{true},
				&SelectChan{0},
			},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 1 => a: 0, b: 1",
			args{
				&SingleChan{true},
				&SelectChan{1},
			},
			&SingleChan{val: false},
			&SingleChan{val: true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			outA := MockOut{}
			outB := MockOut{}
			or := NewDemux(TargetIn, TargetIn, &outA, &outB)

			or.Update(
				UpdateOpts{TargetIn, tt.args.in},
				UpdateOpts{TargetSel, tt.args.sel},
			)

			if !reflect.DeepEqual(tt.expectedA, outA.Result) {
				t.Errorf("A: expected:\n%+v\ngot:\n%+v", tt.expectedA, outA.Result)
			}

			if !reflect.DeepEqual(tt.expectedB, outB.Result) {
				t.Errorf("B: expected:\n%+v\ngot:\n%+v", tt.expectedB, outB.Result)
			}
		})
	}
}
