package components

import (
	"reflect"
	"testing"
)

func TestMux4Way16(t *testing.T) {
	t.Parallel()

	type args struct {
		sel0 Val
		sel1 Val
	}

	tests := []struct {
		name     string
		args     args
		expected Val
	}{
		{
			"a: sel 0x00 => 0x0001",
			args{
				&SingleChan{false},
				&SingleChan{false},
			},
			&SixteenChan{val: 0x0001},
		},
		{
			"b: sel 0x01 => 0x0002",
			args{
				&SingleChan{true},
				&SingleChan{false},
			},
			&SixteenChan{val: 0x0002},
		},
		{
			"c: sel 0x02 => 0x0003",
			args{
				&SingleChan{false},
				&SingleChan{true},
			},
			&SixteenChan{val: 0x0003},
		},
		{
			"d: sel 0x03 => 0x0004",
			args{
				&SingleChan{true},
				&SingleChan{true},
			},
			&SixteenChan{val: 0x0004},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			or := NewMux4Way16()

			result := or.Update(
				UpdateOpts{TargetA, &SixteenChan{0x0001}},
				UpdateOpts{TargetB, &SixteenChan{0x0002}},
				UpdateOpts{TargetC, &SixteenChan{0x0003}},
				UpdateOpts{TargetD, &SixteenChan{0x0004}},
				UpdateOpts{TargetSel0, tt.args.sel0},
				UpdateOpts{TargetSel1, tt.args.sel1},
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
