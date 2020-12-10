package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countBranch(t *testing.T) {
	type args struct {
		max        int
		joltageMap map[int]struct{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "counts the possible paths for a small map",
			args: args{
				max: 3,
				joltageMap: map[int]struct{}{
					0: {},
					1: {},
					2: {},
					3: {},
				},
			},
			want: 4,
		},
		{
			name: "counts the possible paths for a larger map",
			args: args{
				max: 4,
				joltageMap: map[int]struct{}{
					0: {},
					1: {},
					2: {},
					3: {},
					4: {},
				},
			},
			want: 7,
		},
		{
			name: "counts the possible paths for a larger map",
			args: args{
				max: 5,
				joltageMap: map[int]struct{}{
					0: {},
					1: {},
					2: {},
					3: {},
					4: {},
					5: {},
				},
			},
			want: 13,
		},
		{
			name: "counts the possible paths for a larger map",
			args: args{
				max: 6,
				joltageMap: map[int]struct{}{
					0: {},
					1: {},
					2: {},
					3: {},
					4: {},
					5: {},
					6: {},
				},
			},
			want: 24,
		},
		{
			name: "counts the possible paths for the larger example map",
			args: args{
				max: 52,
				joltageMap: map[int]struct{}{
					28: {},
					33: {},
					18: {},
					42: {},
					31: {},
					14: {},
					46: {},
					20: {},
					48: {},
					47: {},
					24: {},
					23: {},
					49: {},
					45: {},
					19: {},
					38: {},
					39: {},
					11: {},
					1:  {},
					32: {},
					25: {},
					35: {},
					8:  {},
					17: {},
					7:  {},
					9:  {},
					4:  {},
					2:  {},
					34: {},
					10: {},
					3:  {},
					52: {},
				},
			},
			want: 19208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countBranch(0, tt.args.max, tt.args.joltageMap))
		})
	}
}

func Test_countBranchEasier(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "counts the possible paths for a small map",
			args: args{
				input: []int{0, 1, 2, 3},
			},
			want: 4,
		},
		{
			name: "counts the possible paths for a small map 2",
			args: args{
				input: []int{0, 1, 2, 3, 4},
			},
			want: 7,
		},
		{
			name: "counts the possible paths for a small map 3 where one is missing",
			args: args{
				input: []int{0, 1, 2, 4},
			},
			want: 3,
		},
		{
			name: "counts the possible paths for a larger map",
			args: args{
				input: []int{0, 1, 2, 3, 4, 5},
			},
			want: 13,
		},
		{
			name: "counts the possible paths for a larger map where one is missing",
			args: args{
				input: []int{0, 1, 2, 4, 5},
			},
			want: 5,
		},
		{
			name: "counts the possible paths for a larger map 2",
			args: args{
				input: []int{0, 1, 2, 3, 4, 5, 6},
			},
			want: 24,
		},
		{
			name: "counts the possible paths for a larger map 2 where a single space is missing",
			args: args{
				input: []int{0, 1, 2, 4, 5, 6},
			},
			want: 8,
		},
		{
			name: "counts the possible paths for a larger map 2 where a double space is missing",
			args: args{
				input: []int{0, 1, 2, 5, 6},
			},
			want: 2,
		},
		{
			name: "counts the possible paths for a larger map 2 where two single spaces are missing",
			args: args{
				input: []int{0, 2, 3, 4, 6},
			},
			want: 5,
		},
		{
			name: "counts the possible paths for small example",
			args: args{
				input: []int{0, 1, 2, 3, 4, 7, 8, 9, 10, 11, 14, 17, 18, 19, 20, 23, 24, 25, 28, 31, 32, 33, 34, 35, 38, 39, 42, 45, 46, 47, 48, 49, 52},
			},
			want: 19208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countBranchEasier(tt.args.input))
		})
	}
}
