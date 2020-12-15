package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nthNumber(t *testing.T) {
	type args struct {
		in []int
		n  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns the first number correctly",
			args: args{
				in: []int{14, 1, 17, 0, 3, 20},
				n:  1,
			},
			want: 14,
		},
		{
			name: "returns the second number correctly",
			args: args{
				in: []int{14, 1, 17, 0, 3, 20},
				n:  2,
			},
			want: 1,
		},
		{
			name: "returns the sixth number correctly",
			args: args{
				in: []int{14, 1, 17, 0, 3, 20},
				n:  6,
			},
			want: 20,
		},
		{
			name: "returns the seventh number correctly",
			args: args{
				in: []int{14, 1, 17, 0, 3, 20},
				n:  7,
			},
			want: 0,
		},
		{
			name: "returns the eight number correctly",
			args: args{
				in: []int{14, 1, 17, 0, 3, 20},
				n:  8,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nthNumber(tt.args.in, tt.args.n))
		})
	}
}
