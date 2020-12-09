package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_t1checkNumber(t *testing.T) {
	type args struct {
		previous []int
		num      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successfully returns true when the number is a sum of any two different in the passed list",
			args: args{
				previous: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
				num:      25,
			},
			want: true,
		},
		{
			name: "successfully returns false when the number is not a sum of any two different in the passed list",
			args: args{
				previous: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
				num:      48,
			},
			want: false,
		},
		{
			name: "successfully returns false when the number is only a sum of two of the same numbers in the passed list",
			args: args{
				previous: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 23},
				num:      46,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, t1checkNumber(tt.args.previous, tt.args.num))
		})
	}
}
