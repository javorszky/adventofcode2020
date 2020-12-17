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
