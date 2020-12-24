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
