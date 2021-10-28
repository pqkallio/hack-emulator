package components

import (
	"reflect"
	"testing"
)

func TestFullAdder(t *testing.T) {
	t.Parallel()

	type args struct {
		a Val
		b Val
		c Val
	}

	tests := []struct {
		name          string
		args          args
		expectedSum   Val
		expectedCarry Val
	}{
		{
			"0 + 0 + 0 = 00",
			args{&SingleChan{false}, &SingleChan{false}, &SingleChan{false}},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"0 + 0 + 1 = 01",
			args{&SingleChan{false}, &SingleChan{false}, &SingleChan{true}},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"0 + 1 + 0 = 01",
			args{&SingleChan{false}, &SingleChan{true}, &SingleChan{false}},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"0 + 1 + 1 = 10",
			args{&SingleChan{false}, &SingleChan{true}, &SingleChan{true}},
			&SingleChan{val: false},
			&SingleChan{val: true},
		},
		{
			"1 + 0 + 0 = 01",
			args{&SingleChan{true}, &SingleChan{false}, &SingleChan{false}},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"1 + 0 + 1 = 10",
			args{&SingleChan{true}, &SingleChan{false}, &SingleChan{true}},
			&SingleChan{val: false},
			&SingleChan{val: true},
		},
		{
			"1 + 1 + 0 = 10",
			args{&SingleChan{true}, &SingleChan{true}, &SingleChan{false}},
			&SingleChan{val: false},
			&SingleChan{val: true},
		},
		{
			"1 + 1 + 1 = 11",
			args{&SingleChan{true}, &SingleChan{true}, &SingleChan{true}},
			&SingleChan{val: true},
			&SingleChan{val: true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fullAdder := NewFullAdder()

			sum, carry := fullAdder.Update(
				UpdateOpts{TargetA, tt.args.a},
				UpdateOpts{TargetB, tt.args.b},
				UpdateOpts{TargetC, tt.args.c},
			)

			if !reflect.DeepEqual(tt.expectedSum, sum) {
				t.Errorf("sum: expected:\n%+v\ngot:\n%+v", tt.expectedSum, sum)
			}

			if !reflect.DeepEqual(tt.expectedCarry, carry) {
				t.Errorf("carry: expected:\n%+v\ngot:\n%+v", tt.expectedSum, sum)
			}
		})
	}
}
