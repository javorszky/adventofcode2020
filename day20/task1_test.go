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
				ID:    3391,
				Side1: "#.......##",
				Side2: "####.#####",
				Side3: ".##.#.####",
				Side4: "##.##..#..",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseTile(tt.args.s))
		})
	}
}
