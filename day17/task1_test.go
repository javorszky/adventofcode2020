package day17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_grid_setStateAt(t *testing.T) {
	type args struct {
		x     int
		y     int
		z     int
		state string
	}
	tests := []struct {
		name string
		g    grid
		args args
		want grid
	}{
		{
			name: "successfully sets a coördinate in an otherwise empty grid",
			g:    grid{},
			args: args{
				x:     0,
				y:     0,
				z:     0,
				state: active,
			},
			want: grid{
				0: {
					0: {
						0: active,
					},
				},
			},
		},
		{
			name: "successfully changes state on a coördinate that exists in grid",
			g: grid{
				0: {
					0: {
						0: inactive,
					},
				},
			},
			args: args{
				x:     0,
				y:     0,
				z:     0,
				state: active,
			},
			want: grid{
				0: {
					0: {
						0: active,
					},
				},
			},
		},
		{
			name: "successfully adds another cube on the xy plane",
			g: grid{
				0: {
					0: {
						0: inactive,
					},
				},
			},
			args: args{
				x:     0,
				y:     0,
				z:     1,
				state: active,
			},
			want: grid{
				0: {
					0: {
						0: inactive,
						1: active,
					},
				},
			},
		},
		{
			name: "successfully adds another cube on the xz plane",
			g: grid{
				0: {
					0: {
						0: inactive,
					},
				},
			},
			args: args{
				x:     0,
				y:     1,
				z:     0,
				state: active,
			},
			want: grid{
				0: {
					0: {
						0: inactive,
					},
					1: {
						0: active,
					},
				},
			},
		},
		{
			name: "successfully adds another cube on the yz plane",
			g: grid{
				0: {
					0: {
						0: inactive,
					},
				},
			},
			args: args{
				x:     1,
				y:     0,
				z:     0,
				state: active,
			},
			want: grid{
				0: {
					0: {
						0: inactive,
					},
				},
				1: {
					0: {
						0: active,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.setStateAt(tt.args.x, tt.args.y, tt.args.z, tt.args.state))
		})
	}
}

func Test_grid_state(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		g    grid
		args args
		want string
	}{
		{
			name: "returns active state of an active cube in a grid where the coördinate exist",
			g: grid{
				0: {
					0: {
						0: active,
					},
				},
			},
			args: args{
				x: 0,
				y: 0,
				z: 0,
			},
			want: active,
		},
		{
			name: "returns inactive state of an inactive cube in a grid where the coördinate exist",
			g: grid{
				0: {
					0: {
						0: inactive,
					},
				},
			},
			args: args{
				x: 0,
				y: 0,
				z: 0,
			},
			want: inactive,
		},
		{
			name: "returns state of a cube in a grid where the z coördinate does not exist",
			g: grid{
				0: {
					0: {
						0: active,
					},
				},
			},
			args: args{
				x: 0,
				y: 0,
				z: 1,
			},
			want: inactive,
		},
		{
			name: "returns state of a cube in a grid where the y coördinate does not exist",
			g: grid{
				0: {
					0: {
						0: active,
					},
				},
			},
			args: args{
				x: 0,
				y: 1,
				z: 0,
			},
			want: inactive,
		},
		{
			name: "returns state of a cube in a grid where the x coördinate does not exist",
			g: grid{
				0: {
					0: {
						0: active,
					},
				},
			},
			args: args{
				x: 1,
				y: 0,
				z: 0,
			},
			want: inactive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.state(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func Test_grid_neighbours(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		g    grid
		args args
		want grid
	}{
		{
			name: "gets non-existent neighbours of a singular existing cube",
			g: grid{
				9: {
					8: {
						7: inactive,
					},
				},
			},
			args: args{
				x: 9,
				y: 8,
				z: 7,
			},
			want: grid{
				8: {
					7: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					8: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
				9: {
					7: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					8: {
						6: inactive,
						8: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
				10: {
					7: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					8: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
			},
		},
		{
			name: "gets mostly non-existent neighbours of an existing cube",
			g: grid{
				9: {
					8: {
						7: inactive,
						8: active,
					},
					7: {
						7: active,
					},
				},
				10: {
					7: {
						8: active,
					},
				},
			},
			args: args{
				x: 9,
				y: 8,
				z: 7,
			},
			want: grid{
				8: {
					7: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					8: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
				9: {
					7: {
						6: inactive,
						7: active,
						8: inactive,
					},
					8: {
						6: inactive,
						8: active,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
				10: {
					7: {
						6: inactive,
						7: inactive,
						8: active,
					},
					8: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
			},
		},
		{
			name: "gets mostly non-existent neighbours of a non-existing cube",
			g: grid{
				9: {
					8: {
						7: active,
					},
				},
			},
			args: args{
				x: 9,
				y: 8,
				z: 6,
			},
			want: grid{
				8: {
					7: {
						6: inactive,
						7: inactive,
						5: inactive,
					},
					8: {
						6: inactive,
						7: inactive,
						5: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						5: inactive,
					},
				},
				9: {
					7: {
						6: inactive,
						7: inactive,
						5: inactive,
					},
					8: {
						7: active,
						5: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						5: inactive,
					},
				},
				10: {
					7: {
						6: inactive,
						7: inactive,
						5: inactive,
					},
					8: {
						6: inactive,
						7: inactive,
						5: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						5: inactive,
					},
				},
			},
		},
		{
			name: "gets non-existent neighbours of a non-existing cube",
			g: grid{
				1: {
					3: {
						2: active,
					},
				},
			},
			args: args{
				x: 9,
				y: 8,
				z: 7,
			},
			want: grid{
				8: {
					7: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					8: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
				9: {
					7: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					8: {
						6: inactive,
						8: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
				10: {
					7: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					8: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
					9: {
						6: inactive,
						7: inactive,
						8: inactive,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.neighbours(tt.args.x, tt.args.y, tt.args.z))
		})
	}
}

func Test_grid_edges(t *testing.T) {
	tests := []struct {
		name string
		g    grid
		want map[int]map[int]map[int]struct{}
	}{
		{
			name: "gets empty edges for an empty grid",
			g:    grid{},
			want: map[int]map[int]map[int]struct{}{},
		},
		{
			name: "gets edges for a singular cube in grid",
			g: grid{
				0: {
					0: {
						0: active,
					},
				},
			},
			want: map[int]map[int]map[int]struct{}{
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
		{
			name: "gets edges for two cubes in grid",
			g: grid{
				0: {
					0: {
						0: active,
						1: inactive,
					},
				},
			},
			want: map[int]map[int]map[int]struct{}{
				-1: {
					-1: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
					0: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
					1: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
				},
				0: {
					-1: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
					0: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
					1: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
				},
				1: {
					-1: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
					0: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
					1: {
						-1: {},
						0:  {},
						1:  {},
						2:  {},
					},
				},
			},
		},
		{
			name: "gets edges for two non-overlapping cubes in grid",
			g: grid{
				0: {
					0: {
						0: active,
					},
				},
				4: {
					4: {
						4: inactive,
					},
				},
			},
			want: map[int]map[int]map[int]struct{}{
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
				3: {
					3: {
						3: {},
						4: {},
						5: {},
					},
					4: {
						3: {},
						4: {},
						5: {},
					},
					5: {
						3: {},
						4: {},
						5: {},
					},
				},
				4: {
					3: {
						3: {},
						4: {},
						5: {},
					},
					4: {
						3: {},
						4: {},
						5: {},
					},
					5: {
						3: {},
						4: {},
						5: {},
					},
				},
				5: {
					3: {
						3: {},
						4: {},
						5: {},
					},
					4: {
						3: {},
						4: {},
						5: {},
					},
					5: {
						3: {},
						4: {},
						5: {},
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

func Test_next(t *testing.T) {
	type args struct {
		state      string
		neighbours grid
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "active state with neighbours with 2 active will remain active",
			args: args{
				state: active,
				neighbours: grid{
					0: {
						0: {
							0: active,
							1: active,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: active,
		},
		{
			name: "active state with neighbours with 3 active will remain active",
			args: args{
				state: active,
				neighbours: grid{
					0: {
						0: {
							0: active,
							1: active,
							2: active,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: active,
		},
		{
			name: "active state with neighbours with 1 active will be inactive",
			args: args{
				state: active,
				neighbours: grid{
					0: {
						0: {
							0: active,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: inactive,
		},
		{
			name: "active state with neighbours with 0 active will be inactive",
			args: args{
				state: active,
				neighbours: grid{
					0: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: inactive,
		},
		{
			name: "active state with neighbours with 4 active will be inactive",
			args: args{
				state: active,
				neighbours: grid{
					0: {
						0: {
							0: active,
							1: active,
							2: active,
						},
						1: {
							0: active,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: inactive,
		},
		{
			name: "inactive state with neighbours with 3 active will become active",
			args: args{
				state: inactive,
				neighbours: grid{
					0: {
						0: {
							0: active,
							1: active,
							2: active,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: active,
		},
		{
			name: "inactive state with neighbours with 2 active will stay inactive",
			args: args{
				state: inactive,
				neighbours: grid{
					0: {
						0: {
							0: active,
							1: active,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: inactive,
		},
		{
			name: "inactive state with neighbours with 4 active will stay inactive",
			args: args{
				state: inactive,
				neighbours: grid{
					0: {
						0: {
							0: active,
							1: active,
							2: active,
						},
						1: {
							0: active,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: inactive,
		},
		{
			name: "inactive state with neighbours with 0 active will stay inactive",
			args: args{
				state: inactive,
				neighbours: grid{
					0: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					1: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
					2: {
						0: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						1: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
						2: {
							0: inactive,
							1: inactive,
							2: inactive,
						},
					},
				},
			},
			want: inactive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := next(tt.args.state, tt.args.neighbours); got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grid_actives(t *testing.T) {
	tests := []struct {
		name string
		g    grid
		want int
	}{
		{
			name: "counts the number of actives on the grid correctly",
			g: grid{
				0: {
					0: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					1: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					2: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					3: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
				},
				1: {
					0: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					1: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					2: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					3: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
				},
				2: {
					0: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					1: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					2: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					3: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
				},
				3: {
					0: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					1: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					2: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
					3: {
						0: inactive,
						1: active,
						2: active,
						3: inactive,
					},
				},
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.g.actives())
		})
	}
}
