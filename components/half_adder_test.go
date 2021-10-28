package components

import (
	"reflect"
	"testing"
)

func TestHalfAdder(t *testing.T) {
	t.Parallel()

	type args struct {
		a Val
		b Val
	}

	tests := []struct {
		name          string
		args          args
		expectedSum   Val
		expectedCarry Val
	}{
		{
			"0 + 0 = 00",
			args{&SingleChan{val: false}, &SingleChan{val: false}},
			&SingleChan{val: false},
			&SingleChan{val: false},
		},
		{
			"0 + 1 = 01",
			args{&SingleChan{val: false}, &SingleChan{val: true}},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"1 + 0 = 01",
			args{&SingleChan{val: true}, &SingleChan{val: false}},
			&SingleChan{val: true},
			&SingleChan{val: false},
		},
		{
			"1 + 1 = 10",
			args{&SingleChan{val: true}, &SingleChan{val: true}},
			&SingleChan{val: false},
			&SingleChan{val: true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			halfAdder := NewHalfAdder()

			sum, carry := halfAdder.Update(UpdateOpts{TargetA, tt.args.a}, UpdateOpts{TargetB, tt.args.b})

			if !reflect.DeepEqual(tt.expectedSum, sum) {
				t.Errorf("sum: expected:\n%+v\ngot:\n%+v", tt.expectedSum, sum)
			}

			if !reflect.DeepEqual(tt.expectedCarry, carry) {
				t.Errorf("carry: expected:\n%+v\ngot:\n%+v", tt.expectedSum, sum)
			}
		})
	}
}
