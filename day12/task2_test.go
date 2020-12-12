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
