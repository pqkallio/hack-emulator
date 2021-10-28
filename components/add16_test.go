package components

import (
	"reflect"
	"testing"
)

func TestAdd16(t *testing.T) {
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
			"0x0000 + 0x0000 = 0x0000",
			args{&SixteenChan{0x0000}, &SixteenChan{0x0000}},
			&SixteenChan{0x0000},
		},
		{
			"0x0000 + 0xffff = 0xffff",
			args{&SixteenChan{0x0000}, &SixteenChan{0xffff}},
			&SixteenChan{0xffff},
		},
		{
			"0xffff & 0x0000 = 0xffff",
			args{&SixteenChan{0xffff}, &SixteenChan{0x0000}},
			&SixteenChan{0xffff},
		},
		{
			"0xfffe & 0x0001 = 0xffff",
			args{&SixteenChan{0xfffe}, &SixteenChan{0x0001}},
			&SixteenChan{0xffff},
		},
		{
			"0x0001 & 0x0001 = 0x0002",
			args{&SixteenChan{0x0001}, &SixteenChan{0x0001}},
			&SixteenChan{0x0002},
		},
		{
			"0xffff + 0x0001 = 0x0000",
			args{&SixteenChan{0xffff}, &SixteenChan{0x0001}},
			&SixteenChan{0x0000},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			add16 := NewAdd16()

			result := add16.Update(UpdateOpts{TargetA, tt.args.a}, UpdateOpts{TargetB, tt.args.b})

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
