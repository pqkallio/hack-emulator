package components

import (
	"reflect"
	"testing"
)

func TestDemux8Way(t *testing.T) {
	t.Parallel()

	type args struct {
		in   Val
		sel0 Val
		sel1 Val
		sel2 Val
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
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
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
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
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
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{false},
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
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{false},
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
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{true},
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
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{true},
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
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{true},
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
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{true},
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

			demux8Way := NewDemux8Way()

			a, b, c, d, e, f, g, h := demux8Way.Update(
				UpdateOpts{TargetIn, tt.args.in},
				UpdateOpts{TargetSel0, tt.args.sel0},
				UpdateOpts{TargetSel1, tt.args.sel1},
				UpdateOpts{TargetSel2, tt.args.sel2},
			)

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

			if !reflect.DeepEqual(tt.expectedE, e) {
				t.Errorf("E: expected:\n%+v\ngot:\n%+v", tt.expectedE, e)
			}

			if !reflect.DeepEqual(tt.expectedF, f) {
				t.Errorf("B: expected:\n%+v\ngot:\n%+v", tt.expectedF, f)
			}

			if !reflect.DeepEqual(tt.expectedG, g) {
				t.Errorf("C: expected:\n%+v\ngot:\n%+v", tt.expectedG, g)
			}

			if !reflect.DeepEqual(tt.expectedH, h) {
				t.Errorf("D: expected:\n%+v\ngot:\n%+v", tt.expectedH, h)
			}
		})
	}
}
