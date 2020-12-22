package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_playString(t *testing.T) {
	type args struct {
		p1 []int
		p2 []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "converts two decks into a string",
			args: args{
				p1: []int{1, 2, 3, 4, 5},
				p2: []int{6, 7, 8, 9, 10},
			},
			want: "1,2,3,4,5|6,7,8,9,10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, playString(tt.args.p1, tt.args.p2))
		})
	}
}
