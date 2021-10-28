package components

import (
	"reflect"
	"testing"
)

func TestDemux4Way(t *testing.T) {
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
		expectedC Val
		expectedD Val
	}{
		{
			"in: 1, sel: 0 => a: 1, b: 0, c: 0, d: 0",
			args{
				&SingleChan{true},
				&SelectChan{0},
			},
			&SingleChan{val: true},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 1 => a: 0, b: 1, c: 0, d: 0",
			args{
				&SingleChan{true},
				&SelectChan{1},
			},
			&SingleChan{val: false},
			&SingleChan{val: true},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 2 => a: 0, b: 0, c: 1, d: 0",
			args{
				&SingleChan{true},
				&SelectChan{2},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 3 => a: 0, b: 0, c: 0, d: 1",
			args{
				&SingleChan{true},
				&SelectChan{3},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
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
			outC := MockOut{}
			outD := MockOut{}

			demux4Way := NewDemux4Way(
				TargetIn,
				TargetIn,
				TargetIn,
				TargetIn,
				&outA,
				&outB,
				&outC,
				&outD,
			)

			demux4Way.Update(
				UpdateOpts{TargetIn, tt.args.in},
				UpdateOpts{TargetSel, tt.args.sel},
			)

			if !reflect.DeepEqual(tt.expectedA, outA.Result) {
				t.Errorf("A: expected:\n%+v\ngot:\n%+v", tt.expectedA, outA.Result)
			}

			if !reflect.DeepEqual(tt.expectedB, outB.Result) {
				t.Errorf("B: expected:\n%+v\ngot:\n%+v", tt.expectedB, outB.Result)
			}

			if !reflect.DeepEqual(tt.expectedC, outC.Result) {
				t.Errorf("C: expected:\n%+v\ngot:\n%+v", tt.expectedC, outC.Result)
			}

			if !reflect.DeepEqual(tt.expectedD, outD.Result) {
				t.Errorf("D: expected:\n%+v\ngot:\n%+v", tt.expectedD, outD.Result)
			}
		})
	}
}
