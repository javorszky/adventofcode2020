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
		{
			name: "returns the ninth number correctly",
			args: args{
				in: []int{14, 1, 17, 0, 3, 20},
				n:  9,
			},
			want: 3,
		},
		{
			name: "returns the tenth number correctly",
			args: args{
				in: []int{14, 1, 17, 0, 3, 20},
				n:  10,
			},
			want: 1,
		},
		{
			name: "returns the eleventh number correctly",
			args: args{
				in: []int{14, 1, 17, 0, 3, 20},
				n:  11,
			},
			want: 8,
		},
		{
			name: "aoc explanation 1 - 1",
			args: args{
				in: []int{0, 3, 6},
				n:  1,
			},
			want: 0,
		},
		{
			name: "aoc explanation 1 - 2",
			args: args{
				in: []int{0, 3, 6},
				n:  2,
			},
			want: 3,
		},
		{
			name: "aoc explanation 1 - 3",
			args: args{
				in: []int{0, 3, 6},
				n:  3,
			},
			want: 6,
		},
		{
			name: "aoc explanation 1 - 4",
			args: args{
				in: []int{0, 3, 6},
				n:  4,
			},
			want: 0,
		},
		{
			name: "aoc explanation 1 - 5",
			args: args{
				in: []int{0, 3, 6},
				n:  5,
			},
			want: 3,
		},
		{
			name: "aoc explanation 1 - 6",
			args: args{
				in: []int{0, 3, 6},
				n:  6,
			},
			want: 3,
		},
		{
			name: "aoc explanation 1 - 7",
			args: args{
				in: []int{0, 3, 6},
				n:  7,
			},
			want: 1,
		},
		{
			name: "aoc explanation 1 - 8",
			args: args{
				in: []int{0, 3, 6},
				n:  8,
			},
			want: 0,
		},
		{
			name: "aoc explanation 1 - 9",
			args: args{
				in: []int{0, 3, 6},
				n:  9,
			},
			want: 4,
		},
		{
			name: "aoc explanation 1 - 10",
			args: args{
				in: []int{0, 3, 6},
				n:  10,
			},
			want: 0,
		},
		{
			name: "aoc explanation 1 - 2020",
			args: args{
				in: []int{0, 3, 6},
				n:  2020,
			},
			want: 436,
		},
		{
			name: "aoc test 1",
			args: args{
				in: []int{1, 3, 2},
				n:  2020,
			},
			want: 1,
		},
		{
			name: "aoc test 2",
			args: args{
				in: []int{2, 1, 3},
				n:  2020,
			},
			want: 10,
		},
		{
			name: "aoc test 3",
			args: args{
				in: []int{1, 2, 3},
				n:  2020,
			},
			want: 27,
		},
		{
			name: "aoc test 4",
			args: args{
				in: []int{2, 3, 1},
				n:  2020,
			},
			want: 78,
		},
		{
			name: "aoc test 5",
			args: args{
				in: []int{3, 2, 1},
				n:  2020,
			},
			want: 438,
		},
		{
			name: "aoc test 6",
			args: args{
				in: []int{3, 1, 2},
				n:  2020,
			},
			want: 1836,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, nthNumber(tt.args.in, tt.args.n))
		})
	}
}
