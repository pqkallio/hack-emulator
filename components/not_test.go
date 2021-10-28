package components

import (
	"reflect"
	"testing"
)

func TestNot(t *testing.T) {
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
			"0 = 1",
			args{&SingleChan{val: false}},
			&SingleChan{val: true},
		},
		{
			"1 = 0",
			args{&SingleChan{val: true}},
			&SingleChan{val: false},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			not := Not{}

			result := not.Update(UpdateOpts{TargetIn, tt.args.input})

			if !reflect.DeepEqual(tt.expected, result) {
				t.Errorf("expected:\n%+v\ngot:\n%+v", tt.expected, result)
			}
		})
	}
}
