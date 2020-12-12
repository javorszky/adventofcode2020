package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newDirection(t *testing.T) {
	type args struct {
		base string
		dir  string
		amnt int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "east plus 90 produces south",
			args: args{
				base: east,
				dir:  "R",
				amnt: 90,
			},
			want: south,
		},
		{
			name: "east plus 180 produces west",
			args: args{
				base: east,
				dir:  "R",
				amnt: 180,
			},
			want: west,
		},
		{
			name: "east plus 270 produces north",
			args: args{
				base: east,
				dir:  "R",
				amnt: 270,
			},
			want: north,
		},
		{
			name: "south plus 90 produces west",
			args: args{
				base: south,
				dir:  "R",
				amnt: 90,
			},
			want: west,
		},
		{
			name: "south plus 180 produces north",
			args: args{
				base: south,
				dir:  "R",
				amnt: 180,
			},
			want: north,
		},
		{
			name: "south plus 270 produces east",
			args: args{
				base: south,
				dir:  "R",
				amnt: 270,
			},
			want: east,
		},
		{
			name: "west plus 90 produces north",
			args: args{
				base: west,
				dir:  "R",
				amnt: 90,
			},
			want: north,
		},
		{
			name: "west plus 180 produces east",
			args: args{
				base: west,
				dir:  "R",
				amnt: 180,
			},
			want: east,
		},
		{
			name: "west plus 270 produces south",
			args: args{
				base: west,
				dir:  "R",
				amnt: 270,
			},
			want: south,
		},
		{
			name: "north plus 90 produces east",
			args: args{
				base: north,
				dir:  "R",
				amnt: 90,
			},
			want: east,
		},
		{
			name: "north plus 180 produces south",
			args: args{
				base: north,
				dir:  "R",
				amnt: 180,
			},
			want: south,
		},
		{
			name: "north plus 270 produces west",
			args: args{
				base: north,
				dir:  "R",
				amnt: 270,
			},
			want: west,
		},
		{
			name: "north minus 270 produces east",
			args: args{
				base: north,
				dir:  "L",
				amnt: 270,
			},
			want: east,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, newDirection(tt.args.base, tt.args.dir, tt.args.amnt))
		})
	}
}

func Test_moveShip(t *testing.T) {
	// starter ship for all tests.
	s := ship{
		X:         0,
		Y:         0,
		Direction: east,
	}
	type args struct {
		s           ship
		instruction string
	}
	tests := []struct {
		name string
		args args
		want ship
	}{
		{
			name: "forward 5 moves starter ship east 5",
			args: args{
				s:           s,
				instruction: "F5",
			},
			want: ship{
				Direction: east,
				X:         5,
				Y:         0,
			},
		},
		{
			name: "east 7 moves starter ship east 7",
			args: args{
				s:           s,
				instruction: "E7",
			},
			want: ship{
				Direction: east,
				X:         7,
				Y:         0,
			},
		},
		{
			name: "west 6 moves starter ship west 6",
			args: args{
				s:           s,
				instruction: "W6",
			},
			want: ship{
				Direction: east,
				X:         -6,
				Y:         0,
			},
		},
		{
			name: "south 8 moves starter ship south 8",
			args: args{
				s:           s,
				instruction: "S8",
			},
			want: ship{
				Direction: east,
				X:         0,
				Y:         8,
			},
		},
		{
			name: "north 9 moves starter ship north 9",
			args: args{
				s:           s,
				instruction: "N9",
			},
			want: ship{
				Direction: east,
				X:         0,
				Y:         -9,
			},
		},
		{
			name: "right 90 moves ship nowhere, but will face south",
			args: args{
				s:           s,
				instruction: "R90",
			},
			want: ship{
				Direction: south,
				X:         0,
				Y:         0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, moveShip(tt.args.s, tt.args.instruction))
		})
	}
}

func Test_manhattanDistance(t *testing.T) {
	type args struct {
		s ship
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "returns manhattan distance for the ship with both coordinates plus",
			args: args{
				s: ship{
					X:         17,
					Y:         8,
					Direction: south,
				},
			},
			want: 25,
		},
		{
			name: "returns manhattan distance for the ship with both coordinates minus",
			args: args{
				s: ship{
					X:         -17,
					Y:         -8,
					Direction: south,
				},
			},
			want: 25,
		},
		{
			name: "returns manhattan distance for the ship with one coordinate minus",
			args: args{
				s: ship{
					X:         -17,
					Y:         8,
					Direction: south,
				},
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, manhattanDistance(tt.args.s))
		})
	}
}

func Test_navigate(t *testing.T) {
	s := ship{
		X:         0,
		Y:         0,
		Direction: east,
	}
	type args struct {
		s   ship
		ins []string
	}
	tests := []struct {
		name string
		args args
		want ship
	}{
		{
			name: "navigates ship to given coordinates",
			args: args{
				s: s,
				ins: []string{
					"F10",
					"N3",
					"F7",
					"R90",
					"F11",
				},
			},
			want: ship{
				X:         17,
				Y:         8,
				Direction: south,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, navigate(tt.args.s, tt.args.ins))
		})
	}
}
