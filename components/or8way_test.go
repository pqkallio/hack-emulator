package components

import (
	"reflect"
	"testing"
)

func TestOr8Way(t *testing.T) {
	t.Parallel()

	type args struct {
		a, b, c, d, e, f, g, h Val
	}

	tests := []struct {
		name     string
		args     args
		expected Val
	}{
		{
			"0|0|0|0|0|0|0|0 = 0",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{false},
		},
		{
			"0|0|0|0|0|0|0|1 = 1",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{true},
			},
			&SingleChan{true},
		},
		{
			"0|0|0|0|0|0|1|0 = 1",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"0|0|0|0|0|1|0|0 = 1",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"0|0|0|0|1|0|0|0 = 1",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"0|0|0|1|0|0|0|0 = 1",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"0|0|1|0|0|0|0|0 = 1",
			args{
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"0|1|0|0|0|0|0|0 = 1",
			args{
				&SingleChan{false},
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"1|0|0|0|0|0|0|0 = 1",
			args{
				&SingleChan{true},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
				&SingleChan{false},
			},
			&SingleChan{true},
		},
		{
			"1|1|1|1|1|1|1|1 = 1",
			args{
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{true},
				&SingleChan{true},
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

			or8Way := NewOr8Way()

			result := or8Way.Update(
				UpdateOpts{TargetA, tt.args.a},
				UpdateOpts{TargetB, tt.args.b},
				UpdateOpts{TargetC, tt.args.c},
				UpdateOpts{TargetD, tt.args.d},
				UpdateOpts{TargetE, tt.args.e},
				UpdateOpts{TargetF, tt.args.f},
				UpdateOpts{TargetG, tt.args.g},
				UpdateOpts{TargetH, tt.args.h},
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
