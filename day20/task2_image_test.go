package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stitchImage(t *testing.T) {
	type args struct {
		img image2
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "stitches image correctly",
			args: args{
				img: image2{
					W: 3,
					H: 3,
					Tiles: map[int]map[int]tilev2{
						0: {
							0: tilev2{
								Content: []string{
									"123",
									"456",
									"789",
								},
							},
							1: tilev2{
								Content: []string{
									"0ab",
									"cde",
									"fgh",
								},
							},
							2: tilev2{
								Content: []string{
									"ijk",
									"lmn",
									"opq",
								},
							},
						},
						1: {
							0: tilev2{
								Content: []string{
									"rst",
									"uvw",
									"xyz",
								},
							},
							1: tilev2{
								Content: []string{
									"321",
									"654",
									"987",
								},
							},
							2: tilev2{
								Content: []string{
									"ba0",
									"edc",
									"hgf",
								},
							},
						},
						2: {
							0: tilev2{
								Content: []string{
									"kji",
									"nml",
									"qpo",
								},
							},
							1: tilev2{
								Content: []string{
									"tsr",
									"wvu",
									"zyx",
								},
							},
							2: tilev2{
								Content: []string{
									"ABC",
									"DEF",
									"GHI",
								},
							},
						},
					},
				},
			},
			want: []string{
				"1230abijk",
				"456cdelmn",
				"789fghopq",
				"rst321ba0",
				"uvw654edc",
				"xyz987hgf",
				"kjitsrABC",
				"nmlwvuDEF",
				"qpozyxGHI",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, stitchImage(tt.args.img))
		})
	}
}
