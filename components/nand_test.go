package components

import (
	"reflect"
	"testing"
)

func TestNand(t *testing.T) {
	t.Parallel()

	type args struct {
		a Val
		b Val
	}

	tests := []struct {
		name     string
		args     args
		expected Val
	}{
		{
			"nand(0, 0) = 1",
			args{&SingleChan{val: false}, &SingleChan{val: false}},
			&SingleChan{val: true},
		},
		{
			"nand(0, 1) = 1",
			args{&SingleChan{val: false}, &SingleChan{val: true}},
			&SingleChan{val: true},
		},
		{
			"nand(1, 0) = 1",
			args{&SingleChan{val: true}, &SingleChan{val: false}},
			&SingleChan{val: true},
		},
		{
			"nand(1, 1) = 0",
			args{&SingleChan{val: true}, &SingleChan{val: true}},
			&SingleChan{val: false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			nand := NewNand()

			result := nand.Update(UpdateOpts{TargetA, tt.args.a}, UpdateOpts{TargetB, tt.args.b})

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
