package components

import (
	"reflect"
	"testing"
)

func TestXor(t *testing.T) {
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
			"0 & 0 = 0",
			args{&SingleChan{val: false}, &SingleChan{val: false}},
			&SingleChan{val: false},
		},
		{
			"0 & 1 = 1",
			args{&SingleChan{val: false}, &SingleChan{val: true}},
			&SingleChan{val: true},
		},
		{
			"1 & 0 = 1",
			args{&SingleChan{val: true}, &SingleChan{val: false}},
			&SingleChan{val: true},
		},
		{
			"1 & 1 = 0",
			args{&SingleChan{val: true}, &SingleChan{val: true}},
			&SingleChan{val: false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			xor := NewXor()

			result := xor.Update(UpdateOpts{TargetA, tt.args.a}, UpdateOpts{TargetB, tt.args.b})

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
