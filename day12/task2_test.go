package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newDirection2(t *testing.T) {
	s := ship2{
		WX: 10,
		WY: -1,
		SX: 0,
		SY: 0,
	}
	type args struct {
		s    ship2
		dir  string
		amnt int
	}
	tests := []struct {
		name string
		args args
		want ship2
	}{
		{
			name: "rotate waypoint 90 deg right around ship",
			args: args{
				s:    s,
				dir:  "R",
				amnt: 90,
			},
			want: ship2{
				SX: 0,
				SY: 0,
				WX: 1,
				WY: 10,
			},
		},
		{
			name: "rotate waypoint 180 deg right around ship",
			args: args{
				s:    s,
				dir:  "R",
				amnt: 180,
			},
			want: ship2{
				SX: 0,
				SY: 0,
				WX: -10,
				WY: 1,
			},
		},
		{
			name: "rotate waypoint 270 deg right around ship",
			args: args{
				s:    s,
				dir:  "R",
				amnt: 270,
			},
			want: ship2{
				SX: 0,
				SY: 0,
				WX: -1,
				WY: -10,
			},
		},
		{
			name: "rotate waypoint 90 deg left around ship (same as r270)",
			args: args{
				s:    s,
				dir:  "L",
				amnt: 90,
			},
			want: ship2{
				SX: 0,
				SY: 0,
				WX: -1,
				WY: -10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, newDirection2(tt.args.s, tt.args.dir, tt.args.amnt))
		})
	}
}

func Test_moveShip2(t *testing.T) {
	s := ship2{
		WX: 10,
		WY: -1,
		SX: 8,
		SY: -6,
	}
	type args struct {
		s           ship2
		instruction string
	}
	tests := []struct {
		name string
		args args
		want ship2
	}{
		{
			name: "east 1 moves waypoint 1 to the east",
			args: args{
				s:           s,
				instruction: "E1",
			},
			want: ship2{
				WX: 11,
				WY: -1,
				SX: 8,
				SY: -6,
			},
		},
		{
			name: "west 1 moves waypoint 1 to the west",
			args: args{
				s:           s,
				instruction: "W1",
			},
			want: ship2{
				WX: 9,
				WY: -1,
				SX: 8,
				SY: -6,
			},
		},
		{
			name: "north 1 moves waypoint 1 to the north",
			args: args{
				s:           s,
				instruction: "N1",
			},
			want: ship2{
				WX: 10,
				WY: -2,
				SX: 8,
				SY: -6,
			},
		},
		{
			name: "south 1 moves waypoint 1 to the south",
			args: args{
				s:           s,
				instruction: "S1",
			},
			want: ship2{
				WX: 10,
				WY: 0,
				SX: 8,
				SY: -6,
			},
		},
		{
			name: "forward 6 moves ship 6 times towards the vector of the waypoint",
			args: args{
				s:           s,
				instruction: "F6",
			},
			want: ship2{
				WX: 10,
				WY: -1,
				SX: 68,
				SY: -12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, moveShip2(tt.args.s, tt.args.instruction))
		})
	}
}

func Test_navigate2(t *testing.T) {
	s := ship2{
		WX: 10,
		WY: -1,
		SX: 0,
		SY: 0,
	}
	type args struct {
		s   ship2
		ins []string
	}
	tests := []struct {
		name string
		args args
		want ship2
	}{
		{
			name: "navigates ship 2 to given coordinates",
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
			want: ship2{
				WX: 4,
				WY: 10,
				SX: 214,
				SY: 72,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, navigate2(tt.args.s, tt.args.ins))
		})
	}
}
