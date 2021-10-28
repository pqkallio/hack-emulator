package components

import (
	"reflect"
	"testing"
)

func TestMux4Way16(t *testing.T) {
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
				UpdateOpts{TargetSel, tt.args.sel},
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
