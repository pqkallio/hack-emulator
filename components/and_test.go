package components

import (
	"reflect"
	"testing"
)

func TestAnd(t *testing.T) {
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
			"0 & 1 = 0",
			args{&SingleChan{val: false}, &SingleChan{val: true}},
			&SingleChan{val: false},
		},
		{
			"1 & 0 = 0",
			args{&SingleChan{val: true}, &SingleChan{val: false}},
			&SingleChan{val: false},
		},
		{
			"1 & 1 = 1",
			args{&SingleChan{val: true}, &SingleChan{val: true}},
			&SingleChan{val: true},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			and := NewAnd()

			result := and.Update(UpdateOpts{TargetA, tt.args.a}, UpdateOpts{TargetB, tt.args.b})

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
