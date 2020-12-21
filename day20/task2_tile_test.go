package day20

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vFlipContent(t *testing.T) {
	type args struct {
		c []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "correctly vertically flips string slice",
			args: args{
				c: []string{
					"1234",
					"5678",
					"90ab",
					"cdef",
				},
			},
			want: []string{
				"cdef",
				"90ab",
				"5678",
				"1234",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, vFlipContent(tt.args.c))
		})
	}
}

func Test_hFlipContent(t *testing.T) {
	type args struct {
		c []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "correctly horizontally flips string slice",
			args: args{
				c: []string{
					"1234",
					"5678",
					"90ab",
					"cdef",
				},
			},
			want: []string{
				"4321",
				"8765",
				"ba09",
				"fedc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hFlipContent(tt.args.c))
		})
	}
}

func Test_rotateContent(t *testing.T) {
	type args struct {
		c []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "correctly rotates string slice",
			args: args{
				c: []string{
					"1234",
					"5678",
					"90ab",
					"cdef",
				},
			},
			want: []string{
				"c951",
				"d062",
				"ea73",
				"fb84",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateContent(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTileTask2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want tilev2
	}{
		{
			name: "correctly parses tile image into tile struct",
			args: args{
				s: `Tile 3391:
#.......##
#..#.....#
..#....#.#
##...##.##
#.##.#....
.#.......#
.#.......#
#..#..#..#
.......#.#
.##.#.####`,
			},
			want: tilev2{
				ID:     "3391000",
				Top:    "#.......##",
				Right:  "####.#####",
				Bottom: ".##.#.####",
				Left:   "##.##..#..",
				Content: []string{
					"..#.....",
					".#....#.",
					"#...##.#",
					".##.#...",
					"#.......",
					"#.......",
					"..#..#..",
					"......#.",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTileTask2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTileTask2() = %v, want %v", got, tt.want)
			}
		})
	}
}
