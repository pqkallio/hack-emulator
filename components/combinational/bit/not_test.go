package bit

import (
	"reflect"
	"testing"
)

func TestNot(t *testing.T) {
	t.Parallel()

	type args struct {
		in bool
	}

	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			"0 = 1",
			args{false},
			true,
		},
		{
			"1 = 0",
			args{true},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			not := NewNot()

			result := not.Update(tt.args.in)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
