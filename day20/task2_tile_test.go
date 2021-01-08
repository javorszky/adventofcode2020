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
				WithBorders: []string{
					"#.......##",
					"#..#.....#",
					"..#....#.#",
					"##...##.##",
					"#.##.#....",
					".#.......#",
					".#.......#",
					"#..#..#..#",
					".......#.#",
					".##.#.####",
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
				WithBorders: []string{
					"..##.#..#.",
					"##..#.....",
					"#...##..#.",
					"####.#...#",
					"##.##.###.",
					"##...#.###",
					".#.#.#..##",
					"..#....#..",
					"###...#.#.",
					"..###..###",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseTileTask2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTileTask2() = %v, row %v", got, tt.want)
			}
		})
	}
}

func Test_integrity(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		want        tilev2
		content     []string
		withBorders []string
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
				WithBorders: []string{
					"..##.#..#.",
					"##..#.....",
					"#...##..#.",
					"####.#...#",
					"##.##.###.",
					"##...#.###",
					".#.#.#..##",
					"..#....#..",
					"###...#.#.",
					"..###..###",
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
			withBorders: []string{
				"..###..###",
				"###...#.#.",
				"..#....#..",
				".#.#.#..##",
				"##...#.###",
				"##.##.###.",
				"####.#...#",
				"#...##..#.",
				"##..#.....",
				"..##.#..#.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tile := parseTileTask2(tt.args.s)
			assert.Equal(t, tt.want, tile)

			if !oneMatches(tt.content, tile.Content) {
				assert.FailNow(t, "no version of tile content matched what was expected")
			}

			if !oneMatches(tt.withBorders, tile.WithBorders) {
				assert.FailNow(t, "no version of tile with borders matched what was expected")
			}
		})
	}
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
				Top:    "123456",
				Right:  "6bhn",
				Bottom: "ijklmn",
				Left:   "17ci",
				Content: []string{
					"890a",
					"defg",
				},
				WithBorders: []string{
					"123456",
					"7890ab",
					"cdefgh",
					"ijklmn",
				},
			},
			want: tilev2{
				ID:     "2311001",
				Top:    "ic71",
				Right:  "123456",
				Bottom: "nhb6",
				Left:   "ijklmn",
				Content: []string{
					"d8",
					"e9",
					"f0",
					"ga",
				},
				WithBorders: []string{
					"ic71",
					"jd82",
					"ke93",
					"lf04",
					"mga5",
					"nhb6",
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
				Top:    "123456",
				Right:  "6bhn",
				Bottom: "ijklmn",
				Left:   "17ci",
				Content: []string{
					"890a",
					"defg",
				},
				WithBorders: []string{
					"123456",
					"7890ab",
					"cdefgh",
					"ijklmn",
				},
			},
			want: tilev2{
				ID:     "2311100",
				Top:    "ijklmn",
				Right:  "nhb6",
				Bottom: "123456",
				Left:   "ic71",
				Content: []string{
					"defg",
					"890a",
				},
				WithBorders: []string{
					"ijklmn",
					"cdefgh",
					"7890ab",
					"123456",
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
				Top:    "123456",
				Right:  "6bhn",
				Bottom: "ijklmn",
				Left:   "17ci",
				Content: []string{
					"890a",
					"defg",
				},
				WithBorders: []string{
					"123456",
					"7890ab",
					"cdefgh",
					"ijklmn",
				},
			},
			want: tilev2{
				ID:     "2311010",
				Top:    "654321",
				Right:  "17ci",
				Bottom: "nmlkji",
				Left:   "6bhn",
				Content: []string{
					"a098",
					"gfed",
				},
				WithBorders: []string{
					"654321",
					"ba0987",
					"hgfedc",
					"nmlkji",
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

func Test_Orientations(t *testing.T) {
	upright := tilev2{
		ID:     "2311000",
		Top:    "123456",
		Right:  "6bhn",
		Bottom: "ijklmn",
		Left:   "17ci",
		Content: []string{
			"890a",
			"defg",
		},
		WithBorders: []string{
			"123456",
			"7890ab",
			"cdefgh",
			"ijklmn",
		},
	}

	tests := []struct {
		name         string
		orientation1 tilev2
		orientation2 tilev2
	}{
		// upright and flipvfliph
		{
			name:         "rotate rotate is the same as flipv fliph",
			orientation1: upright.rotate().rotate(),
			orientation2: upright.flipH().flipV(),
		},
		{
			name:         "rotate rotate is the same as fliph flipv",
			orientation1: upright.rotate().rotate(),
			orientation2: upright.flipH().flipV(),
		},
		{
			name:         "rotate rotate rotate is the same as flipv fliph rotate",
			orientation1: upright.rotate().rotate().rotate(),
			orientation2: upright.flipH().flipV().rotate(),
		},
		{
			name:         "rotate rotate rotate is the same as fliph flipv rotate",
			orientation1: upright.rotate().rotate().rotate(),
			orientation2: upright.flipV().flipH().rotate(),
		},

		// flipV
		{
			name:         "flipv-rotate-rotate is the same as fliph",
			orientation1: upright.flipV().rotate().rotate(),
			orientation2: upright.flipH(),
		},
		{
			name:         "flipv rotate rotate rotate is the same as fliph rotate",
			orientation1: upright.flipV().rotate().rotate().rotate(),
			orientation2: upright.flipH().rotate(),
		},

		// flipH
		{
			name:         "fliph-rotate-rotate is the same as flipv",
			orientation1: upright.flipH().rotate().rotate(),
			orientation2: upright.flipV(),
		},
		{
			name:         "fliph-rotate-rotate-rotate is the same as flipv rotate",
			orientation1: upright.flipH().rotate().rotate().rotate(),
			orientation2: upright.flipV().rotate(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.orientation1, tt.orientation2)
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

func oneMatches(a, b []string) bool {
	if Equal(a, b) {
		return true
	}
	if Equal(a, rotateContent(b)) {
		return true
	}
	if Equal(a, hFlipContent(b)) {
		return true
	}
	if Equal(a, rotateContent(hFlipContent(b))) {
		return true
	}
	if Equal(a, vFlipContent(b)) {
		return true
	}
	if Equal(a, rotateContent(vFlipContent(b))) {
		return true
	}
	if Equal(a, hFlipContent(vFlipContent(b))) {
		return true
	}
	if Equal(a, rotateContent(hFlipContent(vFlipContent(b)))) {
		return true
	}
	return false
}
