package day24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseDirections(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want directions
	}{
		{
			name: "correctly parses direction string where all directions are present",
			args: args{
				s: "wwswesesesesewnwewwswnwwnenwnwsww",
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
		{
			name: "correctly parses direction string where only one direction is present",
			args: args{
				s: "wwww",
			},
			want: directions{
				se: 0,
				sw: 0,
				ne: 0,
				nw: 0,
				e:  0,
				w:  4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseDirections(tt.args.s))
		})
	}
}

func Test_collapseDirections(t *testing.T) {
	type args struct {
		d directions
	}
	tests := []struct {
		name string
		args args
		want directions
	}{
		{
			name: "resolves se / nw",
			args: args{
				d: directions{
					se: 1,
					sw: 0,
					ne: 0,
					nw: 1,
					e:  0,
					w:  0,
				},
			},
			want: directions{
				se: 0,
				nw: 0,
				sw: 0,
				ne: 0,
				e:  0,
				w:  0,
			},
		},
		{
			name: "resolves sw / ne",
			args: args{
				d: directions{
					se: 0,
					sw: 1,
					ne: 1,
					nw: 0,
					e:  0,
					w:  0,
				},
			},
			want: directions{
				se: 0,
				nw: 0,
				sw: 0,
				ne: 0,
				e:  0,
				w:  0,
			},
		},
		{
			name: "resolves e / w",
			args: args{
				d: directions{
					se: 0,
					sw: 0,
					ne: 0,
					nw: 0,
					e:  1,
					w:  1,
				},
			},
			want: directions{
				se: 0,
				nw: 0,
				sw: 0,
				ne: 0,
				e:  0,
				w:  0,
			},
		},
		{
			name: "resolves se / w",
			args: args{
				d: directions{
					se: 1,
					sw: 0,
					ne: 0,
					nw: 0,
					e:  0,
					w:  1,
				},
			},
			want: directions{
				se: 0,
				nw: 0,
				sw: 1,
				ne: 0,
				e:  0,
				w:  0,
			},
		},
		{
			name: "resolves sw / nw",
			args: args{
				d: directions{
					se: 0,
					sw: 1,
					ne: 0,
					nw: 1,
					e:  0,
					w:  0,
				},
			},
			want: directions{
				se: 0,
				nw: 0,
				sw: 0,
				ne: 0,
				e:  0,
				w:  1,
			},
		},
		{
			name: "resolves w / ne",
			args: args{
				d: directions{
					se: 0,
					sw: 0,
					ne: 1,
					nw: 0,
					e:  0,
					w:  1,
				},
			},
			want: directions{
				se: 0,
				nw: 1,
				sw: 0,
				ne: 0,
				e:  0,
				w:  0,
			},
		},
		{
			name: "resolves nw / e",
			args: args{
				d: directions{
					se: 0,
					sw: 0,
					ne: 0,
					nw: 1,
					e:  1,
					w:  0,
				},
			},
			want: directions{
				se: 0,
				nw: 0,
				sw: 0,
				ne: 1,
				e:  0,
				w:  0,
			},
		},
		{
			name: "resolves ne / se",
			args: args{
				d: directions{
					se: 1,
					sw: 0,
					ne: 1,
					nw: 0,
					e:  0,
					w:  0,
				},
			},
			want: directions{
				se: 0,
				nw: 0,
				sw: 0,
				ne: 0,
				e:  1,
				w:  0,
			},
		},
		{
			name: "resolves e / sw",
			args: args{
				d: directions{
					se: 0,
					sw: 1,
					ne: 0,
					nw: 0,
					e:  1,
					w:  0,
				},
			},
			want: directions{
				se: 1,
				nw: 0,
				sw: 0,
				ne: 0,
				e:  0,
				w:  0,
			},
		},
		{
			name: "multiple levels of collapse",
			args: args{
				d: directions{
					se: 4,
					sw: 4,
					ne: 5,
					nw: 1,
					e:  1,
					w:  5,
				},
			},
			want: directions{
				se: 0,
				nw: 0,
				sw: 2,
				ne: 0,
				e:  0,
				w:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, collapseDirections(tt.args.d))
		})
	}
}

func Test_getCoordinate(t *testing.T) {
	type args struct {
		d directions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "gets coordinate for uncollapsed direction",
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
			want: "1.2.4.3.7.4",
		},
		{
			name: "gets coordinate for collapsed direction",
			args: args{
				d: directions{
					se: 0,
					nw: 0,
					sw: 2,
					ne: 0,
					e:  0,
					w:  5,
				},
			},
			want: "0.0.0.2.5.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getCoordinate(tt.args.d))
		})
	}
}

func Test_countBlackTiles(t *testing.T) {
	type args struct {
		tiles []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "counts black tiles for given input - aoc d24 t1 example",
			args: args{
				tiles: []string{
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
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countBlackTiles(tt.args.tiles))
		})
	}
}

func Test_absDiff(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abs diff when a is larger",
			args: args{
				a: 8,
				b: 7,
			},
			want: 1,
		},
		{
			name: "abs diff when b is larger",
			args: args{
				a: 5,
				b: 7,
			},
			want: 2,
		},
		{
			name: "abs diff when a==b",
			args: args{
				a: 5,
				b: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, absDiff(tt.args.a, tt.args.b))
		})
	}
}

func Test_smoler(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a is smaller",
			args: args{
				a: 3,
				b: 9,
			},
			want: 3,
		},
		{
			name: "b is smaller",
			args: args{
				a: 9,
				b: 3,
			},
			want: 3,
		},
		{
			name: "a and b are eq",
			args: args{
				a: 5,
				b: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smoler(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("smoler() = %v, want %v", got, tt.want)
			}
		})
	}
}
