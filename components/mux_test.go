package components

import (
	"reflect"
	"testing"
)

func TestMux(t *testing.T) {
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
			"a: 0, b: 0, sel: 0 = 0",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{val: false},
		},
		{
			"a: 0, b: 0, sel: 1 = 0",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{true},
			},
			&SingleChan{val: false},
		},
		{
			"a: 0, b: 1, sel: 0 = 0",
			args{
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{false},
			},
			&SingleChan{val: false},
		},
		{
			"a: 0, b: 1, sel: 1 = 1",
			args{
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{true},
			},
			&SingleChan{true},
		},
		{
			"a: 1, b: 0, sel: 0 = 1",
			args{
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"a: 1, b: 0, sel: 1 = 0",
			args{
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{true},
			},
			&SingleChan{false},
		},
		{
			"a: 1, b: 1, sel: 0 = 1",
			args{
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"a: 1, b: 1, sel: 1 = 1",
			args{
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{true},
			},
			&SingleChan{true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux := NewMux()

			result := mux.Update(
				UpdateOpts{TargetA, tt.args.a},
				UpdateOpts{TargetB, tt.args.b},
				UpdateOpts{TargetSel0, tt.args.sel},
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
