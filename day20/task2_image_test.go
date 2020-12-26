package day20

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stitchImage(t *testing.T) {
	type args struct {
		img image2
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "stitches image correctly",
			args: args{
				img: image2{
					W: 3,
					H: 3,
					Tiles: map[int]map[int]tilev2{
						0: {
							0: tilev2{
								Content: []string{
									"123",
									"456",
									"789",
								},
							},
							1: tilev2{
								Content: []string{
									"0ab",
									"cde",
									"fgh",
								},
							},
							2: tilev2{
								Content: []string{
									"ijk",
									"lmn",
									"opq",
								},
							},
						},
						1: {
							0: tilev2{
								Content: []string{
									"rst",
									"uvw",
									"xyz",
								},
							},
							1: tilev2{
								Content: []string{
									"321",
									"654",
									"987",
								},
							},
							2: tilev2{
								Content: []string{
									"ba0",
									"edc",
									"hgf",
								},
							},
						},
						2: {
							0: tilev2{
								Content: []string{
									"kji",
									"nml",
									"qpo",
								},
							},
							1: tilev2{
								Content: []string{
									"tsr",
									"wvu",
									"zyx",
								},
							},
							2: tilev2{
								Content: []string{
									"ABC",
									"DEF",
									"GHI",
								},
							},
						},
					},
				},
			},
			want: []string{
				"1230abijk",
				"456cdelmn",
				"789fghopq",
				"rst321ba0",
				"uvw654edc",
				"xyz987hgf",
				"kjitsrABC",
				"nmlwvuDEF",
				"qpozyxGHI",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				if !assert.Equal(t, tt.want, stitchImage(tt.args.img)) {
					assert.FailNow(t, "failed at iteration %d\n", i)
				}
			}
		})
	}
}

func Test_stitchImage_complete(t *testing.T) {
	type args struct {
		fileContent string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "parses example input, and returns an image that matches the example image",
			args: args{
				fileContent: `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`,
			},
			want: []string{
				".#.#..#.##...#.##..#####",
				"###....#.#....#..#......",
				"##.##.###.#.#..######...",
				"###.#####...#.#####.#..#",
				"##.#....#.##.####...#.##",
				"...########.#....#####.#",
				"....#..#...##..#.#.###..",
				".####...#..#.....#......",
				"#..#.##..#..###.#.##....",
				"#.####..#.####.#.#.###..",
				"###.#.#...#.######.#..##",
				"#.####....##..########.#",
				"##..##.#...#...#.#.#.#..",
				"...#..#..#.#.##..###.###",
				".#.#....#.##.#...###.##.",
				"###.#...#..#.##.######..",
				".#.#.###.##.##.#..#.##..",
				".####.###.#...###.#..#.#",
				"..#.#..#..#.#.#.####.###",
				"#..####...#.#.#.###.###.",
				"#####..#####...###....##",
				"#.##..#..#...#..####...#",
				".#.###..##..##..####.##.",
				"...###...##...#...#..###",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tss := make(tileSets2, 0)
			// there are 144 tiles, so that's a 12x12 grid.
			tiles := parseInput(tt.args.fileContent)
			for _, tileString := range tiles {
				tss.addTileSet(newTileSet2(parseTileTask2(tileString)))
			}

			s := int(math.Sqrt(float64(len(tiles))))

			i := image2{
				W:     s,
				H:     s,
				Tiles: make(map[int]map[int]tilev2, 0),
			}

			img, err := fitTileIntoImage2(i, 1, tss)

			if err != nil {
				assert.FailNow(t, "could not fit tiles into an image")
			}

			assert.True(t, oneMatches(tt.want, stitchImage(img)))
		})
	}
}

func Test_getXY2(t *testing.T) {
	type args struct {
		i image2
		n int
	}
	tests := []struct {
		name    string
		args    args
		row     int
		col     int
		wantErr bool
	}{
		{
			name: "finds xy coordinate correctly for tile 1 (top left corner in a 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 1,
			},
			row:     0,
			col:     0,
			wantErr: false,
		},
		{
			name: "finds xy coordinate correctly for tile 2 (top mid in a 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 2,
			},
			row:     0,
			col:     1,
			wantErr: false,
		},
		{
			name: "finds xy coordinate correctly for tile 3 (top right corner in a 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 3,
			},
			row:     0,
			col:     2,
			wantErr: false,
		},
		{
			name: "finds xy coordinate correctly for tile 4 (mid left edge in 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 4,
			},
			row:     1,
			col:     0,
			wantErr: false,
		},
		{
			name: "finds xy coordinate correctly for tile 5 (mid mid in 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 5,
			},
			row:     1,
			col:     1,
			wantErr: false,
		},
		{
			name: "finds xy coordinate correctly for tile 6 (mid right in 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 6,
			},
			row:     1,
			col:     2,
			wantErr: false,
		},
		{
			name: "finds xy coordinate correctly for tile 7 (bottom left in 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 7,
			},
			row:     2,
			col:     0,
			wantErr: false,
		},
		{
			name: "finds xy coordinate correctly for tile 8 (bottom mid in 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 8,
			},
			row:     2,
			col:     1,
			wantErr: false,
		},
		{
			name: "finds xy coordinate correctly for tile 9 (bottom right in 3x3 grid)",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 9,
			},
			row:     2,
			col:     2,
			wantErr: false,
		},
		{
			name: "errors for out of bound number below",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 0,
			},
			row:     0,
			col:     0,
			wantErr: true,
		},
		{
			name: "error for out of bounds above",
			args: args{
				i: image2{
					W: 3,
					H: 3,
				},
				n: 10,
			},
			row:     0,
			col:     0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getXY2(tt.args.i, tt.args.n)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equalf(t, tt.row, got, "x (got) coordinate is different")
			assert.Equalf(t, tt.col, got1, "y (got1) coordinate is different")
		})
	}
}
