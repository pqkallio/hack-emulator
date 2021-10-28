package components

import (
	"reflect"
	"testing"
)

func TestMux8Way16(t *testing.T) {
	t.Parallel()

	type args struct {
		sel Val
	}

	tests := []struct {
		name     string
		args     args
		expected Val
	}{
		{
			"a: sel 0x00 => 0x0001",
			args{
				&SelectChan{0},
			},
			&SixteenChan{val: 0x0001},
		},
		{
			"b: sel 0x01 => 0x0002",
			args{
				&SelectChan{1},
			},
			&SixteenChan{val: 0x0002},
		},
		{
			"c: sel 0x02 => 0x0003",
			args{
				&SelectChan{2},
			},
			&SixteenChan{val: 0x0003},
		},
		{
			"d: sel 0x03 => 0x0004",
			args{
				&SelectChan{3},
			},
			&SixteenChan{val: 0x0004},
		},
		{
			"e: sel 0x04 => 0x0005",
			args{
				&SelectChan{4},
			},
			&SixteenChan{val: 0x0005},
		},
		{
			"f: sel 0x05 => 0x0006",
			args{
				&SelectChan{5},
			},
			&SixteenChan{val: 0x0006},
		},
		{
			"g: sel 0x06 => 0x0007",
			args{
				&SelectChan{6},
			},
			&SixteenChan{val: 0x0007},
		},
		{
			"h: sel 0x07 => 0x0008",
			args{
				&SelectChan{7},
			},
			&SixteenChan{val: 0x0008},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux8Way := NewMux8Way16()

			result := mux8Way.Update(
				UpdateOpts{TargetA, &SixteenChan{0x0001}},
				UpdateOpts{TargetB, &SixteenChan{0x0002}},
				UpdateOpts{TargetC, &SixteenChan{0x0003}},
				UpdateOpts{TargetD, &SixteenChan{0x0004}},
				UpdateOpts{TargetE, &SixteenChan{0x0005}},
				UpdateOpts{TargetF, &SixteenChan{0x0006}},
				UpdateOpts{TargetG, &SixteenChan{0x0007}},
				UpdateOpts{TargetH, &SixteenChan{0x0008}},
				UpdateOpts{TargetSel, tt.args.sel},
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
