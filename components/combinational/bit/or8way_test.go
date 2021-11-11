package bit

import (
	"reflect"
	"testing"
)

func TestOr8Way(t *testing.T) {
	t.Parallel()

	type args struct {
		a, b, c, d, e, f, g, h bool
	}

	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			"0|0|0|0|0|0|0|0 = 0",
			args{
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			false,
		},
		{
			"0|0|0|0|0|0|0|1 = 1",
			args{
				false,
				false,
				false,
				false,
				false,
				false,
				false,
				true,
			},
			true,
		},
		{
			"0|0|0|0|0|0|1|0 = 1",
			args{
				false,
				false,
				false,
				false,
				false,
				false,
				true,
				false,
			},
			true,
		},
		{
			"0|0|0|0|0|1|0|0 = 1",
			args{
				false,
				false,
				false,
				false,
				false,
				true,
				false,
				false,
			},
			true,
		},
		{
			"0|0|0|0|1|0|0|0 = 1",
			args{
				false,
				false,
				false,
				false,
				true,
				false,
				false,
				false,
			},
			true,
		},
		{
			"0|0|0|1|0|0|0|0 = 1",
			args{
				false,
				false,
				false,
				true,
				false,
				false,
				false,
				false,
			},
			true,
		},
		{
			"0|0|1|0|0|0|0|0 = 1",
			args{
				false,
				false,
				true,
				false,
				false,
				false,
				false,
				false,
			},
			true,
		},
		{
			"0|1|0|0|0|0|0|0 = 1",
			args{
				false,
				true,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			true,
		},
		{
			"1|0|0|0|0|0|0|0 = 1",
			args{
				true,
				false,
				false,
				false,
				false,
				false,
				false,
				false,
			},
			true,
		},
		{
			"1|1|1|1|1|1|1|1 = 1",
			args{
				true,
				true,
				true,
				true,
				true,
				true,
				true,
				true,
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			or8Way := NewOr8Way()

			result := or8Way.Update(
				tt.args.a,
				tt.args.b,
				tt.args.c,
				tt.args.d,
				tt.args.e,
				tt.args.f,
				tt.args.g,
				tt.args.h,
			)

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
