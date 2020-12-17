package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hyperGrid_actives_after_cycling(t *testing.T) {
	tests := []struct {
		name   string
		g      hyperGrid
		cycles int
		want   int
	}{
		{
			name: "aoc task 1 example input and 6 cycle",
			g: hyperGrid{
				0: {
					0: {
						0: {
							0: inactive,
							1: active,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: active,
						},
						2: {
							0: active,
							1: active,
							2: active,
						},
					},
				},
			},
			cycles: 6,
			want:   848,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.g
			for i := 0; i < tt.cycles; i++ {
				g = g.cycle()
			}
			assert.Equal(t, tt.want, g.actives())
		})
	}
}

func Test_hyperGrid_edges(t *testing.T) {
	tests := []struct {
		name string
		g    hyperGrid
		want map[int]map[int]map[int]map[int]struct{}
	}{
		{
			name: "gets the edges of a single hypercube correctly",
			g: hyperGrid{
				0: {
					0: {
						0: {
							0: active,
						},
					},
				},
			},
			want: map[int]map[int]map[int]map[int]struct{}{
				-1: {
					-1: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
					0: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
					1: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
				},
				0: {
					-1: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
					0: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
					1: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
				},
				1: {
					-1: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
					0: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
					1: {
						-1: {
							-1: {},
							0:  {},
							1:  {},
						},
						0: {
							-1: {},
							0:  {},
							1:  {},
						},
						1: {
							-1: {},
							0:  {},
							1:  {},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.edges())
		})
	}
}

func Test_hyperGrid_neighbours(t *testing.T) {
	type args struct {
		x int
		y int
		z int
		w int
	}
	tests := []struct {
		name string
		g    hyperGrid
		args args
		want hyperGrid
	}{
		{
			name: "gets neighbours of a point in hyperGrid",
			g: hyperGrid{
				0: {
					0: {
						0: {
							0: active,
						},
					},
				},
			},
			args: args{
				x: 0,
				y: 0,
				z: 0,
				w: 0,
			},
			want: hyperGrid{
				-1: {
					-1: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
					0: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
					1: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
				},
				0: {
					-1: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
					0: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
					1: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
				},
				1: {
					-1: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
					0: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
					1: {
						-1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						0: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
						1: {
							-1: inactive,
							0:  inactive,
							1:  inactive,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.neighbours(tt.args.x, tt.args.y, tt.args.z, tt.args.w))
		})
	}
}
