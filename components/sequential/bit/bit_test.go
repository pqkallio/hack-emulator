package bit

import (
	"reflect"
	"testing"
)

func TestBit(t *testing.T) {
	t.Parallel()

	type opts struct {
		data, load bool
	}

	type args struct {
		opts []opts
	}

	tests := []struct {
		name     string
		args     args
		expected []bool
	}{
		{
			"load 1",
			args{
				[]opts{
					{
						true,
						true,
					},
					{
						false,
						true,
					},
				},
			},
			[]bool{
				false,
				true,
			},
		},
		{
			"don't load 1",
			args{
				[]opts{
					{
						true,
						false,
					},
					{
						false,
						true,
					},
				},
			},
			[]bool{
				false,
				false,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bit := NewBit()

			for i, opt := range tt.args.opts {
				actual := bit.Update(opt.data, opt.load, nil, 0)

				if !reflect.DeepEqual(tt.expected[i], actual) {
					t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected[i], actual)
				}

				bit.Tick(nil)
			}
		})
	}
}
