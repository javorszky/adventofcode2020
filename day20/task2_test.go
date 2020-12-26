package day20

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_replaceMonster(t *testing.T) {
	type args struct {
		image  string
		coords []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "replaces strings at specified coords",
			args: args{
				image:  "....................",
				coords: []int{0, 3, 4, 6, 7, 9},
			},
			want: "O..OO.OO.O..........",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, replaceMonster(tt.args.image, tt.args.coords))
		})
	}
}

func Test_flattenImage(t *testing.T) {
	type args struct {
		img []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{
			name: "flattens image correctly",
			args: args{
				img: []string{
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
			want:  ".#.#..#.##...#.##..########....#.#....#..#......##.##.###.#.#..######...###.#####...#.#####.#..###.#....#.##.####...#.##...########.#....#####.#....#..#...##..#.#.###...####...#..#.....#......#..#.##..#..###.#.##....#.####..#.####.#.#.###..###.#.#...#.######.#..###.####....##..########.###..##.#...#...#.#.#.#.....#..#..#.#.##..###.###.#.#....#.##.#...###.##.###.#...#..#.##.######...#.#.###.##.##.#..#.##...####.###.#...###.#..#.#..#.#..#..#.#.#.####.####..####...#.#.#.###.###.#####..#####...###....###.##..#..#...#..####...#.#.###..##..##..####.##....###...##...#...#..###",
			want1: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := flattenImage(tt.args.img)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_seaRoughness(t *testing.T) {
	type args struct {
		stitchedImage []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "correctly finds the roughness of the sea per the example on aoc day 20",
			args: args{
				stitchedImage: []string{
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
			want: 273,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, seaRoughness(tt.args.stitchedImage))
		})
	}
}

func Test_findMonsterCoords(t *testing.T) {
	type args struct {
		flattenedImage string
		lineLength     int
		seaMonster     seaMonster
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "finds the coords for the monster",
			args: args{
				flattenedImage: ".....................#....#....##....##....###....#..#..#..#..#..#...",
				lineLength:     23,
				seaMonster:     newSeaMonster(strings.Split(seaMonsterPattern, "\n"), 23),
			},
			want: [][]int{
				{21, 26, 31, 32, 37, 38, 43, 44, 45, 50, 53, 56, 59, 62, 65},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMonsterCoords(tt.args.flattenedImage, tt.args.lineLength, tt.args.seaMonster))
		})
	}
}
