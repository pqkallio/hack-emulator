package components

import (
	"reflect"
	"testing"
)

func TestAnd16(t *testing.T) {
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
			"0x0000 & 0x0000 = 0x0000",
			args{&SixteenChan{0x0000}, &SixteenChan{0x0000}},
			&SixteenChan{0x0000},
		},
		{
			"0x0000 & 0xffff = 0x0000",
			args{&SixteenChan{0x0000}, &SixteenChan{0xffff}},
			&SixteenChan{0x0000},
		},
		{
			"0xffff & 0x0000 = 0x0000",
			args{&SixteenChan{0xffff}, &SixteenChan{0x0000}},
			&SixteenChan{0x0000},
		},
		{
			"0xffff & 0xffff = 0xffff",
			args{&SixteenChan{0xffff}, &SixteenChan{0xffff}},
			&SixteenChan{0xffff},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			and := NewAnd16()

			result := and.Update(UpdateOpts{TargetA, tt.args.a}, UpdateOpts{TargetB, tt.args.b})

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
