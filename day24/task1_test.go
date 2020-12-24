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
			name: "resolves opposite directions, and collapses them into a simpler one",
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
				se: 0,
				nw: 0,
				sw: 2,
				ne: 0,
				e:  0,
				w:  5,
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
			if got := getCoordinate(tt.args.d); got != tt.want {
				t.Errorf("getCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}
