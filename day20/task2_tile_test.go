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
			assert.Equal(t, tt.want, rotateContent(tt.args.c))
		})
	}
}

func Test_parseTileTask2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    tilev2
		content []string
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
		{
			name: "correctly parses tile image into tile struct 2",
			args: args{
				s: `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`,
			},
			want: tilev2{
				ID:     "2311000",
				Top:    "..##.#..#.",
				Right:  "...#.##..#",
				Bottom: "..###..###",
				Left:   ".#####..#.",
				Content: []string{
					"#..#....",
					"...##..#",
					"###.#...",
					"#.##.###",
					"#...#.##",
					"#.#.#..#",
					".#....#.",
					"##...#.#",
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

func Test_integrity(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    tilev2
		content []string
	}{
		{
			name: "correctly parses tile image into tile struct 2",
			args: args{
				s: `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`,
			},
			want: tilev2{
				ID:     "2311000",
				Top:    "..##.#..#.",
				Right:  "...#.##..#",
				Bottom: "..###..###",
				Left:   ".#####..#.",
				Content: []string{
					"#..#....",
					"...##..#",
					"###.#...",
					"#.##.###",
					"#...#.##",
					"#.#.#..#",
					".#....#.",
					"##...#.#",
				},
			},
			content: []string{
				"##...#.#",
				".#....#.",
				"#.#.#..#",
				"#...#.##",
				"#.##.###",
				"###.#...",
				"...##..#",
				"#..#....",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tile := parseTileTask2(tt.args.s)
			assert.Equal(t, tt.want, tile)

			if Equal(tt.content, tile.Content) {
				return
			}
			if Equal(tt.content, rotateContent(tile.Content)) {
				return
			}
			if Equal(tt.content, hFlipContent(tile.Content)) {
				return
			}
			if Equal(tt.content, rotateContent(hFlipContent(tile.Content))) {
				return
			}
			if Equal(tt.content, vFlipContent(tile.Content)) {
				return
			}
			if Equal(tt.content, rotateContent(vFlipContent(tile.Content))) {
				return
			}
			if Equal(tt.content, hFlipContent(vFlipContent(tile.Content))) {
				return
			}
			if Equal(tt.content, rotateContent(hFlipContent(vFlipContent(tile.Content)))) {
				return
			}
			assert.FailNow(t, "no version of tile content matched what was expected")
		})
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Test_tilev2_rotate(t *testing.T) {
	tests := []struct {
		name string
		tile tilev2
		want tilev2
	}{
		{
			name: "correctly rotates tile",
			tile: tilev2{
				ID:     "2311000",
				Top:    "1234",
				Right:  "48bf",
				Bottom: "cdef",
				Left:   "159c",
				Content: []string{
					"67",
					"0a",
				},
			},
			want: tilev2{
				ID:     "2311001",
				Top:    "c951",
				Right:  "1234",
				Bottom: "fb84",
				Left:   "cdef",
				Content: []string{
					"06",
					"a7",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.tile.rotate())
		})
	}
}

func Test_tilev2_flipV(t *testing.T) {
	tests := []struct {
		name string
		tile tilev2
		want tilev2
	}{
		{
			name: "correctly flips vertically tile",
			tile: tilev2{
				ID:     "2311000",
				Top:    "1234",
				Right:  "48bf",
				Bottom: "cdef",
				Left:   "159c",
				Content: []string{
					"67",
					"0a",
				},
			},
			want: tilev2{
				ID:     "2311100",
				Top:    "cdef",
				Right:  "fb84",
				Bottom: "1234",
				Left:   "c951",
				Content: []string{
					"0a",
					"67",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.tile.flipV())
		})
	}
}

func Test_tilev2_flipH(t *testing.T) {
	tests := []struct {
		name string
		tile tilev2
		want tilev2
	}{
		{
			name: "correctly flips horizontally tile",
			tile: tilev2{
				ID:     "2311000",
				Top:    "1234",
				Right:  "48bf",
				Bottom: "cdef",
				Left:   "159c",
				Content: []string{
					"67",
					"0a",
				},
			},
			want: tilev2{
				ID:     "2311010",
				Top:    "4321",
				Right:  "159c",
				Bottom: "fedc",
				Left:   "48bf",
				Content: []string{
					"76",
					"a0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			assert.Equal(t1, tt.want, tt.tile.flipH())
		})
	}
}
