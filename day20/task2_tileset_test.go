package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newTileSet2(t *testing.T) {
	type args struct {
		t tilev2
	}
	tests := []struct {
		name string
		args args
		want tileSet2
	}{
		{
			name: "creates tileset for tile",
			args: args{
				t: tilev2{
					ID:     "2311000",
					Top:    "123456",
					Right:  "6bhn",
					Bottom: "ijklmn",
					Left:   "17ci",
					Content: []string{
						"890a",
						"defg",
					},
				},
			},
			want: tileSet2{
				{ // upright
					ID:     "2311000",
					Top:    "123456",
					Right:  "6bhn",
					Bottom: "ijklmn",
					Left:   "17ci",
					Content: []string{
						"890a",
						"defg",
					},
				},
				{ // upright rotated
					ID:     "2311001",
					Top:    "ic71",
					Right:  "123456",
					Bottom: "nhb6",
					Left:   "ijklmn",
					Content: []string{
						"d8",
						"e9",
						"f0",
						"ga",
					},
				},
				{ // fliph
					ID:     "2311010",
					Top:    "654321",
					Right:  "17ci",
					Bottom: "nmlkji",
					Left:   "6bhn",
					Content: []string{
						"a098",
						"gfed",
					},
				},
				{ // fliph rotated
					ID:     "2311011",
					Top:    "nhb6",
					Right:  "654321",
					Bottom: "ic71",
					Left:   "nmlkji",
					Content: []string{
						"ga",
						"f0",
						"e9",
						"d8",
					},
				},
				{ // flipv
					ID:     "2311100",
					Top:    "ijklmn",
					Right:  "nhb6",
					Bottom: "123456",
					Left:   "ic71",
					Content: []string{
						"defg",
						"890a",
					},
				},
				{ // flipv rotated
					ID:     "2311101",
					Top:    "17ci",
					Right:  "ijklmn",
					Bottom: "6bhn",
					Left:   "123456",
					Content: []string{
						"8d",
						"9e",
						"0f",
						"ag",
					},
				},
				{ // flipv fliph
					ID:     "2311110",
					Top:    "nmlkji",
					Right:  "ic71",
					Bottom: "654321",
					Left:   "nhb6",
					Content: []string{
						"gfed",
						"a098",
					},
				},
				{ // flipv fliph rotated
					ID:     "2311111",
					Top:    "6bhn",
					Right:  "nmlkji",
					Bottom: "17ci",
					Left:   "654321",
					Content: []string{
						"ag",
						"0f",
						"9e",
						"8d",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, newTileSet2(tt.args.t))
		})
	}
}
