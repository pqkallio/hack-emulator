package components

import (
	"reflect"
	"testing"
)

func TestDemux8Way(t *testing.T) {
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
		expectedE Val
		expectedF Val
		expectedG Val
		expectedH Val
	}{
		{
			"in: 1, sel: 0 => a: 1, b: 0, c: 0, d: 0, e: 0, f: 0, g: 0, h: 0",
			args{
				&SingleChan{true},
				&SelectChan{0},
			},
			&SingleChan{val: true},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 1 => a: 0, b: 1, c: 0, d: 0, e: 0, f: 0, g: 0, h: 0",
			args{
				&SingleChan{true},
				&SelectChan{1},
			},
			&SingleChan{val: false},
			&SingleChan{val: true},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 2 => a: 0, b: 0, c: 1, d: 0, e: 0, f: 0, g: 0, h: 0",
			args{
				&SingleChan{true},
				&SelectChan{2},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: true},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 3 => a: 0, b: 0, c: 0, d: 1, e: 0, f: 0, g: 0, h: 0",
			args{
				&SingleChan{true},
				&SelectChan{3},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: true},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 4 => a: 0, b: 0, c: 0, d: 0, e: 1, f: 0, g: 0, h: 0",
			args{
				&SingleChan{true},
				&SelectChan{4},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: true},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 5 => a: 0, b: 0, c: 0, d: 0, e: 0, f: 1, g: 0, h: 0",
			args{
				&SingleChan{true},
				&SelectChan{5},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: true},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 6 => a: 0, b: 0, c: 0, d: 0, e: 0, f: 0, g: 1, h: 0",
			args{
				&SingleChan{true},
				&SelectChan{6},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"in: 1, sel: 7 => a: 0, b: 0, c: 0, d: 0, e: 0, f: 0, g: 0, h: 1",
			args{
				&SingleChan{true},
				&SelectChan{7},
			},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
			&SingleChan{val: false},
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
			outE := MockOut{}
			outF := MockOut{}
			outG := MockOut{}
			outH := MockOut{}

			demux8Way := NewDemux8Way(
				TargetIn,
				TargetIn,
				TargetIn,
				TargetIn,
				TargetIn,
				TargetIn,
				TargetIn,
				TargetIn,
				&outA,
				&outB,
				&outC,
				&outD,
				&outE,
				&outF,
				&outG,
				&outH,
			)

			demux8Way.Update(
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

			if !reflect.DeepEqual(tt.expectedE, outE.Result) {
				t.Errorf("E: expected:\n%+v\ngot:\n%+v", tt.expectedE, outE.Result)
			}

			if !reflect.DeepEqual(tt.expectedF, outF.Result) {
				t.Errorf("B: expected:\n%+v\ngot:\n%+v", tt.expectedF, outF.Result)
			}

			if !reflect.DeepEqual(tt.expectedG, outG.Result) {
				t.Errorf("C: expected:\n%+v\ngot:\n%+v", tt.expectedG, outG.Result)
			}

			if !reflect.DeepEqual(tt.expectedH, outH.Result) {
				t.Errorf("D: expected:\n%+v\ngot:\n%+v", tt.expectedH, outH.Result)
			}
		})
	}
}
