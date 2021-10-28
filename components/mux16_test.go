package components

import (
	"reflect"
	"testing"
)

func TestMux16(t *testing.T) {
	t.Parallel()

	type args struct {
		a   Val
		b   Val
		sel Val
	}

	tests := []struct {
		name     string
		args     args
		expected Val
	}{
		{
			"a: 0x0001, b: 0x0010, sel: 0 = 0x0001",
			args{
				&SixteenChan{0x0001},
				&SixteenChan{0x0010},
				&SelectChan{0},
			},
			&SixteenChan{val: 0x0001},
		},
		{
			"a: 0x0001, b: 0x0010, sel: 1 = 0x0010",
			args{
				&SixteenChan{0x0001},
				&SixteenChan{0x0010},
				&SelectChan{1},
			},
			&SixteenChan{val: 0x0010},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux16 := NewMux16()

			result := mux16.Update(
				UpdateOpts{TargetA, tt.args.a},
				UpdateOpts{TargetB, tt.args.b},
				UpdateOpts{TargetSel, tt.args.sel},
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
