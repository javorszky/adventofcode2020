package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_seaMonsterOffsets(t *testing.T) {
	type args struct {
		sm         []string
		lineLength int
	}
	tests := []struct {
		name    string
		args    args
		first   int
		offsets []int
	}{
		{
			name: "calculates offsets for original seamonster for line length 35 beginning with 0 for the first hash",
			args: args{
				sm: []string{
					"                  # ",
					"#    ##    ##    ###",
					" #  #  #  #  #  #   ",
				},
				lineLength: 35,
			},
			first:   18,
			offsets: []int{0, 17, 22, 23, 28, 29, 34, 35, 36, 53, 56, 59, 62, 65, 68},
		},
		{
			name: "calculates offsets for original seamonster for line length 20",
			args: args{
				sm: []string{
					"                  # ",
					"#    ##    ##    ###",
					" #  #  #  #  #  #   ",
				},
				lineLength: 20,
			},
			first:   18,
			offsets: []int{0, 2, 7, 8, 13, 14, 19, 20, 21, 23, 26, 29, 32, 35, 38},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first, offsets := seaMonsterOffsets(tt.args.sm, tt.args.lineLength)
			assert.Equal(t, tt.first, first)
			assert.Equal(t, tt.offsets, offsets)
		})
	}
}
