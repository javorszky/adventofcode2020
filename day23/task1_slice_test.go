package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createCircleSlice(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "creates a linked slice",
			args: args{
				in: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			}, //       0, 1, 2, 3, 4, 5, 6, 7, 8, 9
			want: []int{0, 2, 5, 8, 6, 4, 7, 3, 9, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, createCircleSlice(tt.args.in))
		})
	}
}
