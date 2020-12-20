package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseTile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want tile
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
			want: tile{
				ID:     "3391000",
				Top:    "#.......##",
				Right:  "####.#####",
				Bottom: ".##.#.####",
				Left:   "##.##..#..",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseTile(tt.args.s))
		})
	}
}

func Test_reverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverses string",
			args: args{
				s: "a string",
			},
			want: "gnirts a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseString(tt.args.s))
		})
	}
}

func Test_tile_flipV(t1 *testing.T) {
	tests := []struct {
		name string
		tile tile
		want tile
	}{
		{
			name: "flips tile vertically, changes ID",
			tile: tile{
				ID:     "3391000",
				Top:    "#.......##",
				Right:  "####.#####",
				Bottom: ".##.#.####",
				Left:   "##.##..#..",
			},
			want: tile{
				ID:     "3391100",
				Top:    ".##.#.####",
				Right:  "#####.####",
				Bottom: "#.......##",
				Left:   "..#..##.##",
			},
		},
		{
			name: "flips tile vertically, changes ID",
			tile: tile{
				ID:     "3391100",
				Top:    "#.......##",
				Right:  "####.#####",
				Bottom: ".##.#.####",
				Left:   "##.##..#..",
			},
			want: tile{
				ID:     "3391000",
				Top:    ".##.#.####",
				Right:  "#####.####",
				Bottom: "#.......##",
				Left:   "..#..##.##",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.tile.flipV())
		})
	}
}

func Test_tile_flipH(t1 *testing.T) {
	tests := []struct {
		name string
		tile tile
		want tile
	}{
		{
			name: "flips tile vertically, changes ID",
			tile: tile{
				ID:     "3391000",
				Top:    "#.......##",
				Right:  "####.#####",
				Bottom: ".##.#.####",
				Left:   "##.##..#..",
			},
			want: tile{
				ID:     "3391010",
				Top:    "##.......#",
				Right:  "##.##..#..",
				Bottom: "####.#.##.",
				Left:   "####.#####",
			},
		},
		{
			name: "flips tile vertically, changes ID",
			tile: tile{
				ID:     "3391010",
				Top:    "#.......##",
				Right:  "####.#####",
				Bottom: ".##.#.####",
				Left:   "##.##..#..",
			},
			want: tile{
				ID:     "3391000",
				Top:    "##.......#",
				Right:  "##.##..#..",
				Bottom: "####.#.##.",
				Left:   "####.#####",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.tile.flipH())
		})
	}
}

func Test_tile_rotate(t1 *testing.T) {
	tests := []struct {
		name string
		tile tile
		want tile
	}{
		{
			name: "rotates 000 tile once",
			tile: tile{
				ID:     "3391000",
				Top:    "##.......#",
				Right:  "##.##..#..",
				Bottom: "####.#.##.",
				Left:   "####.#####",
			},
			want: tile{
				ID:     "3391001",
				Right:  "##.......#",
				Bottom: "##.##..#..",
				Left:   "####.#.##.",
				Top:    "####.#####",
			},
		},
		{
			name: "rotates 001 tile once to become a 110 tile",
			tile: tile{
				ID:     "3391001",
				Top:    "##.......#",
				Right:  "##.##..#..",
				Bottom: "####.#.##.",
				Left:   "####.#####",
			},
			want: tile{
				ID:     "3391110",
				Right:  "##.......#",
				Bottom: "##.##..#..",
				Left:   "####.#.##.",
				Top:    "####.#####",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.tile.rotate())
		})
	}
}
