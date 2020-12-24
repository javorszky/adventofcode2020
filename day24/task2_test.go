package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getWorldBlackWhite(t *testing.T) {
	type args struct {
		in map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "turns a world into colors",
			args: args{
				in: map[string]int{
					"1": 0,
					"2": 1,
					"3": 2,
					"4": 89,
					"5": 110,
				},
			},
			want: map[string]string{
				"1": white,
				"2": black,
				"3": white,
				"4": black,
				"5": white,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getWorldBlackWhite(tt.args.in))
		})
	}
}

func Test_getDirectionsFromCoordinate(t *testing.T) {
	type args struct {
		d directions
	}
	tests := []struct {
		name string
		args args
		want directions
	}{
		{
			name: "converts directions to coordinate, then back to directions",
			args: args{
				d: directions{
					se: 4,
					sw: 3,
					ne: 1,
					nw: 4,
					e:  2,
					w:  7,
				},
			},
			want: directions{
				se: 4,
				sw: 3,
				ne: 1,
				nw: 4,
				e:  2,
				w:  7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getDirectionsFromCoordinate(getCoordinate(tt.args.d)))
		})
	}
}

func Test_getAdjacentAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "gets neighbours for origin tile",
			args: args{
				s: "0.0.0.0.0.0",
			},
			want: []string{
				"1.0.0.0.0.0",
				"0.1.0.0.0.0",
				"0.0.1.0.0.0",
				"0.0.0.1.0.0",
				"0.0.0.0.1.0",
				"0.0.0.0.0.1",
			},
		},
		{
			name: "gets neighbours for a tile wsw to origin",
			args: args{
				s: "0.0.0.3.2.0",
			},
			want: []string{
				"0.0.0.2.2.0", // confirmed
				"0.0.0.3.1.0", // confirmed
				"0.0.0.4.1.0", // confirmed
				"0.0.0.4.2.0", // confirmed
				"0.0.0.3.3.0", // confirmed
				"0.0.0.2.3.0", // confirmed
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, getAdjacentAddresses(tt.args.s))
		})
	}
}

func Test_getExpandedWorldCoordinates(t *testing.T) {
	type args struct {
		in map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{
			name: "creates expanded coordinates from one tile",
			args: args{
				in: map[string]string{
					"0.0.0.0.0.0": white,
				},
			},
			want: map[string]struct{}{
				"0.0.0.0.0.0": {},
				"1.0.0.0.0.0": {},
				"0.1.0.0.0.0": {},
				"0.0.1.0.0.0": {},
				"0.0.0.1.0.0": {},
				"0.0.0.0.1.0": {},
				"0.0.0.0.0.1": {},
			},
		},
		{
			name: "creates expanded coordinates from two tile",
			args: args{
				in: map[string]string{
					"0.0.0.0.0.0": white,
					"1.0.0.0.0.0": white,
				},
			},
			want: map[string]struct{}{
				"0.0.0.0.0.0": {},
				"1.0.0.0.0.0": {},
				"2.0.0.0.0.0": {},
				"1.1.0.0.0.0": {},
				"1.0.0.0.0.1": {},
				"0.1.0.0.0.0": {},
				"0.0.1.0.0.0": {},
				"0.0.0.1.0.0": {},
				"0.0.0.0.1.0": {},
				"0.0.0.0.0.1": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getExpandedWorldCoordinates(tt.args.in))
		})
	}
}

func Test_flip(t *testing.T) {
	type args struct {
		in    map[string]string
		world map[string]struct{}
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "flips tiles correctly",
			args: args{
				in: map[string]string{
					"0.0.0.0.0.0": black,
					"0.0.0.1.0.0": black,
					"0.0.0.0.1.0": white,
				},
				world: map[string]struct{}{
					"0.0.0.0.0.0": {},
					"0.0.0.1.0.0": {},
					"0.0.0.0.1.0": {},
					"1.0.0.0.0.0": {},
					"0.1.0.0.0.0": {},
					"0.0.1.0.0.0": {},
					"0.0.0.0.0.1": {},
					"0.0.0.2.0.0": {},
					"0.0.0.1.1.0": {},
					"0.0.0.0.1.1": {},
					"0.0.0.0.2.0": {},
				},
			},
			want: map[string]string{
				"0.0.0.0.0.0": black,
				"1.0.0.0.0.0": white,
				"0.1.0.0.0.0": white,
				"0.0.1.0.0.0": black,
				"0.0.0.1.0.0": black,
				"0.0.0.0.1.0": black,
				"0.0.0.0.0.1": white,
				"0.0.0.2.0.0": white,
				"0.0.0.1.1.0": white,
				"0.0.0.0.2.0": white,
				"0.0.0.0.1.1": white,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, flip(tt.args.in, tt.args.world))
		})
	}
}

func Test_blackTilesAfterNFlips(t *testing.T) {
	tiles := getWorldBlackWhite(getWorld([]string{
		"sesenwnenenewseeswwswswwnenewsewsw",
		"neeenesenwnwwswnenewnwwsewnenwseswesw",
		"seswneswswsenwwnwse",
		"nwnwneseeswswnenewneswwnewseswneseene",
		"swweswneswnenwsewnwneneseenw",
		"eesenwseswswnenwswnwnwsewwnwsene",
		"sewnenenenesenwsewnenwwwse",
		"wenwwweseeeweswwwnwwe",
		"wsweesenenewnwwnwsenewsenwwsesesenwne",
		"neeswseenwwswnwswswnw",
		"nenwswwsewswnenenewsenwsenwnesesenew",
		"enewnwewneswsewnwswenweswnenwsenwsw",
		"sweneswneswneneenwnewenewwneswswnese",
		"swwesenesewenwneswnwwneseswwne",
		"enesenwswwswneneswsenwnewswseenwsese",
		"wnwnesenesenenwwnenwsewesewsesesew",
		"nenewswnwewswnenesenwnesewesw",
		"eneswnwswnwsenenwnwnwwseeswneewsenese",
		"neswnwewnwnwseenwseesewsenwsweewe",
		"wseweeenwnesenwwwswnew",
	}))

	type args struct {
		in map[string]string
		n  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "counts black tiles in example input after 1 round",
			args: args{
				in: tiles,
				n:  1,
			},
			want: 15,
		},
		{
			name: "counts black tiles in example input after 2 rounds",
			args: args{
				in: tiles,
				n:  2,
			},
			want: 12,
		},
		{
			name: "counts black tiles in example input after 10 rounds",
			args: args{
				in: tiles,
				n:  10,
			},
			want: 37,
		},
		{
			name: "counts black tiles in example input after 100 rounds",
			args: args{
				in: tiles,
				n:  100,
			},
			want: 2208,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, blackTilesAfterNFlips(tt.args.in, tt.args.n))
		})
	}
}
