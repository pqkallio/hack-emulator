package components

import (
	"reflect"
	"testing"
)

func TestInc16(t *testing.T) {
	t.Parallel()

	type args struct {
		in Val
	}

	tests := []struct {
		name     string
		args     args
		expected Val
	}{
		{
			"inc 0x0000 = 0x0001",
			args{&SixteenChan{0x0000}},
			&SixteenChan{0x0001},
		},
		{
			"inc 0x0001 = 0x0002",
			args{&SixteenChan{0x0001}},
			&SixteenChan{0x0002},
		},
		{
			"inc 0xffff = 0x0000",
			args{&SixteenChan{0xffff}},
			&SixteenChan{0x0000},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inc16 := NewInc16()

			result := inc16.Update(UpdateOpts{TargetIn, tt.args.in})

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
