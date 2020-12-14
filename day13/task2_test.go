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
		in map[int][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 1",
			args: args{
				in: map[int][]int{
					0: {17},
					2: {13},
					3: {19},
				},
			},
			want: 3417,
		},
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 2",
			args: args{
				in: map[int][]int{
					0: {67},
					1: {7},
					2: {59},
					3: {61},
				},
			},
			want: 754018,
		},
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 3",
			args: args{
				in: map[int][]int{
					0: {67},
					2: {7},
					3: {59},
					4: {61},
				},
			},
			want: 779210,
		},
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 4",
			args: args{
				in: map[int][]int{
					0: {67},
					1: {7},
					3: {59},
					4: {61},
				},
			},
			want: 1261476,
		},
		{
			name: "returns earliest timestamp where consecutive departures happen, aoc test 5",
			args: args{
				in: map[int][]int{
					0: {1789},
					1: {37},
					2: {47},
					3: {1889},
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

func Test_togetherBus(t *testing.T) {
	type args struct {
		bus1Start  int
		bus1ID     int
		bus1offset int
		bus2Start  int
		bus2ID     int
		bus2offset int
	}
	tests := []struct {
		name       string
		args       args
		wantStart  int
		wantID     int
		wantOffset int
	}{
		{
			name: "buses 0/17/0 and 0/13/2 will give 102/221/0",
			args: args{
				bus1Start:  0,
				bus1ID:     17,
				bus1offset: 0,
				bus2Start:  0,
				bus2ID:     13,
				bus2offset: 2,
			},
			wantStart:  102,
			wantID:     221,
			wantOffset: 0,
		},
		{
			name: "buses 0/13/2 and 0/17/0 will give 102/221/0",
			args: args{
				bus1Start:  0,
				bus1ID:     13,
				bus1offset: 2,
				bus2Start:  0,
				bus2ID:     17,
				bus2offset: 0,
			},
			wantStart:  102,
			wantID:     221,
			wantOffset: 0,
		},
		//{
		//	name: "buses 8/4/1 and 5/3/2 will give 9/12/0",
		//	args: args{
		//		bus1Start:  8,
		//		bus1ID:     4,
		//		bus1offset: 1,
		//		bus2Start:  5,
		//		bus2ID:     3,
		//		bus2offset: 2,
		//	},
		//	wantStart:  9,
		//	wantID:     12,
		//	wantOffset: 0,
		//},
		//{
		//	name: "buses 0/4/0 and 0/3/1 will give 8/12/0",
		//	args: args{
		//		bus1Start:  0,
		//		bus1ID:     4,
		//		bus1offset: 0,
		//		bus2Start:  0,
		//		bus2ID:     3,
		//		bus2offset: 1,
		//	},
		//	wantStart:  8,
		//	wantID:     12,
		//	wantOffset: 0,
		//},
		//{
		//	name: "buses 0/4/1 and 0/3/0 will give 13/12/0",
		//	args: args{
		//		bus1Start:  0,
		//		bus1ID:     4,
		//		bus1offset: 1,
		//		bus2Start:  0,
		//		bus2ID:     3,
		//		bus2offset: 0,
		//	},
		//	wantStart:  9,
		//	wantID:     12,
		//	wantOffset: 0,
		//},
		//{
		//	name: "given two buses, returns a new bus that both buses will be encompassed by",
		//	args: args{
		//		bus1Start:  0,
		//		bus1ID:     17,
		//		bus1offset: 0,
		//		bus2Start:  0,
		//		bus2ID:     19,
		//		bus2offset: 2,
		//	},
		//	wantStart:  17,
		//	wantID:     323,
		//	wantOffset: 0,
		//},
		//{
		//	name: "given two buses, returns a new bus that both buses will be encompassed by",
		//	args: args{
		//		bus1Start:  0,
		//		bus1ID:     19,
		//		bus1offset: 0,
		//		bus2Start:  0,
		//		bus2ID:     17,
		//		bus2offset: 2,
		//	},
		//	wantStart:  304,
		//	wantID:     323,
		//	wantOffset: 0,
		//},
		//{
		//	name: "given two buses, returns a new bus that both buses will be encompassed by 2",
		//	args: args{
		//		bus1ID:     23,
		//		bus1offset: 4,
		//		bus2ID:     41,
		//		bus2offset: 9,
		//	},
		//	want:  943,
		//	want1: 50,
		//},
		//{
		//	name: "given two buses, where one is a 1,0, returns bus 2",
		//	args: args{
		//		bus1ID:     1,
		//		bus1offset: 0,
		//		bus2ID:     41,
		//		bus2offset: 9,
		//	},
		//	want:  41,
		//	want1: 9,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := togetherBus(tt.args.bus1Start, tt.args.bus1ID, tt.args.bus1offset, tt.args.bus2Start, tt.args.bus2ID, tt.args.bus2offset)
			assert.Equal(t, tt.wantStart, got)
			assert.Equal(t, tt.wantID, got1)
			assert.Equal(t, tt.wantOffset, got2)
		})
	}
}
