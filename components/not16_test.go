package components

import (
	"reflect"
	"testing"
)

func TestNot16(t *testing.T) {
	t.Parallel()

	type args struct {
		input Val
	}

	tests := []struct {
		name     string
		args     args
		expected Val
	}{
		{
			"0x0000 = 0xffff",
			args{&SixteenChan{0x0000}},
			&SixteenChan{0xffff},
		},
		{
			"0xffff = 0x0000",
			args{&SixteenChan{0xffff}},
			&SixteenChan{0x0000},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			not := NewNot16()

			result := not.Update(UpdateOpts{TargetIn, tt.args.input})

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
