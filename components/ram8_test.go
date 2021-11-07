package components

import (
	"reflect"
	"testing"
)

func TestRAM8(t *testing.T) {
	t.Parallel()

	opts := func(load bool, data uint16, addr []UpdateOpts) []UpdateOpts {
		return append([]UpdateOpts{
			{TargetIn, &SixteenChan{data}},
			{TargetLoad, &SingleChan{load}},
		}, addr...)
	}

	selAddr0 := []UpdateOpts{
		{TargetSel0, &SingleChan{false}},
		{TargetSel1, &SingleChan{false}},
		{TargetSel2, &SingleChan{false}},
	}

	selAddr1 := []UpdateOpts{
		{TargetSel0, &SingleChan{true}},
		{TargetSel1, &SingleChan{false}},
		{TargetSel2, &SingleChan{false}},
	}

	selAddr2 := []UpdateOpts{
		{TargetSel0, &SingleChan{false}},
		{TargetSel1, &SingleChan{true}},
		{TargetSel2, &SingleChan{false}},
	}

	selAddr3 := []UpdateOpts{
		{TargetSel0, &SingleChan{true}},
		{TargetSel1, &SingleChan{true}},
		{TargetSel2, &SingleChan{false}},
	}

	selAddr4 := []UpdateOpts{
		{TargetSel0, &SingleChan{false}},
		{TargetSel1, &SingleChan{false}},
		{TargetSel2, &SingleChan{true}},
	}

	selAddr5 := []UpdateOpts{
		{TargetSel0, &SingleChan{true}},
		{TargetSel1, &SingleChan{false}},
		{TargetSel2, &SingleChan{true}},
	}

	selAddr6 := []UpdateOpts{
		{TargetSel0, &SingleChan{false}},
		{TargetSel1, &SingleChan{true}},
		{TargetSel2, &SingleChan{true}},
	}

	selAddr7 := []UpdateOpts{
		{TargetSel0, &SingleChan{true}},
		{TargetSel1, &SingleChan{true}},
		{TargetSel2, &SingleChan{true}},
	}

	type args struct {
		opts [][]UpdateOpts
	}

	tests := []struct {
		name     string
		args     args
		expected []Val
	}{
		{
			"sel 0, load 0x1234",
			args{
				[][]UpdateOpts{
					opts(true, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr7),
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0x1234},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
			},
		},
		{
			"sel 1, load 0x1234",
			args{
				[][]UpdateOpts{
					opts(true, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr7),
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0x1234},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
			},
		},
		{
			"sel 2, load 0x1234",
			args{
				[][]UpdateOpts{
					opts(true, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr7),
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0x1234},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
			},
		},
		{
			"sel 3, load 0x1234",
			args{
				[][]UpdateOpts{
					opts(true, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr7),
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0x1234},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
			},
		},
		{
			"sel 4, load 0x1234",
			args{
				[][]UpdateOpts{
					opts(true, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr7),
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0x1234},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
			},
		},
		{
			"sel 5, load 0x1234",
			args{
				[][]UpdateOpts{
					opts(true, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr7),
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0x1234},
				&SixteenChan{0},
				&SixteenChan{0},
			},
		},
		{
			"sel 6, load 0x1234",
			args{
				[][]UpdateOpts{
					opts(true, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr7),
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0x1234},
				&SixteenChan{0},
			},
		},
		{
			"sel 7, load 0x1234",
			args{
				[][]UpdateOpts{
					opts(true, 0x1234, selAddr7),
					opts(false, 0x1234, selAddr0),
					opts(false, 0x1234, selAddr1),
					opts(false, 0x1234, selAddr2),
					opts(false, 0x1234, selAddr3),
					opts(false, 0x1234, selAddr4),
					opts(false, 0x1234, selAddr5),
					opts(false, 0x1234, selAddr6),
					opts(false, 0x1234, selAddr7),
				},
			},
			[]Val{
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0},
				&SixteenChan{0x1234},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ram8 := NewRAM8()

			for i, opt := range tt.args.opts {
				actual := ram8.Update(opt...)

				if !reflect.DeepEqual(tt.expected[i], actual) {
					t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected[i], actual)
				}

				ram8.Tick()
			}
		})
	}
}
