package bit

import (
	"reflect"
	"testing"
)

func TestFullAdder(t *testing.T) {
	t.Parallel()

	type args struct {
		a bool
		b bool
		c bool
	}

	tests := []struct {
		name          string
		args          args
		expectedSum   bool
		expectedCarry bool
	}{
		{
			"0 + 0 + 0 = 00",
			args{false, false, false},
			false,
			false,
		},
		{
			"0 + 0 + 1 = 01",
			args{false, false, true},
			true,
			false,
		},
		{
			"0 + 1 + 0 = 01",
			args{false, true, false},
			true,
			false,
		},
		{
			"0 + 1 + 1 = 10",
			args{false, true, true},
			false,
			true,
		},
		{
			"1 + 0 + 0 = 01",
			args{true, false, false},
			true,
			false,
		},
		{
			"1 + 0 + 1 = 10",
			args{true, false, true},
			false,
			true,
		},
		{
			"1 + 1 + 0 = 10",
			args{true, true, false},
			false,
			true,
		},
		{
			"1 + 1 + 1 = 11",
			args{true, true, true},
			true,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fullAdder := NewFullAdder()

			sum, carry := fullAdder.Update(tt.args.a, tt.args.b, tt.args.c)

			if !reflect.DeepEqual(tt.expectedSum, sum) {
				t.Errorf("sum: expected:\n%+v\ngot:\n%+v", tt.expectedSum, sum)
			}

			if !reflect.DeepEqual(tt.expectedCarry, carry) {
				t.Errorf("carry: expected:\n%+v\ngot:\n%+v", tt.expectedSum, sum)
			}
		})
	}
}
