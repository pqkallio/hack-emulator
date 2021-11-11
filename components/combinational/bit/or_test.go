package bit

import (
	"reflect"
	"testing"
)

func TestOr(t *testing.T) {
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
			"0 | 0 = 0",
			args{false, false},
			false,
		},
		{
			"0 | 1 = 1",
			args{false, true},
			true,
		},
		{
			"1 | 0 = 1",
			args{true, false},
			true,
		},
		{
			"1 | 1 = 1",
			args{true, true},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			or := NewOr()

			result := or.Update(tt.args.a, tt.args.b)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
