package bit

import (
	"reflect"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	t.Parallel()

	type args struct {
		a bool
		b bool
	}

	tests := []struct {
		name          string
		args          args
		expectedSum   bool
		expectedCarry bool
	}{
		{
			"0 + 0 = 00",
			args{false, false},
			false,
			false,
		},
		{
			"0 + 1 = 01",
			args{false, true},
			true,
			false,
		},
		{
			"1 + 0 = 01",
			args{true, false},
			true,
			false,
		},
		{
			"1 + 1 = 10",
			args{true, true},
			false,
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			halfAdder := NewHalfAdder()

			sum, carry := halfAdder.Update(tt.args.a, tt.args.b)

			if !reflect.DeepEqual(tt.expectedSum, sum) {
				t.Errorf("sum: expected:\n%+v\ngot:\n%+v", tt.expectedSum, sum)
			}

			if !reflect.DeepEqual(tt.expectedCarry, carry) {
				t.Errorf("carry: expected:\n%+v\ngot:\n%+v", tt.expectedSum, sum)
			}
		})
	}
}
