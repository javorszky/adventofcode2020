package day23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createInputSlice(t *testing.T) {
	type args struct {
		start []int
		total int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "creates expanded slice",
			args: args{
				start: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
				total: 15,
			},
			want: []int{3, 8, 9, 1, 2, 5, 4, 6, 7, 10, 11, 12, 13, 14, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, createInputSlice(tt.args.start, tt.args.total))
		})
	}
}
