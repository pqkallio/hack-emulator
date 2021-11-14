package bit

import (
	"reflect"
	"testing"
)

func TestAnd(t *testing.T) {
	t.Parallel()

	type args struct {
		a bool
		b bool
	}

	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			"0 & 0 = 0",
			args{false, false},
			false,
		},
		{
			"0 & 1 = 0",
			args{false, true},
			false,
		},
		{
			"1 & 0 = 0",
			args{true, false},
			false,
		},
		{
			"1 & 1 = 1",
			args{true, true},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			and := NewAnd()

			result := and.Update(tt.args.a, tt.args.b, nil, 0)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
