package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getXY(t *testing.T) {
	type args struct {
		i image
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantx   int
		wanty   int
		wantErr bool
	}{
		{
			name: "gets error when out of bounds number is requested below",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 0,
			},
			wantx:   0,
			wanty:   0,
			wantErr: true,
		},
		{
			name: "gets error when out of bounds number is requested above",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 99,
			},
			wantx:   0,
			wanty:   0,
			wantErr: true,
		},
		{
			name: "gets correct coordinates for 1 in 3x3 grid",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 1,
			},
			wantx:   0,
			wanty:   0,
			wantErr: false,
		},
		{
			name: "gets correct coordinates for 2 in 3x3 grid",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 2,
			},
			wantx:   1,
			wanty:   0,
			wantErr: false,
		},
		{
			name: "gets correct coordinates for 3 in 3x3 grid",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 3,
			},
			wantx:   2,
			wanty:   0,
			wantErr: false,
		},
		{
			name: "gets correct coordinates for 4 in 3x3 grid",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 4,
			},
			wantx:   0,
			wanty:   1,
			wantErr: false,
		},
		{
			name: "gets correct coordinates for 6 in 3x3 grid",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 6,
			},
			wantx:   2,
			wanty:   1,
			wantErr: false,
		},
		{
			name: "gets correct coordinates for 7 in 3x3 grid",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 7,
			},
			wantx:   0,
			wanty:   2,
			wantErr: false,
		},
		{
			name: "gets correct coordinates for 9 in 3x3 grid",
			args: args{
				i: image{
					W: 3,
					H: 3,
				},
				n: 9,
			},
			wantx:   2,
			wanty:   2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y, err := getXY(tt.args.i, tt.args.n)
			assert.Equal(t, tt.wantx, x)
			assert.Equal(t, tt.wanty, y)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
