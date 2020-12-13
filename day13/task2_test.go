package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_t2formatInput(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "turns input string into map correctly",
			args: args{
				in: []string{
					"192837918237",
					"7,x,3,x,x,x,9,10,42,x,x,6,x",
				},
			},
			want: map[int]int{
				0:  7,
				2:  3,
				6:  9,
				7:  10,
				8:  42,
				11: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, t2formatInput(tt.args.in))
		})
	}
}

func Test_earliestTime(t *testing.T) {
	type args struct {
		in map[int]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 1",
			args: args{
				in: map[int]int{
					0: 17,
					2: 13,
					3: 19,
				},
			},
			want: 3417,
		},
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 2",
			args: args{
				in: map[int]int{
					0: 67,
					1: 7,
					2: 59,
					3: 61,
				},
			},
			want: 754018,
		},
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 3",
			args: args{
				in: map[int]int{
					0: 67,
					2: 7,
					3: 59,
					4: 61,
				},
			},
			want: 779210,
		},
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 4",
			args: args{
				in: map[int]int{
					0: 67,
					1: 7,
					3: 59,
					4: 61,
				},
			},
			want: 1261476,
		},
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 4",
			args: args{
				in: map[int]int{
					0: 1789,
					1: 37,
					3: 47,
					4: 1889,
				},
			},
			want: 1202161486,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, earliestTime(tt.args.in))
		})
	}
}
