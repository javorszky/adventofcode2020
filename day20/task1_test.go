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
				ID:     3391,
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
			name: "flips tile vertically",
			tile: tile{
				ID:     3391,
				Top:    "#.......##",
				Right:  "####.#####",
				Bottom: ".##.#.####",
				Left:   "##.##..#..",
			},
			want: tile{
				ID:     3391,
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
			name: "flips tile vertically",
			tile: tile{
				ID:     3391,
				Top:    "#.......##",
				Right:  "####.#####",
				Bottom: ".##.#.####",
				Left:   "##.##..#..",
			},
			want: tile{
				ID:     3391,
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
